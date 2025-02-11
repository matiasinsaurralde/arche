package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterCache(t *testing.T) {
	world := NewWorld()
	posID := ComponentID[Position](&world)
	velID := ComponentID[Velocity](&world)
	rotID := ComponentID[rotation](&world)

	cache := world.Cache()

	world.NewEntity()
	world.NewEntity(posID, velID)
	world.NewEntity(posID, velID, rotID)

	all1 := All(posID, velID)
	all2 := All(posID, velID, rotID)

	f1 := cache.Register(all1)
	f2 := cache.Register(all2)
	assert.Equal(t, 0, int(f1.id))
	assert.Equal(t, 1, int(f2.id))

	assert.Equal(t, 2, len(world.getArchetypes(&f1).pointers))
	assert.Equal(t, 1, len(world.getArchetypes(&f2).pointers))

	assert.Panics(t, func() { cache.Register(&f2) })

	e1 := cache.get(&f1)
	e2 := cache.get(&f2)

	assert.Equal(t, f1.filter, e1.Filter)
	assert.Equal(t, f2.filter, e2.Filter)

	ff1 := cache.Unregister(&f1)
	ff2 := cache.Unregister(&f2)

	assert.Equal(t, all1, ff1)
	assert.Equal(t, all2, ff2)

	assert.Panics(t, func() { cache.Unregister(&f1) })
	assert.Panics(t, func() { cache.get(&f1) })
}

func ExampleCache() {
	world := NewWorld()
	posID := ComponentID[Position](&world)

	filter := All(posID)
	cached := world.Cache().Register(filter)
	query := world.Query(&cached)

	for query.Next() {
		// ...
	}
	// Output:
}
