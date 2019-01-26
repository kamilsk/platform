package cobra_test

import (
	"bytes"
	"testing"

	. "github.com/kamilsk/platform/cmd/cobra"
	"github.com/stretchr/testify/assert"
)

func TestVersionCommand(t *testing.T) {
	tests := []struct {
		name                  string
		commit, date, release string
		expected              string
	}{
		{"stable version", "050fb2f", "Fri Jan 25", "1.0.0", "Version 1.0.0 (commit: 050fb2f, build date: Fri Jan 25"},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			cmd := NewVersionCommand(tc.commit, tc.date, tc.release)
			cmd.SetOutput(buf)
			cmd.Run(cmd, nil)
			assert.Contains(t, buf.String(), tc.expected)
		})
	}
}
