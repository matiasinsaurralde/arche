package iterate

import (
	"testing"

	c "github.com/mlange-42/arche/benchmark/arche/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/filter"
	"github.com/mlange-42/arche/generic"
)

func runArcheIter(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()

	posID := ecs.ComponentID[c.Position](&world)
	rotID := ecs.ComponentID[c.Rotation](&world)

	ecs.NewBuilder(&world, posID, rotID).NewBatch(count)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		query := world.Query(ecs.All(posID, rotID))
		cnt := 0
		b.StartTimer()
		for query.Next() {
			cnt++
		}
		_ = cnt
	}
}

func runArcheGet(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()

	posID := ecs.ComponentID[c.Position](&world)
	rotID := ecs.ComponentID[c.Rotation](&world)

	ecs.NewBuilder(&world, posID, rotID).NewBatch(count)

	query := world.Query(ecs.All(posID, rotID))
	for query.Next() {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			for i := 0; i < count; i++ {
				pos := (*c.Position)(query.Get(posID))
				pos.X = 1.0
			}
		}
		b.StopTimer()
		query.Close()
		if true {
			break
		}
	}
}

func runArcheGetEntity(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()

	posID := ecs.ComponentID[c.Position](&world)
	rotID := ecs.ComponentID[c.Rotation](&world)

	ecs.NewBuilder(&world, posID, rotID).NewBatch(count)

	query := world.Query(ecs.All(posID, rotID))
	for query.Next() {
		b.StartTimer()
		var e ecs.Entity
		for i := 0; i < b.N; i++ {
			for i := 0; i < count; i++ {
				e = query.Entity()
			}
		}
		b.StopTimer()
		_ = e
		query.Close()
		if true {
			break
		}
	}
}

func runArcheQuery(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()

	posID := ecs.ComponentID[c.Position](&world)
	rotID := ecs.ComponentID[c.Rotation](&world)

	ecs.NewBuilder(&world, posID, rotID).NewBatch(count)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		query := world.Query(ecs.All(posID, rotID))
		b.StartTimer()
		for query.Next() {
			pos := (*c.Position)(query.Get(posID))
			pos.X = 1.0
		}
	}
}

func runArcheQueryCached(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()

	posID := ecs.ComponentID[c.Position](&world)
	rotID := ecs.ComponentID[c.Rotation](&world)

	ecs.NewBuilder(&world, posID, rotID).NewBatch(count)

	cf := world.Cache().Register(ecs.All(posID, rotID))
	var filter ecs.Filter = &cf

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		query := world.Query(filter)
		b.StartTimer()
		for query.Next() {
			pos := (*c.Position)(query.Get(posID))
			pos.X = 1.0
		}
	}
}

func runArcheFilter(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()

	posID := ecs.ComponentID[c.Position](&world)
	rotID := ecs.ComponentID[c.Rotation](&world)

	ecs.NewBuilder(&world, posID, rotID).NewBatch(count)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		query := world.Query(filter.All(posID, rotID))
		b.StartTimer()
		for query.Next() {
			pos := (*c.Position)(query.Get(posID))
			pos.X = 1.0
		}
	}
}

func runArcheQueryGeneric(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()

	posID := ecs.ComponentID[c.Position](&world)
	rotID := ecs.ComponentID[c.Rotation](&world)

	ecs.NewBuilder(&world, posID, rotID).NewBatch(count)

	query := generic.NewFilter1[c.Position]()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q := query.Query(&world)
		b.StartTimer()
		for q.Next() {
			pos := q.Get()
			pos.X = 1.0
		}
	}
}

func runArcheQuery5C(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()

	id0 := ecs.ComponentID[c.TestStruct0](&world)
	id1 := ecs.ComponentID[c.TestStruct1](&world)
	id2 := ecs.ComponentID[c.TestStruct2](&world)
	id3 := ecs.ComponentID[c.TestStruct3](&world)
	id4 := ecs.ComponentID[c.TestStruct4](&world)

	ecs.NewBuilder(&world, id0, id1, id2, id3, id4).NewBatch(count)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		query := world.Query(ecs.All(id0, id1, id2, id3, id4))
		b.StartTimer()
		for query.Next() {
			t1 := (*c.TestStruct0)(query.Get(id0))
			t2 := (*c.TestStruct1)(query.Get(id1))
			t3 := (*c.TestStruct2)(query.Get(id2))
			t4 := (*c.TestStruct3)(query.Get(id3))
			t5 := (*c.TestStruct4)(query.Get(id4))
			t1.Val, t2.Val, t3.Val, t4.Val, t5.Val = 1, 1, 1, 1, 1
		}
	}
}

