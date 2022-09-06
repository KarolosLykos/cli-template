package cmd_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KarolosLykos/cli-template/cmd"
)

func TestRootCommandOutput(t *testing.T) {
	command := cmd.NewRootCmd()
	b := bytes.NewBufferString("")

	command.SetArgs([]string{"-h"})
	command.SetOut(b)

	cmdErr := command.Execute()
	require.NoError(t, cmdErr)

	out, err := ioutil.ReadAll(b)
	require.NoError(t, err)

	assert.Equal(t, "This is a template CLI application, which can be used as a boilerplate for awesome CLI tools written in Go.\n\n"+command.UsageString(), string(out))
	assert.Nil(t, cmdErr)
}
