package types

import "sort"

type (
	IndexedByDateEntry[T any] func() (key Date, value T)
	IndexedByDate[T any]      map[Date]T
)

func newIndexedByDateEntry[T any](key Date, value T) IndexedByDateEntry[T] {
	return func() (Date, T) { return key, value }
}

func (i IndexedByDate[T]) OrderedListAsc() []IndexedByDateEntry[T] {
	var tuples []IndexedByDateEntry[T]
	for key, value := range i {
		tuples = append(tuples, newIndexedByDateEntry(key, value))
	}
	sort.Slice(tuples, func(i, j int) bool {
		iDate, _ := tuples[i]()
		jDate, _ := tuples[j]()
		return iDate.Time().Before(jDate.Time())
	})
	return tuples
}
