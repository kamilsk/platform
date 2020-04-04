package http

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	CacheControlHeader = "Cache-Control"
	XStrictHeader      = "X-Strict"
	XTimeoutHeader     = "X-Timeout"
)

// Header extends built-in http.Header.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http.Header instead.
type Header http.Header

// NoCache returns true if the header has no-cache value of cache control.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http.Header instead.
func (header Header) NoCache() bool {
	return strings.EqualFold(http.Header(header).Get(CacheControlHeader), "no-cache")
}

// Strict returns true if the header has this value.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http.Header instead.
func (header Header) Strict() bool {
	var strict bool
	if v := http.Header(header).Get(XStrictHeader); v != "" {
		strict, _ = strconv.ParseBool(v)
	}
	return strict
}

// Timeout returns the percentage of a timeout value from the header
// or the fallback value.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http.Header instead.
func (header Header) Timeout(fallback time.Duration, percentage float64) time.Duration {
	d, err := time.ParseDuration(http.Header(header).Get(XTimeoutHeader))
	if err != nil {
		d = fallback
	}
	return time.Duration(percentage * float64(d))
}
