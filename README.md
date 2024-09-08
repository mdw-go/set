# github.com/mdwhatcott/go-set/v2


	package set // import "github.com/mdwhatcott/go-set/v2/set"
	
	Package set implements a generic set type. Finally!
	https://en.wikipedia.org/wiki/Set_(mathematics)
	
	TYPES
	
	type Set[T comparable] map[T]struct{}
	
	func FromSeq[T comparable](seq iter.Seq[T]) (result Set[T])
	
	func Make[T comparable](size int) Set[T]
	
	func Of[T comparable](items ...T) (result Set[T])
	
	func (s Set[T]) Add(items ...T) Set[T]
	
	func (s Set[T]) All() iter.Seq[T]
	
	func (s Set[T]) Clear() Set[T]
	
	func (s Set[T]) Contains(item T) bool
	
	func (s Set[T]) Difference(that Set[T]) (result Set[T])
	
	func (s Set[T]) Empty() bool
	
	func (s Set[T]) Equal(that Set[T]) bool
	
	func (s Set[T]) Intersection(that Set[T]) (result Set[T])
	
	func (s Set[T]) IsSubset(that Set[T]) bool
	
	func (s Set[T]) IsSuperset(that Set[T]) bool
	
	func (s Set[T]) Len() int
	
	func (s Set[T]) Remove(items ...T) Set[T]
	
	func (s Set[T]) Slice() (result []T)
	
	func (s Set[T]) SymmetricDifference(that Set[T]) (result Set[T])
	
	func (s Set[T]) Union(that Set[T]) (result Set[T])
	
