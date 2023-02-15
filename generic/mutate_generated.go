package generic

// Code generated by go generate; DO NOT EDIT.

import (
	"reflect"

	"github.com/mlange-42/arche/ecs"
)

//////////////////////////////////////////////////////////////////////////

// Mutate1 is a helper for mutating one components.
//
// It can be used to create entities ([Mutate1.NewEntity] and [Mutate1.NewEntityWith]),
// and to add or remove components ([Mutate1.Add], [Mutate1.Assign] and [Mutate1.Remove]).
type Mutate1[A any] struct {
	ids      []ecs.ID
	exchange []ecs.ID
	world    *ecs.World
}

// NewMutate1 creates a new Mutate1 object.
func NewMutate1[A any](w *ecs.World) *Mutate1[A] {
	return &Mutate1[A]{
		ids:   []ecs.ID{ecs.ComponentID[A](w)},
		world: w,
	}
}

// WithExchange sets components to remove in calls to [Mutate1.Exchange].
//
// Create the required mask with [Mask1], [Mask2], etc.
func (m *Mutate1[A]) WithExchange(remove []reflect.Type) *Mutate1[A] {
	m.exchange = toIds(m.world, remove)
	return m
}

// NewEntity creates a new [ecs.Entity] with the Mutate1's components.
//
// See also [ecs.World.NewEntity].
func (m *Mutate1[A]) NewEntity() (ecs.Entity, *A) {
	entity := m.world.NewEntity(m.ids...)
	return entity, (*A)(m.world.Get(entity, m.ids[0]))
}

// NewEntityWith creates a new [ecs.Entity] with the Mutate1's components, using the supplied values.
//
// See also [ecs.World.NewEntityWith].
func (m *Mutate1[A]) NewEntityWith(a *A) (ecs.Entity, *A) {
	entity := m.world.NewEntityWith(
		ecs.Component{ID: m.ids[0], Component: a},
	)
	return entity, (*A)(m.world.Get(entity, m.ids[0]))
}

// Add the Mutate1's components to the given entity.
//
// See also [ecs.World.Add].
func (m *Mutate1[A]) Add(entity ecs.Entity) *A {
	m.world.Add(entity, m.ids...)
	return (*A)(m.world.Get(entity, m.ids[0]))
}

// Assign the Mutate1's components to the given entity, using the supplied values.
//
// See also [ecs.World.Assign] and [ecs.World.AssignN].
func (m *Mutate1[A]) Assign(entity ecs.Entity, a *A) *A {
	m.world.AssignN(entity,
		ecs.Component{ID: m.ids[0], Component: a},
	)
	return (*A)(m.world.Get(entity, m.ids[0]))
}

// Remove the Mutate1's components from the given entity.
//
// See also [ecs.World.Remove].
func (m *Mutate1[A]) Remove(entity ecs.Entity) {
	m.world.Remove(entity, m.ids...)
}

// Exchange components on an entity.
//
// Removes the components set via [Mutate1.WithExchange].
// Adds Mutate1's components as with [Mutate1.Add].
//
// See also [ecs.World.Exchange].
func (m *Mutate1[A]) Exchange(entity ecs.Entity) *A {
	m.world.Exchange(entity, m.ids, m.exchange)
	return (*A)(m.world.Get(entity, m.ids[0]))
}

//////////////////////////////////////////////////////////////////////////

// Mutate2 is a helper for mutating two components.
//
// It can be used to create entities ([Mutate2.NewEntity] and [Mutate2.NewEntityWith]),
// and to add or remove components ([Mutate2.Add], [Mutate2.Assign] and [Mutate2.Remove]).
type Mutate2[A any, B any] struct {
	ids      []ecs.ID
	exchange []ecs.ID
	world    *ecs.World
}

// NewMutate2 creates a new Mutate2 object.
func NewMutate2[A any, B any](w *ecs.World) *Mutate2[A, B] {
	return &Mutate2[A, B]{
		ids:   []ecs.ID{ecs.ComponentID[A](w), ecs.ComponentID[B](w)},
		world: w,
	}
}

