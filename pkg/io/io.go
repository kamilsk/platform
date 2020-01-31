package io

import "io"

// TeeReadCloser returns a ReadCloser that writes to w what it reads from rc.
// All reads from rc performed through it are matched with
// corresponding writes to w. There is no internal buffering -
// the write must complete before the read completes.
// Any error encountered while writing is reported as a read error.
//
// Deprecated: use go.octolab.org/io.TeeReadCloser instead.
func TeeReadCloser(rc io.ReadCloser, w io.Writer) io.ReadCloser {
	type pipe struct {
		io.Reader
		io.Closer
	}
	return pipe{io.TeeReader(rc, w), rc}
}
