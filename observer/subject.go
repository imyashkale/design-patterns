package observer

type Subject interface {
	register(o Observer)
	deregister(o Observer)
	notifyAll(o Observer)
}