// WithExchange sets components to remove in calls to [Mutate2.Exchange].
//
// Create the required mask with [Mask1], [Mask2], etc.
func (m *Mutate2[A, B]) WithExchange(remove []reflect.Type) *Mutate2[A, B] {
	m.exchange = toIds(m.world, remove)
	return m
}

// NewEntity creates a new [ecs.Entity] with the Mutate2's components.
//
// See also [ecs.World.NewEntity].
func (m *Mutate2[A, B]) NewEntity() (ecs.Entity, *A, *B) {
	entity := m.world.NewEntity(m.ids...)
	return entity, (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1]))
}

// NewEntityWith creates a new [ecs.Entity] with the Mutate2's components, using the supplied values.
//
// See also [ecs.World.NewEntityWith].
func (m *Mutate2[A, B]) NewEntityWith(a *A, b *B) (ecs.Entity, *A, *B) {
	entity := m.world.NewEntityWith(
		ecs.Component{ID: m.ids[0], Component: a},
		ecs.Component{ID: m.ids[1], Component: b},
	)
	return entity, (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1]))
}

// Add the Mutate2's components to the given entity.
//
// See also [ecs.World.Add].
func (m *Mutate2[A, B]) Add(entity ecs.Entity) (*A, *B) {
	m.world.Add(entity, m.ids...)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1]))
}

// Assign the Mutate2's components to the given entity, using the supplied values.
//
// See also [ecs.World.Assign] and [ecs.World.AssignN].
func (m *Mutate2[A, B]) Assign(entity ecs.Entity, a *A, b *B) (*A, *B) {
	m.world.AssignN(entity,
		ecs.Component{ID: m.ids[0], Component: a},
		ecs.Component{ID: m.ids[1], Component: b},
	)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1]))
}

// Remove the Mutate2's components from the given entity.
//
// See also [ecs.World.Remove].
func (m *Mutate2[A, B]) Remove(entity ecs.Entity) {
	m.world.Remove(entity, m.ids...)
}

// Exchange components on an entity.
//
// Removes the components set via [Mutate2.WithExchange].
// Adds Mutate2's components as with [Mutate2.Add].
//
// See also [ecs.World.Exchange].
func (m *Mutate2[A, B]) Exchange(entity ecs.Entity) (*A, *B) {
	m.world.Exchange(entity, m.ids, m.exchange)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1]))
}

//////////////////////////////////////////////////////////////////////////

// Mutate3 is a helper for mutating three components.
//
// It can be used to create entities ([Mutate3.NewEntity] and [Mutate3.NewEntityWith]),
// and to add or remove components ([Mutate3.Add], [Mutate3.Assign] and [Mutate3.Remove]).
type Mutate3[A any, B any, C any] struct {
	ids      []ecs.ID
	exchange []ecs.ID
	world    *ecs.World
}

// NewMutate3 creates a new Mutate3 object.
func NewMutate3[A any, B any, C any](w *ecs.World) *Mutate3[A, B, C] {
	return &Mutate3[A, B, C]{
		ids:   []ecs.ID{ecs.ComponentID[A](w), ecs.ComponentID[B](w), ecs.ComponentID[C](w)},
		world: w,
	}
}

// WithExchange sets components to remove in calls to [Mutate3.Exchange].
//
// Create the required mask with [Mask1], [Mask2], etc.
func (m *Mutate3[A, B, C]) WithExchange(remove []reflect.Type) *Mutate3[A, B, C] {
	m.exchange = toIds(m.world, remove)
	return m
}

// NewEntity creates a new [ecs.Entity] with the Mutate3's components.
//
// See also [ecs.World.NewEntity].
func (m *Mutate3[A, B, C]) NewEntity() (ecs.Entity, *A, *B, *C) {
	entity := m.world.NewEntity(m.ids...)
	return entity, (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2]))
}

// NewEntityWith creates a new [ecs.Entity] with the Mutate3's components, using the supplied values.
//
// See also [ecs.World.NewEntityWith].
func (m *Mutate3[A, B, C]) NewEntityWith(a *A, b *B, c *C) (ecs.Entity, *A, *B, *C) {
	entity := m.world.NewEntityWith(
		ecs.Component{ID: m.ids[0], Component: a},
		ecs.Component{ID: m.ids[1], Component: b},
		ecs.Component{ID: m.ids[2], Component: c},
	)
	return entity, (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2]))
}

