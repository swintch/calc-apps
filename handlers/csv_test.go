package handlers

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewCSVHandler(t *testing.T) {
	output := bytes.Buffer{}
	logs := bytes.Buffer{}
	input := "1,+,2"
	handle := NewCSVHandler(strings.NewReader(input), &output, &logs)
	err := handle.Handle()
	AssertError(t, err, nil)
	AssertEquals(t, output.String(), "1,+,2,3\n")
	if logs.Len() > 0 {
		t.Logf("\n%s", logs.String())
	}
}