func runArcheQueryGeneric5C(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()

	id0 := ecs.ComponentID[c.TestStruct0](&world)
	id1 := ecs.ComponentID[c.TestStruct1](&world)
	id2 := ecs.ComponentID[c.TestStruct2](&world)
	id3 := ecs.ComponentID[c.TestStruct3](&world)
	id4 := ecs.ComponentID[c.TestStruct4](&world)

	ecs.NewBuilder(&world, id0, id1, id2, id3, id4).NewBatch(count)

	query := generic.NewFilter5[c.TestStruct0, c.TestStruct1, c.TestStruct2, c.TestStruct3, c.TestStruct4]()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q := query.Query(&world)
		b.StartTimer()
		for q.Next() {
			t1, t2, t3, t4, t5 := q.Get()
			t1.Val, t2.Val, t3.Val, t4.Val, t5.Val = 1, 1, 1, 1, 1
		}
	}
}

func runArcheQuery1kArch(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()
	c.RegisterAll(&world)

	perArch := count / 1000

	for i := 0; i < 1024; i++ {
		mask := i
		add := make([]ecs.ID, 0, 11)
		add = append(add, 10)
		for j := 0; j < 10; j++ {
			id := ecs.ID(j)
			m := 1 << j
			if mask&m == m {
				add = append(add, id)
			}
		}
		for j := 0; j < perArch; j++ {
			world.NewEntity(add...)
		}
	}

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		query := world.Query(ecs.All(10))
		b.StartTimer()
		for query.Next() {
			pos := (*c.TestStruct10)(query.Get(10))
			pos.Val = 1
		}
	}
}

func runArcheQuery1kArchCached(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()
	c.RegisterAll(&world)

	perArch := count / 1000

	for i := 0; i < 1024; i++ {
		mask := i
		add := make([]ecs.ID, 0, 11)
		add = append(add, 10)
		for j := 0; j < 10; j++ {
			id := ecs.ID(j)
			m := 1 << j
			if mask&m == m {
				add = append(add, id)
			}
		}
		for j := 0; j < perArch; j++ {
			world.NewEntity(add...)
		}
	}

	cf := world.Cache().Register(ecs.All(10))
	var filter ecs.Filter = &cf

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		query := world.Query(filter)
		b.StartTimer()
		for query.Next() {
			pos := (*c.TestStruct10)(query.Get(10))
			pos.Val = 1
		}
	}
}

func runArcheFilter1kArch(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()
	c.RegisterAll(&world)

	perArch := count / 1000

	for i := 0; i < 1024; i++ {
		mask := i
		add := make([]ecs.ID, 0, 11)
		add = append(add, 10)
		for j := 0; j < 10; j++ {
			id := ecs.ID(j)
			m := 1 << j
			if mask&m == m {
				add = append(add, id)
			}
		}
		for j := 0; j < perArch; j++ {
			entity := world.NewEntity()
			world.Add(entity, add...)
		}
	}

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		query := world.Query(filter.All(10))
		b.StartTimer()
		for query.Next() {
			pos := (*c.TestStruct10)(query.Get(10))
			pos.Val = 1
		}
	}
}

func runArcheQuery1Of1kArch(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()
	c.RegisterAll(&world)

	perArch := 2 * count / 1000

	for i := 0; i < 1024; i++ {
		mask := i
		add := make([]ecs.ID, 0, 10)
		for j := 0; j < 10; j++ {
			id := ecs.ID(j)
			m := 1 << j
			if mask&m == m {
				add = append(add, id)
			}
		}
		for j := 0; j < perArch; j++ {
			entity := world.NewEntity()
			world.Add(entity, add...)
		}
	}

	ecs.NewBuilder(&world, 10).NewBatch(count)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		query := world.Query(ecs.All(10))
		b.StartTimer()
		for query.Next() {
			pos := (*c.TestStruct6)(query.Get(10))
			pos.Val = 1
		}
	}
}

