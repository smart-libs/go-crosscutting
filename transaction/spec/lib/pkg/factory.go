package transaction

type (
	// FactoryOptions is a markup interface where specific options can be provided to a specific transaction
	// implementation
	FactoryOptions interface{}
	Factory        interface {
		Create(options ...FactoryOptions) (Transaction, error)
	}
)
