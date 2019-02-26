package protocol

import "context"

// Interface defines the generic behavior of all protocol listeners.
type Interface interface {
	// Listen starts listening to its protocol.
	// It also listens to Context's Done channel to try to do
	// a graceful shutdown when it closes.
	Listen(context.Context) error
}

// Server represents a generic server to listen some protocol.
type Server interface {
	// ListenAndServe listens some protocol and serves it.
	ListenAndServe() error
	// Shutdown tries to do a graceful shutdown.
	Shutdown(context.Context) error
}
