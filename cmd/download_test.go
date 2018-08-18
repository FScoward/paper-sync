package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestDownloadCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := NewRootCmd()
	cmd.SetOutput(buf)
	cmd.SetArgs([]string{"download", "-i test"})

	if err := cmd.Execute(); err != nil {
		t.Errorf("??????")
	}

	actual := buf.String()
	expect := "download called\n"

	if strings.Contains(actual, expect) {
		t.Errorf("unexpected: expect:%+v, actual:%+v", expect, actual)
	}
}
