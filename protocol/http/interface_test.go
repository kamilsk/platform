package http_test

import (
	"net/http"

	. "github.com/kamilsk/platform/protocol/http"
)

var _ Router = http.NewServeMux()