// Add the Mutate3's components to the given entity.
//
// See also [ecs.World.Add].
func (m *Mutate3[A, B, C]) Add(entity ecs.Entity) (*A, *B, *C) {
	m.world.Add(entity, m.ids...)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2]))
}

// Assign the Mutate3's components to the given entity, using the supplied values.
//
// See also [ecs.World.Assign] and [ecs.World.AssignN].
func (m *Mutate3[A, B, C]) Assign(entity ecs.Entity, a *A, b *B, c *C) (*A, *B, *C) {
	m.world.AssignN(entity,
		ecs.Component{ID: m.ids[0], Component: a},
		ecs.Component{ID: m.ids[1], Component: b},
		ecs.Component{ID: m.ids[2], Component: c},
	)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2]))
}

// Remove the Mutate3's components from the given entity.
//
// See also [ecs.World.Remove].
func (m *Mutate3[A, B, C]) Remove(entity ecs.Entity) {
	m.world.Remove(entity, m.ids...)
}

// Exchange components on an entity.
//
// Removes the components set via [Mutate3.WithExchange].
// Adds Mutate3's components as with [Mutate3.Add].
//
// See also [ecs.World.Exchange].
func (m *Mutate3[A, B, C]) Exchange(entity ecs.Entity) (*A, *B, *C) {
	m.world.Exchange(entity, m.ids, m.exchange)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2]))
}

//////////////////////////////////////////////////////////////////////////

// Mutate4 is a helper for mutating four components.
//
// It can be used to create entities ([Mutate4.NewEntity] and [Mutate4.NewEntityWith]),
// and to add or remove components ([Mutate4.Add], [Mutate4.Assign] and [Mutate4.Remove]).
type Mutate4[A any, B any, C any, D any] struct {
	ids      []ecs.ID
	exchange []ecs.ID
	world    *ecs.World
}

// NewMutate4 creates a new Mutate4 object.
func NewMutate4[A any, B any, C any, D any](w *ecs.World) *Mutate4[A, B, C, D] {
	return &Mutate4[A, B, C, D]{
		ids:   []ecs.ID{ecs.ComponentID[A](w), ecs.ComponentID[B](w), ecs.ComponentID[C](w), ecs.ComponentID[D](w)},
		world: w,
	}
}

// WithExchange sets components to remove in calls to [Mutate4.Exchange].
//
// Create the required mask with [Mask1], [Mask2], etc.
func (m *Mutate4[A, B, C, D]) WithExchange(remove []reflect.Type) *Mutate4[A, B, C, D] {
	m.exchange = toIds(m.world, remove)
	return m
}

// NewEntity creates a new [ecs.Entity] with the Mutate4's components.
//
// See also [ecs.World.NewEntity].
func (m *Mutate4[A, B, C, D]) NewEntity() (ecs.Entity, *A, *B, *C, *D) {
	entity := m.world.NewEntity(m.ids...)
	return entity, (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3]))
}

// NewEntityWith creates a new [ecs.Entity] with the Mutate4's components, using the supplied values.
//
// See also [ecs.World.NewEntityWith].
func (m *Mutate4[A, B, C, D]) NewEntityWith(a *A, b *B, c *C, d *D) (ecs.Entity, *A, *B, *C, *D) {
	entity := m.world.NewEntityWith(
		ecs.Component{ID: m.ids[0], Component: a},
		ecs.Component{ID: m.ids[1], Component: b},
		ecs.Component{ID: m.ids[2], Component: c},
		ecs.Component{ID: m.ids[3], Component: d},
	)
	return entity, (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3]))
}

// Add the Mutate4's components to the given entity.
//
// See also [ecs.World.Add].
func (m *Mutate4[A, B, C, D]) Add(entity ecs.Entity) (*A, *B, *C, *D) {
	m.world.Add(entity, m.ids...)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3]))
}

