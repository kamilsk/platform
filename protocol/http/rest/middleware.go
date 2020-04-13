package rest

import "net/http"

// Middleware provides functionality to compose http handlers.
//
// Deprecated: use go.octolab.org/toolkit/protocol/http/middleware.Type instead.
type Middleware func(http.Handler) http.Handler
