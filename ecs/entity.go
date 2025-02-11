package ecs

import "reflect"

// Reflection type of an [Entity].
var entityType = reflect.TypeOf(Entity{})

// Size of an [Entity] in memory, in bytes.
var entitySize = entityType.Size()

// Size of an [entityIndex] in memory.
var entityIndexSize = reflect.TypeOf(entityIndex{}).Size()

// Entity identifier.
// Holds an entity ID and it's generation for recycling.
//
// Entities are only created via the [World], using [World.NewEntity] or [World.NewEntityWith].
// Batch creation of entities is possible via [Builder].
//
// Entities are intended to be stored and passed around via copy, not via pointers.
type Entity struct {
	id  eid    // Entity ID
	gen uint16 // Entity generation
}

// newEntity creates a new Entity.
func newEntity(id eid) Entity {
	return Entity{id, 0}
}

// newEntityGen creates a new Entity with a given generation.
func newEntityGen(id eid, gen uint16) Entity {
	return Entity{id, gen}
}

// IsZero returns whether this entity is the reserved zero entity.
func (e Entity) IsZero() bool {
	return e.id == 0
}

// entityIndex indicates where an entity is currently stored.
type entityIndex struct {
	arch  *archetype // Entity's current archetype
	index uintptr    // Entity's current index in the archetype
}
