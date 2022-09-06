package cmd_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KarolosLykos/cli-template/cmd"
)

func TestVersionCommand(t *testing.T) {
	command := cmd.NewRootCmd()

	b := bytes.NewBufferString("")

	command.SetArgs([]string{"version"})
	command.SetOut(b)

	_, err := command.ExecuteC()
	require.NoError(t, err)

	out, err := ioutil.ReadAll(b)
	require.NoError(t, err)

	assert.Equal(t, fmt.Sprintln(command.Version), string(out))
}
