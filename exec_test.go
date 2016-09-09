package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironRun(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	e := &Environ{
		dir:    "/tmp",
		env:    []string{"A=1"},
		stdout: stdout,
		stderr: stderr,
	}

	err := e.Run("/bin/echo", "hello, ae")
	if assert.NoError(t, err) {
		assert.Equal(t, "hello, ae\n", stdout.String())
		assert.Equal(t, "", stderr.String())
	}
}
