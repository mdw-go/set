package set_test

import (
	"slices"
	"sort"
	"testing"

	"github.com/mdwhatcott/go-set/v2/internal/should"
	"github.com/mdwhatcott/go-set/v2/set"
)

func TestCreation(t *testing.T) {
	should.So(t, len(set.Of[int]()), should.Equal, 0)
	should.So(t, len(set.Make[int](0)), should.Equal, 0)
	should.So(t, set.Of[int]().Len(), should.Equal, 0)
	should.So(t, set.Of[int]().Empty(), should.BeTrue)
	should.So(t, set.Of[int](1).Empty(), should.BeFalse)
	should.So(t, set.FromSeq(set.Of(1, 2, 3).All()), should.Equal, set.Of(1, 2, 3))
}
func TestContains(t *testing.T) {
	should.So(t, set.Of[int](1).Contains(1), should.BeTrue)
	should.So(t, set.Of[int]().Contains(1), should.BeFalse)
}
func TestAdd(t *testing.T) {
	values := set.Of[int]().Add(1, 2, 3)
	should.So(t, values.Contains(1), should.BeTrue)
	should.So(t, values.Contains(2), should.BeTrue)
	should.So(t, values.Contains(3), should.BeTrue)
	should.So(t, values.Len(), should.Equal, 3)
}
func TestRemove(t *testing.T) {
	values := set.Of[int](1, 2, 3).Remove(2)
	should.So(t, values.Contains(1), should.BeTrue)
	should.So(t, values.Contains(2), should.BeFalse)
	should.So(t, values.Contains(3), should.BeTrue)
	should.So(t, values.Len(), should.Equal, 2)
}
func TestClear(t *testing.T) {
	values := set.Of[int](1, 2, 3).Clear()
	should.So(t, values.Contains(1), should.BeFalse)
	should.So(t, values.Contains(2), should.BeFalse)
	should.So(t, values.Contains(3), should.BeFalse)
	should.So(t, values.Len(), should.Equal, 0)
}
func TestSlice(t *testing.T) {
	items := set.Of[int](1, 2, 3, 4, 5).Slice()
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})
	should.So(t, len(items), should.Equal, 5)
	should.So(t, items, should.Equal, []int{1, 2, 3, 4, 5})
}
func TestAll(t *testing.T) {
	items := slices.Collect(set.Of[int](1, 2, 3, 4, 5).All())
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})
	should.So(t, len(items), should.Equal, 5)
	should.So(t, items, should.Equal, []int{1, 2, 3, 4, 5})
}
func TestEqual(t *testing.T) {
	should.So(t, set.Of[int](1, 2, 3).Equal(set.Of[int](3, 2, 1)), should.BeTrue)
	should.So(t, set.Of[int](1, 2).Equal(set.Of[int](3, 2, 1)), should.BeFalse)
	should.So(t, set.Of[int](1, 2, 2).Equal(set.Of[int](1, 2, 3)), should.BeFalse)
	should.So(t, set.Of[int](1, 2, 3).Equal(set.Of[int](1, 2, 4)), should.BeFalse)
}
func TestIsSubset(t *testing.T) {
	should.So(t, set.Of[int](1, 2, 3).IsSubset(set.Of[int](1, 2, 3, 4, 5)), should.BeTrue)
	should.So(t, set.Of[int](4, 5, 6).IsSubset(set.Of[int](1, 2, 3, 4, 5)), should.BeFalse)
}
func TestIsSuperset(t *testing.T) {
	should.So(t, set.Of[int](1, 2, 3, 4, 5).IsSuperset(set.Of[int](1, 2, 3)), should.BeTrue)
	should.So(t, set.Of[int](1, 2, 3, 4, 5).IsSuperset(set.Of[int](4, 5, 6)), should.BeFalse)
}
func TestUnion(t *testing.T) {
	should.So(t, set.Of[int](1, 2, 3).Union(set.Of[int](1, 2, 3)), should.Equal, set.Of[int](1, 2, 3))
	should.So(t, set.Of[int](1, 2, 3).Union(set.Of[int](2, 3, 4)), should.Equal, set.Of[int](1, 2, 3, 4))
	should.So(t, set.Of[int](1, 2, 3).Union(set.Of[int](4, 5, 6)), should.Equal, set.Of[int](1, 2, 3, 4, 5, 6))
}
func TestIntersection(t *testing.T) {
	should.So(t, set.Of[int](1, 2, 3).Intersection(set.Of[int](4, 5, 6)), should.Equal, set.Of[int]())
	should.So(t, set.Of[int](1, 2, 3).Intersection(set.Of[int](2, 3, 4)), should.Equal, set.Of[int](2, 3))
}
func TestDifference(t *testing.T) {
	should.So(t, set.Of[int](1, 2, 3).Difference(set.Of[int](4, 5, 6)), should.Equal, set.Of[int](1, 2, 3))
	should.So(t, set.Of[int](4, 5, 6).Difference(set.Of[int](1, 2, 3)), should.Equal, set.Of[int](4, 5, 6))
	should.So(t, set.Of[int](1, 2, 3).Difference(set.Of[int](1, 2, 3)), should.Equal, set.Of[int]())
	should.So(t, set.Of[int](1, 2, 3).Difference(set.Of[int](2, 3)), should.Equal, set.Of[int](1))
	should.So(t, set.Of[int](1, 2, 3).Difference(set.Of[int](1)), should.Equal, set.Of[int](2, 3))
	should.So(t, set.Of[int](2, 3).Difference(set.Of[int](1, 2, 3)), should.Equal, set.Of[int]())
	should.So(t, set.Of[int](1).Difference(set.Of[int](1, 2, 3)), should.Equal, set.Of[int]())
}
func TestSymmetricDifference(t *testing.T) {
	should.So(t, set.Of[int](1, 2, 3).SymmetricDifference(set.Of[int](4, 5, 6)), should.Equal, set.Of[int](1, 2, 3, 4, 5, 6))
	should.So(t, set.Of[int](1, 2, 3).SymmetricDifference(set.Of[int](2, 3, 4)), should.Equal, set.Of[int](1, 4))
}
