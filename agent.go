package main

// Event it the central unit of data in Hermes. It represents something
// that has happened somewhere, and may include data about how that something
// occured. For now it is stubbed out.
type Event bool

// Agent is an asynchronous activity which has a state, and may generate events
type Agent interface {
	State() interface{}
	Run()
	Events() chan Event
}