func runArcheQuery1Of1kArchCached(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()
	c.RegisterAll(&world)

	perArch := 2 * count / 1000

	for i := 0; i < 1024; i++ {
		mask := i
		add := make([]ecs.ID, 0, 10)
		for j := 0; j < 10; j++ {
			id := ecs.ID(j)
			m := 1 << j
			if mask&m == m {
				add = append(add, id)
			}
		}
		for j := 0; j < perArch; j++ {
			entity := world.NewEntity()
			world.Add(entity, add...)
		}
	}

	ecs.NewBuilder(&world, 10).NewBatch(count)

	cf := world.Cache().Register(ecs.All(10))
	var filter ecs.Filter = &cf

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		query := world.Query(filter)
		b.StartTimer()
		for query.Next() {
			pos := (*c.TestStruct10)(query.Get(10))
			pos.Val = 1
		}
	}
}

func runArcheQuery1kTargets(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()
	posID := ecs.ComponentID[c.TestStruct0](&world)
	relID := ecs.ComponentID[c.ChildOf](&world)

	perArch := count / 1000

	builder := ecs.NewBuilder(&world)
	targetQuery := builder.NewQuery(1000)
	targets := make([]ecs.Entity, 0, 1000)
	for targetQuery.Next() {
		targets = append(targets, targetQuery.Entity())
	}

	childBuilder := ecs.NewBuilder(&world, posID, relID).WithRelation(relID)
	for _, target := range targets {
		childBuilder.NewBatch(perArch, target)
	}

	filter := ecs.All(posID, relID)

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		query := world.Query(filter)
		b.StartTimer()
		for query.Next() {
			pos := (*c.TestStruct0)(query.Get(posID))
			pos.Val = 1
		}
	}
}

func runArcheQuery1kTargetsCached(b *testing.B, count int) {
	b.StopTimer()
	world := ecs.NewWorld()
	posID := ecs.ComponentID[c.TestStruct0](&world)
	relID := ecs.ComponentID[c.ChildOf](&world)

	perArch := count / 1000

	builder := ecs.NewBuilder(&world)
	targetQuery := builder.NewQuery(1000)
	targets := make([]ecs.Entity, 0, 1000)
	for targetQuery.Next() {
		targets = append(targets, targetQuery.Entity())
	}

	childBuilder := ecs.NewBuilder(&world, posID, relID).WithRelation(relID)
	for _, target := range targets {
		childBuilder.NewBatch(perArch, target)
	}

	cf := world.Cache().Register(ecs.All(posID, relID))
	var filter ecs.Filter = &cf

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		query := world.Query(filter)
		b.StartTimer()
		for query.Next() {
			pos := (*c.TestStruct0)(query.Get(posID))
			pos.Val = 1
		}
	}
}

func BenchmarkArcheIter_1_000(b *testing.B) {
	runArcheIter(b, 1000)
}

func BenchmarkArcheIter_10_000(b *testing.B) {
	runArcheIter(b, 10000)
}

func BenchmarkArcheIter_100_000(b *testing.B) {
	runArcheIter(b, 100000)
}

func BenchmarkArcheGet_1_000(b *testing.B) {
	runArcheGet(b, 1000)
}

func BenchmarkArcheGet_10_000(b *testing.B) {
	runArcheGet(b, 10000)
}

func BenchmarkArcheGet_100_000(b *testing.B) {
	runArcheGet(b, 100000)
}

func BenchmarkArcheGetEntity_1_000(b *testing.B) {
	runArcheGetEntity(b, 1000)
}

func BenchmarkArcheGetEntity_10_000(b *testing.B) {
	runArcheGetEntity(b, 10000)
}

func BenchmarkArcheGetEntity_100_000(b *testing.B) {
	runArcheGetEntity(b, 100000)
}

func BenchmarkArcheIterQueryID_1_000(b *testing.B) {
	runArcheQuery(b, 1000)
}

func BenchmarkArcheIterQueryID_10_000(b *testing.B) {
	runArcheQuery(b, 10000)
}

func BenchmarkArcheIterQueryID_100_000(b *testing.B) {
	runArcheQuery(b, 100000)
}

func BenchmarkArcheIterQueryIDCached_1_000(b *testing.B) {
	runArcheQueryCached(b, 1000)
}

func BenchmarkArcheIterQueryIDCached_10_000(b *testing.B) {
	runArcheQueryCached(b, 10000)
}

func BenchmarkArcheIterQueryIDCached_100_000(b *testing.B) {
	runArcheQueryCached(b, 100000)
}

func BenchmarkArcheIterFilter_1_000(b *testing.B) {
	runArcheFilter(b, 1000)
}

func BenchmarkArcheIterFilter_10_000(b *testing.B) {
	runArcheFilter(b, 10000)
}

