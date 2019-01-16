package protocol

import "context"

// Interface defines the generic behavior of all protocol listeners.
type Interface interface {
	// Listen starts listening to its protocol.
	Listen(context.Context) error
	// Shutdown tries to do a graceful shutdown.
	Shutdown(context.Context) error
}
