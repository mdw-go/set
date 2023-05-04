package set_test

import (
	"sort"
	"testing"

	"github.com/mdwhatcott/go-set/internal/should"
	"github.com/mdwhatcott/go-set/set"
)

func TestCreation(t *testing.T) {
	should.So(t, len(set.New[int](0)), should.Equal, 0)
	should.So(t, len(set.From[int]()), should.Equal, 0)
	should.So(t, set.From[int]().Len(), should.Equal, 0)
	should.So(t, set.From[int]().Empty(), should.BeTrue)
	should.So(t, set.From[int](1).Empty(), should.BeFalse)
}
func TestContains(t *testing.T) {
	should.So(t, set.From[int](1).Contains(1), should.BeTrue)
	should.So(t, set.From[int]().Contains(1), should.BeFalse)
}
func TestAdd(t *testing.T) {
	values := set.New[int](0)
	values.Add(1, 2, 3)
	should.So(t, values.Contains(1), should.BeTrue)
	should.So(t, values.Contains(2), should.BeTrue)
	should.So(t, values.Contains(3), should.BeTrue)
	should.So(t, values.Len(), should.Equal, 3)
}
func TestRemove(t *testing.T) {
	values := set.From[int](1, 2, 3)
	values.Remove(2)
	should.So(t, values.Contains(1), should.BeTrue)
	should.So(t, values.Contains(2), should.BeFalse)
	should.So(t, values.Contains(3), should.BeTrue)
	should.So(t, values.Len(), should.Equal, 2)
}
func TestClear(t *testing.T) {
	values := set.From[int](1, 2, 3)
	values.Clear()
	should.So(t, values.Contains(1), should.BeFalse)
	should.So(t, values.Contains(2), should.BeFalse)
	should.So(t, values.Contains(3), should.BeFalse)
	should.So(t, values.Len(), should.Equal, 0)
}
func TestSlice(t *testing.T) {
	values := set.From[int](1, 2, 3, 4, 5)
	items := values.Slice()
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})
	should.So(t, values.Len(), should.Equal, 5)
	should.So(t, len(items), should.Equal, 5)
	should.So(t, items, should.Equal, []int{1, 2, 3, 4, 5})
}
func TestEqual(t *testing.T) {
	should.So(t, set.From[int](1, 2, 3).Equal(set.From[int](3, 2, 1)), should.BeTrue)
	should.So(t, set.From[int](1, 2).Equal(set.From[int](3, 2, 1)), should.BeFalse)
	should.So(t, set.From[int](1, 2, 2).Equal(set.From[int](1, 2, 3)), should.BeFalse)
	should.So(t, set.From[int](1, 2, 3).Equal(set.From[int](1, 2, 4)), should.BeFalse)
}
func TestIsSubset(t *testing.T) {
	should.So(t, set.From[int](1, 2, 3).IsSubset(set.From[int](1, 2, 3, 4, 5)), should.BeTrue)
	should.So(t, set.From[int](4, 5, 6).IsSubset(set.From[int](1, 2, 3, 4, 5)), should.BeFalse)
}
func TestIsSuperset(t *testing.T) {
	should.So(t, set.From[int](1, 2, 3, 4, 5).IsSuperset(set.From[int](1, 2, 3)), should.BeTrue)
	should.So(t, set.From[int](1, 2, 3, 4, 5).IsSuperset(set.From[int](4, 5, 6)), should.BeFalse)
}
func TestUnion(t *testing.T) {
	should.So(t, set.From[int](1, 2, 3).Union(set.From[int](1, 2, 3)), should.Equal, set.From[int](1, 2, 3))
	should.So(t, set.From[int](1, 2, 3).Union(set.From[int](2, 3, 4)), should.Equal, set.From[int](1, 2, 3, 4))
	should.So(t, set.From[int](1, 2, 3).Union(set.From[int](4, 5, 6)), should.Equal, set.From[int](1, 2, 3, 4, 5, 6))
}
func TestIntersection(t *testing.T) {
	should.So(t, set.From[int](1, 2, 3).Intersection(set.From[int](4, 5, 6)), should.Equal, set.From[int]())
	should.So(t, set.From[int](1, 2, 3).Intersection(set.From[int](2, 3, 4)), should.Equal, set.From[int](2, 3))
}
func TestDifference(t *testing.T) {
	should.So(t, set.From[int](1, 2, 3).Difference(set.From[int](4, 5, 6)), should.Equal, set.From[int](1, 2, 3))
	should.So(t, set.From[int](1, 2, 3).Difference(set.From[int](2, 3)), should.Equal, set.From[int](1))
}
func TestSymmetricDifference(t *testing.T) {
	should.So(t, set.From[int](1, 2, 3).SymmetricDifference(set.From[int](4, 5, 6)), should.Equal, set.From[int](1, 2, 3, 4, 5, 6))
	should.So(t, set.From[int](1, 2, 3).SymmetricDifference(set.From[int](2, 3, 4)), should.Equal, set.From[int](1, 4))
}
