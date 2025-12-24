package types

type (
	IndexedByStringLess    struct{}
	IndexedByString[V any] struct {
		IndexedBy[string, V, IndexedByStringLess]
	}
)

func (i2 IndexedByStringLess) Less(i, j string) bool { return i < j }

func AsIndexedByString[V any](indexed map[string]V) IndexedBy[string, V, IndexedByStringLess] {
	return indexed
}
