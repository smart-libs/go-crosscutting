package currency

type (
	unknownSerializableType struct{}
)

func (u unknownSerializableType) String() string { return "" }
