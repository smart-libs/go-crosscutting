package event

type (
	ListenerMock[E any] struct {
		OnEventFunc func(E)
	}
)

func (l ListenerMock[E]) OnEvent(e E) {
	if l.OnEventFunc != nil {
		l.OnEventFunc(e)
	}
}
