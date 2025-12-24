package types

import "sort"

type (
	IndexedLessFunc[K comparable]                        interface{ Less(k1, k2 K) bool }
	IndexedByEntry[K comparable, V any]                  func() (key K, value V)
	IndexedBy[K comparable, V any, L IndexedLessFunc[K]] map[K]V
)

func newIndexedByEntry[K comparable, T any](key K, value T) IndexedByEntry[K, T] {
	return func() (K, T) { return key, value }
}

func (i IndexedBy[K, T, L]) OrderedListAsc() []IndexedByEntry[K, T] {
	var tuples []IndexedByEntry[K, T]
	for key, value := range i {
		tuples = append(tuples, newIndexedByEntry(key, value))
	}
	var less L
	sort.Slice(tuples, func(i, j int) bool {
		ki, _ := tuples[i]()
		kj, _ := tuples[j]()
		return less.Less(ki, kj)
	})
	return tuples
}
