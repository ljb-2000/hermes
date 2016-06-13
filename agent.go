package main

// Agent is an asynchronous activity which has a state, and may generate events
type Agent interface {
	State() interface{}
}