// Assign the Mutate4's components to the given entity, using the supplied values.
//
// See also [ecs.World.Assign] and [ecs.World.AssignN].
func (m *Mutate4[A, B, C, D]) Assign(entity ecs.Entity, a *A, b *B, c *C, d *D) (*A, *B, *C, *D) {
	m.world.AssignN(entity,
		ecs.Component{ID: m.ids[0], Component: a},
		ecs.Component{ID: m.ids[1], Component: b},
		ecs.Component{ID: m.ids[2], Component: c},
		ecs.Component{ID: m.ids[3], Component: d},
	)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3]))
}

// Remove the Mutate4's components from the given entity.
//
// See also [ecs.World.Remove].
func (m *Mutate4[A, B, C, D]) Remove(entity ecs.Entity) {
	m.world.Remove(entity, m.ids...)
}

// Exchange components on an entity.
//
// Removes the components set via [Mutate4.WithExchange].
// Adds Mutate4's components as with [Mutate4.Add].
//
// See also [ecs.World.Exchange].
func (m *Mutate4[A, B, C, D]) Exchange(entity ecs.Entity) (*A, *B, *C, *D) {
	m.world.Exchange(entity, m.ids, m.exchange)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3]))
}

//////////////////////////////////////////////////////////////////////////

// Mutate5 is a helper for mutating five components.
//
// It can be used to create entities ([Mutate5.NewEntity] and [Mutate5.NewEntityWith]),
// and to add or remove components ([Mutate5.Add], [Mutate5.Assign] and [Mutate5.Remove]).
type Mutate5[A any, B any, C any, D any, E any] struct {
	ids      []ecs.ID
	exchange []ecs.ID
	world    *ecs.World
}

// NewMutate5 creates a new Mutate5 object.
func NewMutate5[A any, B any, C any, D any, E any](w *ecs.World) *Mutate5[A, B, C, D, E] {
	return &Mutate5[A, B, C, D, E]{
		ids:   []ecs.ID{ecs.ComponentID[A](w), ecs.ComponentID[B](w), ecs.ComponentID[C](w), ecs.ComponentID[D](w), ecs.ComponentID[E](w)},
		world: w,
	}
}

// WithExchange sets components to remove in calls to [Mutate5.Exchange].
//
// Create the required mask with [Mask1], [Mask2], etc.
func (m *Mutate5[A, B, C, D, E]) WithExchange(remove []reflect.Type) *Mutate5[A, B, C, D, E] {
	m.exchange = toIds(m.world, remove)
	return m
}

// NewEntity creates a new [ecs.Entity] with the Mutate5's components.
//
// See also [ecs.World.NewEntity].
func (m *Mutate5[A, B, C, D, E]) NewEntity() (ecs.Entity, *A, *B, *C, *D, *E) {
	entity := m.world.NewEntity(m.ids...)
	return entity, (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4]))
}

// NewEntityWith creates a new [ecs.Entity] with the Mutate5's components, using the supplied values.
//
// See also [ecs.World.NewEntityWith].
func (m *Mutate5[A, B, C, D, E]) NewEntityWith(a *A, b *B, c *C, d *D, e *E) (ecs.Entity, *A, *B, *C, *D, *E) {
	entity := m.world.NewEntityWith(
		ecs.Component{ID: m.ids[0], Component: a},
		ecs.Component{ID: m.ids[1], Component: b},
		ecs.Component{ID: m.ids[2], Component: c},
		ecs.Component{ID: m.ids[3], Component: d},
		ecs.Component{ID: m.ids[4], Component: e},
	)
	return entity, (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4]))
}

// Add the Mutate5's components to the given entity.
//
// See also [ecs.World.Add].
func (m *Mutate5[A, B, C, D, E]) Add(entity ecs.Entity) (*A, *B, *C, *D, *E) {
	m.world.Add(entity, m.ids...)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4]))
}

// Assign the Mutate5's components to the given entity, using the supplied values.
//
// See also [ecs.World.Assign] and [ecs.World.AssignN].
func (m *Mutate5[A, B, C, D, E]) Assign(entity ecs.Entity, a *A, b *B, c *C, d *D, e *E) (*A, *B, *C, *D, *E) {
	m.world.AssignN(entity,
		ecs.Component{ID: m.ids[0], Component: a},
		ecs.Component{ID: m.ids[1], Component: b},
		ecs.Component{ID: m.ids[2], Component: c},
		ecs.Component{ID: m.ids[3], Component: d},
		ecs.Component{ID: m.ids[4], Component: e},
	)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4]))
}

