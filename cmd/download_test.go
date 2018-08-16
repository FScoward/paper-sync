package cmd

import (
	"bytes"
	"testing"
)

func TestDownloadCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	cmd := NewRootCmd()
	cmd.SetOutput(buf)
	cmd.SetArgs([]string{"download"})

	if err := cmd.Execute(); err != nil {
		t.Errorf("??????")
	}

	actual := buf.String()
	expect := "download called\n"
	if actual != expect {
		t.Errorf("unexpected: expect:%+v, actual:%+v", expect, actual)
	}
}
