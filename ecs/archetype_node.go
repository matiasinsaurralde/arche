package ecs

import (
	"reflect"
	"unsafe"

	"github.com/mlange-42/arche/ecs/stats"
)

// archNode is a node in the archetype graph.
type archNode struct {
	*nodeData
	Mask             Mask             // Mask of the archetype
	TransitionAdd    idMap[*archNode] // Mapping from component ID to add to the resulting archetype
	TransitionRemove idMap[*archNode] // Mapping from component ID to remove to the resulting archetype
	Relation         int8             // The node's relation component ID. Negative value stands for no relation
	IsActive         bool
	HasRelation      bool
}

type nodeData struct {
	Ids               []ID                  // List of component IDs
	Types             []reflect.Type        // Component type per column
	archetype         *archetype            // The single archetype for nodes without entity
	archetypes        pagedSlice[archetype] // Storage for archetypes in nodes with entity relation
	archetypeMap      map[Entity]*archetype // Mapping from relation targets to archetypes
	freeIndices       []int32               // Indices of free/inactive archetypes
	zeroValue         []byte                // Used as source for setting storage to zero
	zeroPointer       unsafe.Pointer        // Points to zeroValue for fast access
	capacityIncrement uint32                // Capacity increment
}

// Creates a new archNode
func newArchNode(mask Mask, data *nodeData, relation int8, capacityIncrement int, components []componentType) archNode {
	var arch map[Entity]*archetype
	if relation >= 0 {
		arch = map[Entity]*archetype{}
	}
	ids := make([]ID, len(components))
	types := make([]reflect.Type, len(components))

	var maxSize uintptr = 0
	prev := -1
	for i, c := range components {
		if int(c.ID) <= prev {
			panic("component arguments must be sorted by ID")
		}
		prev = int(c.ID)

		ids[i] = c.ID
		types[i] = c.Type
		size, align := c.Type.Size(), uintptr(c.Type.Align())
		size = (size + (align - 1)) / align * align
		if size > maxSize {
			maxSize = size
		}
	}

	var zeroValue []byte
	var zeroPointer unsafe.Pointer
	if maxSize > 0 {
		zeroValue = make([]byte, maxSize)
		zeroPointer = unsafe.Pointer(&zeroValue[0])
	}

	data.Ids = ids
	data.Types = types
	data.archetypeMap = arch
	data.capacityIncrement = uint32(capacityIncrement)
	data.zeroValue = zeroValue
	data.zeroPointer = zeroPointer

	return archNode{
		nodeData:         data,
		Mask:             mask,
		TransitionAdd:    newIDMap[*archNode](),
		TransitionRemove: newIDMap[*archNode](),
		Relation:         relation,
		HasRelation:      relation >= 0,
	}
}

// Matches the archetype node against a filter.
// Ignores the relation target.
func (a *archNode) Matches(f Filter) bool {
	return f.Matches(a.Mask, nil)
}

// Archetypes of the node.
// Returns a single wrapped archetype if there are no relations.
// Returns nil if the node has no archetype(s).
func (a *archNode) Archetypes() archetypes {
	if a.archetype == nil {
		return &a.archetypes
	}
	return singleArchetype{Archetype: a.archetype}
}

// GetArchetype returns the archetype for the given relation target.
//
// The target is ignored if the node has no relation component.
func (a *archNode) GetArchetype(target Entity) *archetype {
	if a.Relation >= 0 {
		return a.archetypeMap[target]
	}
	return a.archetype
}

// SetArchetype sets the archetype for a node without a relation.
//
// Do not use on nodes without a relation component!
func (a *archNode) SetArchetype(arch *archetype) {
	a.archetype = arch
}

// CreateArchetype creates a new archetype in nodes with relation component.
func (a *archNode) CreateArchetype(target Entity) *archetype {
	var arch *archetype
	var archIndex int32
	lenFree := len(a.freeIndices)
	if lenFree > 0 {
		archIndex = a.freeIndices[lenFree-1]
		arch = a.archetypes.Get(archIndex)
		a.freeIndices = a.freeIndices[:lenFree-1]
		arch.Activate(target, archIndex)
	} else {
		a.archetypes.Add(archetype{})
		archIndex := a.archetypes.Len() - 1
		arch = a.archetypes.Get(archIndex)
		arch.Init(a, archIndex, true, target)
	}
	a.archetypeMap[target] = arch
	return arch
}

// RemoveArchetype de-activates an archetype.
// The archetype will be re-used by CreateArchetype.
func (a *archNode) RemoveArchetype(arch *archetype) {
	delete(a.archetypeMap, arch.RelationTarget)
	idx := arch.index
	a.freeIndices = append(a.freeIndices, idx)
	a.archetypes.Get(idx).Deactivate()
}

// Stats generates statistics for an archetype node.
func (a *archNode) Stats(reg *componentRegistry[ID]) stats.NodeStats {
	ids := a.Ids
	aCompCount := len(ids)
	aTypes := make([]reflect.Type, aCompCount)
	for j, id := range ids {
		aTypes[j], _ = reg.ComponentType(id)
	}

	arches := a.Archetypes()
	var numArches int32
	cap := 0
	count := 0
	memory := 0
	var archStats []stats.ArchetypeStats
	if arches != nil {
		numArches = arches.Len()
		archStats = make([]stats.ArchetypeStats, numArches)
		var i int32
		for i = 0; i < numArches; i++ {
			archStats[i] = arches.Get(i).Stats(reg)
			stats := &archStats[i]
			cap += stats.Capacity
			count += stats.Size
			memory += stats.Memory
		}
	}

	memPerEntity := 0
	for j := range ids {
		memPerEntity += int(aTypes[j].Size())
	}

	return stats.NodeStats{
		ArchetypeCount:       int(numArches),
		ActiveArchetypeCount: int(numArches) - len(a.freeIndices),
		IsActive:             a.IsActive,
		HasRelation:          a.HasRelation,
		Components:           aCompCount,
		ComponentIDs:         ids,
		ComponentTypes:       aTypes,
		Memory:               memory,
		MemoryPerEntity:      memPerEntity,
		Size:                 count,
		Capacity:             cap,
		Archetypes:           archStats,
	}
}

// UpdateStats updates statistics for an archetype node.
func (a *archNode) UpdateStats(stats *stats.NodeStats, reg *componentRegistry[ID]) {
	if !a.IsActive {
		return
	}

	arches := a.Archetypes()

	if !stats.IsActive {
		temp := a.Stats(reg)
		*stats = temp
		return
	}

	cap := 0
	count := 0
	memory := 0

	cntOld := int32(len(stats.Archetypes))
	cntNew := int32(arches.Len())
	var i int32
	for i = 0; i < cntOld; i++ {
		arch := &stats.Archetypes[i]
		arches.Get(i).UpdateStats(stats, arch, reg)
		cap += arch.Capacity
		count += arch.Size
		memory += arch.Memory
	}
	for i = cntOld; i < cntNew; i++ {
		arch := arches.Get(i).Stats(reg)
		stats.Archetypes = append(stats.Archetypes, arch)
		cap += arch.Capacity
		count += arch.Size
		memory += arch.Memory
	}

	stats.IsActive = true
	stats.ArchetypeCount = int(cntNew)
	stats.ActiveArchetypeCount = int(cntNew) - len(a.freeIndices)
	stats.Capacity = cap
	stats.Size = count
	stats.Memory = memory
}
