package event

import "io"

type (
	// Source represents any source of event that is able to accept listeners to be registered and unregistered.
	// Extracting the Source interface from the Notifier allows the components to protect the Notifier object from being
	// used incorrectly by the clients that just want to register/unregister listeners, because the client will not have
	// access to Notifier.Notify() method.
	Source[E any] interface {
		io.Closer
		Register(Listener[E])
		Unregister(Listener[E])
	}
)