// Remove the Mutate5's components from the given entity.
//
// See also [ecs.World.Remove].
func (m *Mutate5[A, B, C, D, E]) Remove(entity ecs.Entity) {
	m.world.Remove(entity, m.ids...)
}

// Exchange components on an entity.
//
// Removes the components set via [Mutate5.WithExchange].
// Adds Mutate5's components as with [Mutate5.Add].
//
// See also [ecs.World.Exchange].
func (m *Mutate5[A, B, C, D, E]) Exchange(entity ecs.Entity) (*A, *B, *C, *D, *E) {
	m.world.Exchange(entity, m.ids, m.exchange)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4]))
}

//////////////////////////////////////////////////////////////////////////

// Mutate6 is a helper for mutating six components.
//
// It can be used to create entities ([Mutate6.NewEntity] and [Mutate6.NewEntityWith]),
// and to add or remove components ([Mutate6.Add], [Mutate6.Assign] and [Mutate6.Remove]).
type Mutate6[A any, B any, C any, D any, E any, F any] struct {
	ids      []ecs.ID
	exchange []ecs.ID
	world    *ecs.World
}

// NewMutate6 creates a new Mutate6 object.
func NewMutate6[A any, B any, C any, D any, E any, F any](w *ecs.World) *Mutate6[A, B, C, D, E, F] {
	return &Mutate6[A, B, C, D, E, F]{
		ids:   []ecs.ID{ecs.ComponentID[A](w), ecs.ComponentID[B](w), ecs.ComponentID[C](w), ecs.ComponentID[D](w), ecs.ComponentID[E](w), ecs.ComponentID[F](w)},
		world: w,
	}
}

// WithExchange sets components to remove in calls to [Mutate6.Exchange].
//
// Create the required mask with [Mask1], [Mask2], etc.
func (m *Mutate6[A, B, C, D, E, F]) WithExchange(remove []reflect.Type) *Mutate6[A, B, C, D, E, F] {
	m.exchange = toIds(m.world, remove)
	return m
}

// NewEntity creates a new [ecs.Entity] with the Mutate6's components.
//
// See also [ecs.World.NewEntity].
func (m *Mutate6[A, B, C, D, E, F]) NewEntity() (ecs.Entity, *A, *B, *C, *D, *E, *F) {
	entity := m.world.NewEntity(m.ids...)
	return entity, (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4])), (*F)(m.world.Get(entity, m.ids[5]))
}

// NewEntityWith creates a new [ecs.Entity] with the Mutate6's components, using the supplied values.
//
// See also [ecs.World.NewEntityWith].
func (m *Mutate6[A, B, C, D, E, F]) NewEntityWith(a *A, b *B, c *C, d *D, e *E, f *F) (ecs.Entity, *A, *B, *C, *D, *E, *F) {
	entity := m.world.NewEntityWith(
		ecs.Component{ID: m.ids[0], Component: a},
		ecs.Component{ID: m.ids[1], Component: b},
		ecs.Component{ID: m.ids[2], Component: c},
		ecs.Component{ID: m.ids[3], Component: d},
		ecs.Component{ID: m.ids[4], Component: e},
		ecs.Component{ID: m.ids[5], Component: f},
	)
	return entity, (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4])), (*F)(m.world.Get(entity, m.ids[5]))
}

// Add the Mutate6's components to the given entity.
//
// See also [ecs.World.Add].
func (m *Mutate6[A, B, C, D, E, F]) Add(entity ecs.Entity) (*A, *B, *C, *D, *E, *F) {
	m.world.Add(entity, m.ids...)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4])), (*F)(m.world.Get(entity, m.ids[5]))
}

