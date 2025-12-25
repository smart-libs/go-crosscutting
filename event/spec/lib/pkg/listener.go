package event

type Listener[E any] interface {
	OnEvent(E)
}
