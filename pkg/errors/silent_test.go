package errors_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	. "github.com/kamilsk/platform/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestDoSilent(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	DoSilent(fmt.Fprintf(buf, "test"))
	assert.Equal(t, "test", buf.String())

	to, from := bytes.NewBuffer(nil), strings.NewReader("test")
	DoSilent(io.Copy(to, from))
	assert.Equal(t, "test", to.String())
}
