package cmd_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/KarolosLykos/cli-template/cmd"
)

func TestDateCmd(t *testing.T) {
	root := cmd.RootCmd

	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs([]string{"date"})

	if err := root.Execute(); err != nil {
		panic(err)
	}

	assert.Equal(t, time.Now().Format("02 Jan 06"), buf.String())
}
