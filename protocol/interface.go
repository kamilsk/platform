package protocol

import "context"

// Interface defines the generic behavior of all protocol listeners.
type Interface interface {
	// Listen starts listening to its protocol.
	// It also listens to Context's Done channel to try to do
	// a graceful shutdown when it closes.
	Listen(context.Context) error
}
