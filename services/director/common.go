package director

// Common interfaces for pipe.
type (
	Task  chan interface{}
	Queue chan chan interface{}
)
