package cobra_test

import (
	"bytes"
	"testing"

	. "github.com/kamilsk/platform/cmd/cobra"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestCompletionCommand(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		expected string
	}{
		{"Bash", "bash", "# bash completion for test"},
		{"Zsh", "zsh", "#compdef test"},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			cmd := &cobra.Command{Use: "test"}
			cmd.AddCommand(NewCompletionCommand())
			cmd.SetArgs([]string{"completion", tc.format})
			cmd.SetOutput(buf)
			assert.Len(t, cmd.Commands(), 1)
			assert.NoError(t, cmd.Execute())
			assert.Contains(t, buf.String(), tc.expected)
		})
	}
}
