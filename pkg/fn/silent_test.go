package fn_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	. "github.com/kamilsk/platform/pkg/fn"
	"github.com/stretchr/testify/assert"
)

func TestDoSilent(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	DoSilent(fmt.Fprintf(buf, "test"))
	assert.Equal(t, "test", buf.String())
}

func TestDoSilent64(t *testing.T) {
	to, from := bytes.NewBuffer(nil), strings.NewReader("test")

	DoSilent64(io.Copy(to, from))
	assert.Equal(t, "test", to.String())
}
