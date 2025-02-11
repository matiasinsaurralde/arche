package generic

import "github.com/mlange-42/arche/ecs"

// Map provides a type-safe way to access a component type by entity ID.
//
// Create one with [NewMap].
type Map[T any] struct {
	id    ecs.ID
	world *ecs.World
}

// NewMap creates a new [Map] for a component type.
//
// Map provides a type-safe way to access a component type by entity ID.
//
// See also [ecs.World.Get], [ecs.World.Has] and [ecs.World.Set].
func NewMap[T any](w *ecs.World) Map[T] {
	return Map[T]{
		id:    ecs.ComponentID[T](w),
		world: w,
	}
}

// ID returns the component ID for this Map.
func (g *Map[T]) ID() ecs.ID {
	return g.id
}

// Get gets the component for the given entity.
//
// See [Map.HasUnchecked] for an optimized version for static entities.
// See also [ecs.World.Get].
func (g *Map[T]) Get(entity ecs.Entity) *T {
	return (*T)(g.world.Get(entity, g.id))
}

// GetUnchecked gets the component for the given entity.
//
// GetUnchecked is an optimized version of [Map.Get],
// for cases where entities are static or checked with [ecs.World.Alive] in user code.
//
// See also [ecs.World.GetUnchecked].
func (g *Map[T]) GetUnchecked(entity ecs.Entity) *T {
	return (*T)(g.world.GetUnchecked(entity, g.id))
}

// Has returns whether the entity has the component.
//
// See [Map.HasUnchecked] for an optimized version for static entities.
// See also [ecs.World.Has].
func (g *Map[T]) Has(entity ecs.Entity) bool {
	return g.world.Has(entity, g.id)
}

// HasUnchecked returns whether the entity has the component.
//
// HasUnchecked is an optimized version of [Map.Has],
// for cases where entities are static or checked with [ecs.World.Alive] in user code.
//
// See also [ecs.World.HasUnchecked].
func (g *Map[T]) HasUnchecked(entity ecs.Entity) bool {
	return g.world.HasUnchecked(entity, g.id)
}

// Set overwrites the component for the given entity.
//
// Panics if the entity does not have a component of that type.
//
// See also [ecs.World.Set].
func (g *Map[T]) Set(entity ecs.Entity, comp *T) *T {
	return (*T)(g.world.Set(entity, g.id, comp))
}

// GetRelation returns the target entity for the given entity and the Map's relation component.
//
// Panics:
//   - if the entity does not have a component of that type.
//   - if the component is not an [ecs.Relation].
//   - if the entity has been removed/recycled.
//
// See also [ecs.World.GetRelation].
func (g *Map[T]) GetRelation(entity ecs.Entity) ecs.Entity {
	return g.world.Relations().Get(entity, g.id)
}

// GetRelation returns the target entity for the given entity and the Map's relation component.
//
// Returns the zero entity if the entity does not have the given component,
// or if the component is not an [ecs.Relation].
//
// GetRelationUnchecked is an optimized version of [Map.GetRelation].
// Does not check if the entity is alive or that the component ID is applicable.
//
// See also [ecs.World.GetRelationUnchecked].
func (g *Map[T]) GetRelationUnchecked(entity ecs.Entity) ecs.Entity {
	return g.world.Relations().GetUnchecked(entity, g.id)
}

// SetRelation sets the target entity for the given entity and the Map's relation component.
//
// Panics if the entity does not have a component of that type.
// Panics if the component is not a relation.
//
// See also [ecs.World.SetRelation].
func (g *Map[T]) SetRelation(entity, target ecs.Entity) {
	g.world.Relations().Set(entity, g.id, target)
}