func BenchmarkArcheIterFilter_100_000(b *testing.B) {
	runArcheFilter(b, 100000)
}

func BenchmarkArcheIterQueryGeneric_1_000(b *testing.B) {
	runArcheQueryGeneric(b, 1000)
}

func BenchmarkArcheIterQueryGeneric_10_000(b *testing.B) {
	runArcheQueryGeneric(b, 10000)
}

func BenchmarkArcheIterQueryGeneric_100_000(b *testing.B) {
	runArcheQueryGeneric(b, 100000)
}

func BenchmarkArcheIterQueryID_5C_1_000(b *testing.B) {
	runArcheQuery5C(b, 1000)
}

func BenchmarkArcheIterQueryID_5C_10_000(b *testing.B) {
	runArcheQuery5C(b, 10000)
}

func BenchmarkArcheIterQueryID_5C_100_000(b *testing.B) {
	runArcheQuery5C(b, 100000)
}

func BenchmarkArcheIterQueryGeneric_5C_1_000(b *testing.B) {
	runArcheQueryGeneric5C(b, 1000)
}

func BenchmarkArcheIterQueryGeneric_5C_10_000(b *testing.B) {
	runArcheQueryGeneric5C(b, 10000)
}

func BenchmarkArcheIterQueryGeneric_5C_100_000(b *testing.B) {
	runArcheQueryGeneric5C(b, 100000)
}

func BenchmarkArcheIter1kArchID_1_000(b *testing.B) {
	runArcheQuery1kArch(b, 1000)
}

func BenchmarkArcheIter1kArchID_10_000(b *testing.B) {
	runArcheQuery1kArch(b, 10000)
}

func BenchmarkArcheIter1kArchID_100_000(b *testing.B) {
	runArcheQuery1kArch(b, 100000)
}

func BenchmarkArcheIter1kArchIDCached_1_000(b *testing.B) {
	runArcheQuery1kArchCached(b, 1000)
}

func BenchmarkArcheIter1kArchIDCached_10_000(b *testing.B) {
	runArcheQuery1kArchCached(b, 10000)
}

func BenchmarkArcheIter1kArchIDCached_100_000(b *testing.B) {
	runArcheQuery1kArchCached(b, 100000)
}

func BenchmarkArcheFilter1kArchID_1_000(b *testing.B) {
	runArcheFilter1kArch(b, 1000)
}

func BenchmarkArcheFilter1kArchID_10_000(b *testing.B) {
	runArcheFilter1kArch(b, 10000)
}

func BenchmarkArcheFilter1kArchID_100_000(b *testing.B) {
	runArcheFilter1kArch(b, 100000)
}

func BenchmarkArcheIter1kTargets_1_000(b *testing.B) {
	runArcheQuery1kTargets(b, 1000)
}

func BenchmarkArcheIter1kTargets_10_000(b *testing.B) {
	runArcheQuery1kTargets(b, 10000)
}

func BenchmarkArcheIter1kTargets_100_000(b *testing.B) {
	runArcheQuery1kTargets(b, 100000)
}

func BenchmarkArcheIter1kTargetsCached_1_000(b *testing.B) {
	runArcheQuery1kTargetsCached(b, 1000)
}

func BenchmarkArcheIter1kTargetsCached_10_000(b *testing.B) {
	runArcheQuery1kTargetsCached(b, 10000)
}

func BenchmarkArcheIter1kTargetsCached_100_000(b *testing.B) {
	runArcheQuery1kTargetsCached(b, 100000)
}

func BenchmarkArcheIter1Of1kArch_1_000(b *testing.B) {
	runArcheQuery1Of1kArch(b, 1000)
}

func BenchmarkArcheIter1Of1kArch_10_000(b *testing.B) {
	runArcheQuery1Of1kArch(b, 10000)
}

func BenchmarkArcheIter1Of1kArch_100_000(b *testing.B) {
	runArcheQuery1Of1kArch(b, 100000)
}

func BenchmarkArcheIter1Of1kArchCached_1_000(b *testing.B) {
	runArcheQuery1Of1kArchCached(b, 1000)
}

func BenchmarkArcheIter1Of1kArchCached_10_000(b *testing.B) {
	runArcheQuery1Of1kArchCached(b, 10000)
}

func BenchmarkArcheIter1Of1kArchCached_100_000(b *testing.B) {
	runArcheQuery1Of1kArchCached(b, 100000)
}
