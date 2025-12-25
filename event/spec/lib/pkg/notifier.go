package event

type (
	Notifier[E any] interface {
		Notify(E)
		Source[E]
	}

	// NotifierFactoryOptions is a markup interface to allow client to provide specific options to a
	// specific implementation
	NotifierFactoryOptions interface{}

	NotifierFactory[E any] interface {
		Create(options ...NotifierFactoryOptions) Notifier[E]
	}
)