// Assign the Mutate6's components to the given entity, using the supplied values.
//
// See also [ecs.World.Assign] and [ecs.World.AssignN].
func (m *Mutate6[A, B, C, D, E, F]) Assign(entity ecs.Entity, a *A, b *B, c *C, d *D, e *E, f *F) (*A, *B, *C, *D, *E, *F) {
	m.world.AssignN(entity,
		ecs.Component{ID: m.ids[0], Component: a},
		ecs.Component{ID: m.ids[1], Component: b},
		ecs.Component{ID: m.ids[2], Component: c},
		ecs.Component{ID: m.ids[3], Component: d},
		ecs.Component{ID: m.ids[4], Component: e},
		ecs.Component{ID: m.ids[5], Component: f},
	)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4])), (*F)(m.world.Get(entity, m.ids[5]))
}

// Remove the Mutate6's components from the given entity.
//
// See also [ecs.World.Remove].
func (m *Mutate6[A, B, C, D, E, F]) Remove(entity ecs.Entity) {
	m.world.Remove(entity, m.ids...)
}

// Exchange components on an entity.
//
// Removes the components set via [Mutate6.WithExchange].
// Adds Mutate6's components as with [Mutate6.Add].
//
// See also [ecs.World.Exchange].
func (m *Mutate6[A, B, C, D, E, F]) Exchange(entity ecs.Entity) (*A, *B, *C, *D, *E, *F) {
	m.world.Exchange(entity, m.ids, m.exchange)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4])), (*F)(m.world.Get(entity, m.ids[5]))
}

//////////////////////////////////////////////////////////////////////////

// Mutate7 is a helper for mutating seven components.
//
// It can be used to create entities ([Mutate7.NewEntity] and [Mutate7.NewEntityWith]),
// and to add or remove components ([Mutate7.Add], [Mutate7.Assign] and [Mutate7.Remove]).
type Mutate7[A any, B any, C any, D any, E any, F any, G any] struct {
	ids      []ecs.ID
	exchange []ecs.ID
	world    *ecs.World
}

// NewMutate7 creates a new Mutate7 object.
func NewMutate7[A any, B any, C any, D any, E any, F any, G any](w *ecs.World) *Mutate7[A, B, C, D, E, F, G] {
	return &Mutate7[A, B, C, D, E, F, G]{
		ids:   []ecs.ID{ecs.ComponentID[A](w), ecs.ComponentID[B](w), ecs.ComponentID[C](w), ecs.ComponentID[D](w), ecs.ComponentID[E](w), ecs.ComponentID[F](w), ecs.ComponentID[G](w)},
		world: w,
	}
}

// WithExchange sets components to remove in calls to [Mutate7.Exchange].
//
// Create the required mask with [Mask1], [Mask2], etc.
func (m *Mutate7[A, B, C, D, E, F, G]) WithExchange(remove []reflect.Type) *Mutate7[A, B, C, D, E, F, G] {
	m.exchange = toIds(m.world, remove)
	return m
}

// NewEntity creates a new [ecs.Entity] with the Mutate7's components.
//
// See also [ecs.World.NewEntity].
func (m *Mutate7[A, B, C, D, E, F, G]) NewEntity() (ecs.Entity, *A, *B, *C, *D, *E, *F, *G) {
	entity := m.world.NewEntity(m.ids...)
	return entity, (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4])), (*F)(m.world.Get(entity, m.ids[5])), (*G)(m.world.Get(entity, m.ids[6]))
}

// NewEntityWith creates a new [ecs.Entity] with the Mutate7's components, using the supplied values.
//
// See also [ecs.World.NewEntityWith].
func (m *Mutate7[A, B, C, D, E, F, G]) NewEntityWith(a *A, b *B, c *C, d *D, e *E, f *F, g *G) (ecs.Entity, *A, *B, *C, *D, *E, *F, *G) {
	entity := m.world.NewEntityWith(
		ecs.Component{ID: m.ids[0], Component: a},
		ecs.Component{ID: m.ids[1], Component: b},
		ecs.Component{ID: m.ids[2], Component: c},
		ecs.Component{ID: m.ids[3], Component: d},
		ecs.Component{ID: m.ids[4], Component: e},
		ecs.Component{ID: m.ids[5], Component: f},
		ecs.Component{ID: m.ids[6], Component: g},
	)
	return entity, (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4])), (*F)(m.world.Get(entity, m.ids[5])), (*G)(m.world.Get(entity, m.ids[6]))
}

