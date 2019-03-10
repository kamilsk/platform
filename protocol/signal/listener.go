package signal

import (
	"context"
	"io"

	"github.com/kamilsk/platform/pkg/safe"
	"github.com/kamilsk/platform/pkg/sync"
)

// New returns a listener of the termination signals.
func New() *Listener {
	return &Listener{}
}

// Listener listens the termination signals and holds resources
// to release their later when the termination signals caught.
type Listener struct {
	resources []resource
}

// AddResource registers the resource to release it later
// when the termination signals caught.
func (listener *Listener) AddResource(src io.Closer, cleaners ...func(error)) {
	listener.resources = append(listener.resources, resource{src, cleaners})
}

// Listen starts listening to the termination signals.
// It releases registered resources when the termination signals caught
// and never returns an error.
func (listener *Listener) Listen(ctx context.Context) error {
	if err := sync.Termination().Wait(ctx); err == sync.ErrSignalTrapped {
		for _, resource := range listener.resources {
			safe.Close(resource, resource.cleaners...)
		}
	}
	return nil
}

type resource struct {
	io.Closer
	cleaners []func(error)
}
