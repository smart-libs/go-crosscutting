package event

type (
	OrdinaryNotifier[E any] struct {
		listeners []Listener[E]
	}
)

func (o *OrdinaryNotifier[E]) Notify(e E) {
	for _, l := range o.listeners {
		l.OnEvent(e)
	}
}

func (o *OrdinaryNotifier[E]) Close() error {
	o.listeners = nil
	return nil
}

func (o *OrdinaryNotifier[E]) Register(l Listener[E]) {
	o.listeners = append(o.listeners, l)
}

func (o *OrdinaryNotifier[E]) Unregister(l Listener[E]) {
	for i, elem := range o.listeners {
		if elem == l {
			o.listeners = append(o.listeners[:i], o.listeners[i+1:]...)
			return
		}
	}
}
