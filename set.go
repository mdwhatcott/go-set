// Package set implements a generic set type. Finally!
// https://en.wikipedia.org/wiki/Set_(mathematics)
package set

import (
	"iter"
	"maps"
	"slices"
)

type Set[T comparable] map[T]struct{}

func Make[T comparable](size int) Set[T] {
	return make(Set[T], size)
}
func Of[T comparable](items ...T) (result Set[T]) {
	return Make[T](len(items)).Add(items...)
}
func FromSeq[T comparable](seq iter.Seq[T]) (result Set[T]) {
	result = make(Set[T])
	for t := range seq {
		result.Add(t)
	}
	return result
}
func (s Set[T]) Contains(item T) bool {
	_, found := s[item]
	return found
}
func (s Set[T]) Add(items ...T) Set[T] {
	return s.AddIter(slices.Values(items))
}
func (s Set[T]) AddIter(it iter.Seq[T]) Set[T] {
	for item := range it {
		s[item] = struct{}{}
	}
	return s
}
func (s Set[T]) Remove(items ...T) Set[T] {
	return s.RemoveIter(slices.Values(items))
}
func (s Set[T]) RemoveIter(it iter.Seq[T]) Set[T] {
	for item := range it {
		delete(s, item)
	}
	return s
}
func (s Set[T]) Empty() bool {
	return s.Len() == 0
}
func (s Set[T]) Len() int {
	return len(s)
}
func (s Set[T]) All() iter.Seq[T] {
	return maps.Keys(s)
}
func (s Set[T]) Slice() (result []T) {
	return slices.Collect(s.All())
}
func (s Set[T]) Clear() Set[T] {
	clear(s)
	return s
}
func (s Set[T]) Equal(that Set[T]) bool {
	if len(s) != len(that) {
		return false
	}
	for item := range s {
		if !that.Contains(item) {
			return false
		}
	}
	return true
}
func (s Set[T]) IsSubset(that Set[T]) bool {
	for item := range s {
		if !that.Contains(item) {
			return false
		}
	}
	return true
}
func (s Set[T]) IsSuperset(that Set[T]) bool {
	for item := range that {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}
func (s Set[T]) Union(that Set[T]) (result Set[T]) {
	result = make(Set[T])
	for item := range s {
		result.Add(item)
	}
	for item := range that {
		result.Add(item)
	}
	return result
}
func (s Set[T]) Intersection(that Set[T]) (result Set[T]) {
	result = make(Set[T])
	for item := range s {
		if that.Contains(item) {
			result.Add(item)
		}
	}
	return result
}
func (s Set[T]) Difference(that Set[T]) (result Set[T]) {
	result = make(Set[T])
	for item := range s {
		if !that.Contains(item) {
			result.Add(item)
		}
	}
	return result
}
func (s Set[T]) SymmetricDifference(that Set[T]) (result Set[T]) {
	return s.Difference(that).Union(that.Difference(s))
}