// Add the Mutate7's components to the given entity.
//
// See also [ecs.World.Add].
func (m *Mutate7[A, B, C, D, E, F, G]) Add(entity ecs.Entity) (*A, *B, *C, *D, *E, *F, *G) {
	m.world.Add(entity, m.ids...)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4])), (*F)(m.world.Get(entity, m.ids[5])), (*G)(m.world.Get(entity, m.ids[6]))
}

// Assign the Mutate7's components to the given entity, using the supplied values.
//
// See also [ecs.World.Assign] and [ecs.World.AssignN].
func (m *Mutate7[A, B, C, D, E, F, G]) Assign(entity ecs.Entity, a *A, b *B, c *C, d *D, e *E, f *F, g *G) (*A, *B, *C, *D, *E, *F, *G) {
	m.world.AssignN(entity,
		ecs.Component{ID: m.ids[0], Component: a},
		ecs.Component{ID: m.ids[1], Component: b},
		ecs.Component{ID: m.ids[2], Component: c},
		ecs.Component{ID: m.ids[3], Component: d},
		ecs.Component{ID: m.ids[4], Component: e},
		ecs.Component{ID: m.ids[5], Component: f},
		ecs.Component{ID: m.ids[6], Component: g},
	)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4])), (*F)(m.world.Get(entity, m.ids[5])), (*G)(m.world.Get(entity, m.ids[6]))
}

// Remove the Mutate7's components from the given entity.
//
// See also [ecs.World.Remove].
func (m *Mutate7[A, B, C, D, E, F, G]) Remove(entity ecs.Entity) {
	m.world.Remove(entity, m.ids...)
}

// Exchange components on an entity.
//
// Removes the components set via [Mutate7.WithExchange].
// Adds Mutate7's components as with [Mutate7.Add].
//
// See also [ecs.World.Exchange].
func (m *Mutate7[A, B, C, D, E, F, G]) Exchange(entity ecs.Entity) (*A, *B, *C, *D, *E, *F, *G) {
	m.world.Exchange(entity, m.ids, m.exchange)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4])), (*F)(m.world.Get(entity, m.ids[5])), (*G)(m.world.Get(entity, m.ids[6]))
}

//////////////////////////////////////////////////////////////////////////

// Mutate8 is a helper for mutating eight components.
//
// It can be used to create entities ([Mutate8.NewEntity] and [Mutate8.NewEntityWith]),
// and to add or remove components ([Mutate8.Add], [Mutate8.Assign] and [Mutate8.Remove]).
type Mutate8[A any, B any, C any, D any, E any, F any, G any, H any] struct {
	ids      []ecs.ID
	exchange []ecs.ID
	world    *ecs.World
}

// NewMutate8 creates a new Mutate8 object.
func NewMutate8[A any, B any, C any, D any, E any, F any, G any, H any](w *ecs.World) *Mutate8[A, B, C, D, E, F, G, H] {
	return &Mutate8[A, B, C, D, E, F, G, H]{
		ids:   []ecs.ID{ecs.ComponentID[A](w), ecs.ComponentID[B](w), ecs.ComponentID[C](w), ecs.ComponentID[D](w), ecs.ComponentID[E](w), ecs.ComponentID[F](w), ecs.ComponentID[G](w), ecs.ComponentID[H](w)},
		world: w,
	}
}

// WithExchange sets components to remove in calls to [Mutate8.Exchange].
//
// Create the required mask with [Mask1], [Mask2], etc.
func (m *Mutate8[A, B, C, D, E, F, G, H]) WithExchange(remove []reflect.Type) *Mutate8[A, B, C, D, E, F, G, H] {
	m.exchange = toIds(m.world, remove)
	return m
}

// NewEntity creates a new [ecs.Entity] with the Mutate8's components.
//
// See also [ecs.World.NewEntity].
func (m *Mutate8[A, B, C, D, E, F, G, H]) NewEntity() (ecs.Entity, *A, *B, *C, *D, *E, *F, *G, *H) {
	entity := m.world.NewEntity(m.ids...)
	return entity, (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4])), (*F)(m.world.Get(entity, m.ids[5])), (*G)(m.world.Get(entity, m.ids[6])), (*H)(m.world.Get(entity, m.ids[7]))
}

