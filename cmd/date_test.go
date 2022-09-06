package cmd_test

import (
	"bytes"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/KarolosLykos/cli-template/cmd"
)

var testCases = []struct {
	name     string
	args     []string
	expected string
	err      bool
}{
	{
		name:     "default",
		args:     nil,
		expected: "02 Jan 06",
		err:      false,
	},
	{
		name:     "02 01 06",
		args:     []string{"-f", "02 Jan 06"},
		expected: "02 Jan 06",
		err:      false,
	},
	{
		name:     "02 01 06 15:04:05",
		args:     []string{"-f", "02 01 06 15:04:05"},
		expected: "02 01 06 15:04:05",
		err:      false,
	},
	{
		name:     "02 Jan 06",
		args:     []string{"-f", "02 Jan 06"},
		expected: "02 Jan 06",
		err:      false,
	},
	{
		name:     "02 Jan 06 15:04:05",
		args:     []string{"-f", "02 Jan 06 15:04:05"},
		expected: "02 Jan 06 15:04:05",
		err:      false,
	},
	{
		name:     "Mon Jan 06",
		args:     []string{"-f", "Mon Jan 06"},
		expected: "Mon Jan 06",
		err:      false,
	},
	{
		name:     "Mon Jan 06 15:04:05",
		args:     []string{"-f", "Mon Jan 06 15:04:05"},
		expected: "Mon Jan 06 15:04:05",
		err:      false,
	},
	{
		name:     "02 01 2006",
		args:     []string{"-f", "02 01 2006"},
		expected: "02 01 2006",
		err:      false,
	},
	{
		name:     "02 01 2006 15:04:05 MST",
		args:     []string{"-f", "02 01 2006 15:04:05 MST"},
		expected: "02 01 2006 15:04:05 MST",
		err:      false,
	},
	{
		name:     "Wrong format",
		args:     []string{"-f", "Something wrong"},
		expected: "",
		err:      true,
	},
}

func TestDateCommand(t *testing.T) {
	for _, tc := range testCases {
		command := cmd.NewDateCmd()
		b := bytes.NewBufferString("")

		command.SetArgs(tc.args)
		command.SetOut(b)

		err := command.Execute()
		out, _ := ioutil.ReadAll(b)

		if tc.err {
			assert.Error(t, err, tc.name)
		} else {
			assert.Equal(t, time.Now().Format(tc.expected)+"\n", string(out), tc.name)
		}
	}
}
