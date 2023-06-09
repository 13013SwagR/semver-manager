package filter

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFilterCommand(t *testing.T) {
	tests := []struct {
		name          string
		inputArgs     []string
		expectedError bool
		expectedOut   string
	}{
		{
			name:        "Provided single version highest",
			inputArgs:   []string{"--versions", "1.2.3", "--highest"},
			expectedOut: "1.2.3",
		},
		{
			name:        "Provided multiple versions highest",
			inputArgs:   []string{"--versions", "1.2.3, 1.1.1", "--highest"},
			expectedOut: "1.2.3",
		},
		{
			name:        "Provided multiple versions highest with some bad versions",
			inputArgs:   []string{"--versions", "1.2.3, 1.1.1, bad.version", "--highest"},
			expectedOut: "1.2.3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filtercmd := NewFilterCommand()
			output, err := executeCommand(filtercmd, test.inputArgs...)

			if test.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.NoError(t, err)
			}

			assert.Equal(t, test.expectedOut, output)
		})
	}
}

func executeCommand(cmd *cobra.Command, args ...string) (string, error) {
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)
	err := cmd.Execute()
	return strings.TrimSpace(buf.String()), err
}
