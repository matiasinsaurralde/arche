package filter

import (
	"github.com/mlange-42/arche/ecs"
)

// All matches entities that have all the given components.
//
// Like [And] for combining individual components.
func All(comps ...ecs.ID) ecs.Mask {
	return ecs.All(comps...)
}

// ANY matches entities that have any of the given components.
type ANY ecs.Mask

// Any matches entities that have any of the given components.
func Any(comps ...ecs.ID) ANY {
	return ANY(ecs.All(comps...))
}

// Matches matches a filter against a bitmask
func (f ANY) Matches(bits ecs.Mask, target *ecs.Entity) bool {
	return bits.ContainsAny(ecs.Mask(f))
}

// NoneOF matches entities that are missing all the given components.
type NoneOF ecs.Mask

// NoneOf matches entities that are missing all the given components.
func NoneOf(comps ...ecs.ID) NoneOF {
	return NoneOF(ecs.All(comps...))
}

// Matches matches a filter against a bitmask
func (f NoneOF) Matches(bits ecs.Mask, target *ecs.Entity) bool {
	return !bits.ContainsAny(ecs.Mask(f))
}

// AnyNOT matches entities that are missing any of the given components.
type AnyNOT ecs.Mask

// AnyNot matches entities that are missing any of the given components.
func AnyNot(comps ...ecs.ID) AnyNOT {
	return AnyNOT(ecs.All(comps...))
}

// Matches matches a filter against a bitmask
func (f AnyNOT) Matches(bits ecs.Mask, target *ecs.Entity) bool {
	return !bits.Contains(ecs.Mask(f))
}

// AND combines two filters using AND.
//
// Ignores potential relation targets.
type AND struct {
	L ecs.Filter
	R ecs.Filter
}

// And combines two filters using AND.
//
// Ignores potential relation targets.
func And(l, r ecs.Filter) *AND {
	return &AND{L: l, R: r}
}

// Matches matches a filter against a bitmask.
func (f *AND) Matches(bits ecs.Mask, target *ecs.Entity) bool {
	return f.L.Matches(bits, nil) && f.R.Matches(bits, nil)
}

// OR combines two filters using OR.
//
// Ignores potential relation targets.
type OR struct {
	L ecs.Filter
	R ecs.Filter
}

// Or combines two filters using OR.
//
// Ignores potential relation targets.
func Or(l, r ecs.Filter) *OR {
	return &OR{L: l, R: r}
}

// Matches matches a filter against a bitmask.
func (f *OR) Matches(bits ecs.Mask, target *ecs.Entity) bool {
	return f.L.Matches(bits, nil) || f.R.Matches(bits, nil)
}

// XOR combines two filters using XOR.
//
// Ignores potential relation targets.
type XOR struct {
	L ecs.Filter
	R ecs.Filter
}

// XOr combines two filters using XOR.
//
// Ignores potential relation targets.
func XOr(l, r ecs.Filter) *XOR {
	return &XOR{L: l, R: r}
}

// Matches matches a filter against a bitmask.
func (f *XOR) Matches(bits ecs.Mask, target *ecs.Entity) bool {
	return f.L.Matches(bits, nil) != f.R.Matches(bits, nil)
}

// NOT inverts a filter. It matches if the inner filter does not.
//
// Does NOT ignore a potential relation target.
type NOT struct {
	F ecs.Filter
}

// Not inverts a filter. It matches if the inner filter does not.
func Not(f ecs.Filter) *NOT {
	return &NOT{F: f}
}

// Matches matches a filter against a bitmask.
func (f *NOT) Matches(bits ecs.Mask, target *ecs.Entity) bool {
	return !f.F.Matches(bits, target)
}