// NewEntityWith creates a new [ecs.Entity] with the Mutate8's components, using the supplied values.
//
// See also [ecs.World.NewEntityWith].
func (m *Mutate8[A, B, C, D, E, F, G, H]) NewEntityWith(a *A, b *B, c *C, d *D, e *E, f *F, g *G, h *H) (ecs.Entity, *A, *B, *C, *D, *E, *F, *G, *H) {
	entity := m.world.NewEntityWith(
		ecs.Component{ID: m.ids[0], Component: a},
		ecs.Component{ID: m.ids[1], Component: b},
		ecs.Component{ID: m.ids[2], Component: c},
		ecs.Component{ID: m.ids[3], Component: d},
		ecs.Component{ID: m.ids[4], Component: e},
		ecs.Component{ID: m.ids[5], Component: f},
		ecs.Component{ID: m.ids[6], Component: g},
		ecs.Component{ID: m.ids[7], Component: h},
	)
	return entity, (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4])), (*F)(m.world.Get(entity, m.ids[5])), (*G)(m.world.Get(entity, m.ids[6])), (*H)(m.world.Get(entity, m.ids[7]))
}

// Add the Mutate8's components to the given entity.
//
// See also [ecs.World.Add].
func (m *Mutate8[A, B, C, D, E, F, G, H]) Add(entity ecs.Entity) (*A, *B, *C, *D, *E, *F, *G, *H) {
	m.world.Add(entity, m.ids...)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4])), (*F)(m.world.Get(entity, m.ids[5])), (*G)(m.world.Get(entity, m.ids[6])), (*H)(m.world.Get(entity, m.ids[7]))
}

// Assign the Mutate8's components to the given entity, using the supplied values.
//
// See also [ecs.World.Assign] and [ecs.World.AssignN].
func (m *Mutate8[A, B, C, D, E, F, G, H]) Assign(entity ecs.Entity, a *A, b *B, c *C, d *D, e *E, f *F, g *G, h *H) (*A, *B, *C, *D, *E, *F, *G, *H) {
	m.world.AssignN(entity,
		ecs.Component{ID: m.ids[0], Component: a},
		ecs.Component{ID: m.ids[1], Component: b},
		ecs.Component{ID: m.ids[2], Component: c},
		ecs.Component{ID: m.ids[3], Component: d},
		ecs.Component{ID: m.ids[4], Component: e},
		ecs.Component{ID: m.ids[5], Component: f},
		ecs.Component{ID: m.ids[6], Component: g},
		ecs.Component{ID: m.ids[7], Component: h},
	)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4])), (*F)(m.world.Get(entity, m.ids[5])), (*G)(m.world.Get(entity, m.ids[6])), (*H)(m.world.Get(entity, m.ids[7]))
}

// Remove the Mutate8's components from the given entity.
//
// See also [ecs.World.Remove].
func (m *Mutate8[A, B, C, D, E, F, G, H]) Remove(entity ecs.Entity) {
	m.world.Remove(entity, m.ids...)
}

// Exchange components on an entity.
//
// Removes the components set via [Mutate8.WithExchange].
// Adds Mutate8's components as with [Mutate8.Add].
//
// See also [ecs.World.Exchange].
func (m *Mutate8[A, B, C, D, E, F, G, H]) Exchange(entity ecs.Entity) (*A, *B, *C, *D, *E, *F, *G, *H) {
	m.world.Exchange(entity, m.ids, m.exchange)
	return (*A)(m.world.Get(entity, m.ids[0])), (*B)(m.world.Get(entity, m.ids[1])), (*C)(m.world.Get(entity, m.ids[2])), (*D)(m.world.Get(entity, m.ids[3])), (*E)(m.world.Get(entity, m.ids[4])), (*F)(m.world.Get(entity, m.ids[5])), (*G)(m.world.Get(entity, m.ids[6])), (*H)(m.world.Get(entity, m.ids[7]))
}
