package fn_test

import (
	"bytes"
	"fmt"
	"testing"

	. "github.com/kamilsk/platform/pkg/fn"
	"github.com/stretchr/testify/assert"
)

func TestSilent(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	PrintSilent(fmt.Fprintf(buf, "test"))
	assert.Equal(t, "test", buf.String())
}
