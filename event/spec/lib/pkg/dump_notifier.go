package event

type (
	DumbNotifier[E any]        struct{}
	DumbNotifierFactory[E any] struct{}
)

func (d DumbNotifierFactory[E]) Create(_ ...NotifierFactoryOptions) Notifier[E] {
	return DumbNotifier[E]{}
}

func (d DumbNotifier[E]) Close() error { return nil /* Do nothing */ }

func (d DumbNotifier[E]) Notify(_ E) { /* Do nothing */ }

func (d DumbNotifier[E]) Register(_ Listener[E]) { /* Do nothing */ }

func (d DumbNotifier[E]) Unregister(_ Listener[E]) { /* Do nothing */ }

var (
	_ Notifier[string]        = DumbNotifier[string]{}
	_ NotifierFactory[string] = DumbNotifierFactory[string]{}
)
