package io

import "io"

func Tee(rc io.ReadCloser, w io.Writer) io.ReadCloser {
	type pipe struct {
		io.Reader
		io.Closer
	}
	return pipe{io.TeeReader(rc, w), rc}
}
