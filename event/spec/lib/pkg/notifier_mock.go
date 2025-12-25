package event

type (
	// NotifierMock is provided to help tests
	NotifierMock[E any] struct {
		OnNotify     func(E)
		OnRegister   func(Listener[E])
		OnUnregister func(Listener[E])
		OnClose      func() error
	}
)

func (n NotifierMock[E]) Notify(e E) {
	if n.OnNotify != nil {
		n.OnNotify(e)
	}
}

func (n NotifierMock[E]) Close() error {
	if n.OnNotify != nil {
		return n.OnClose()
	}
	return nil
}

func (n NotifierMock[E]) Register(l Listener[E]) {
	if n.OnRegister != nil {
		n.OnRegister(l)
	}
}

func (n NotifierMock[E]) Unregister(l Listener[E]) {
	if n.OnUnregister != nil {
		n.OnUnregister(l)
	}
}
