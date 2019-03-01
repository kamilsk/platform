// +build go1.8

package protocol_test

import (
	"net/http"

	. "github.com/kamilsk/platform/protocol"
)

var _ Server = &http.Server{}
