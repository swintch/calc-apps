package handlers

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewCSVHandler_TestAddition(t *testing.T) {
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

func TestNewCSVHandler_MultipleValues(t *testing.T) {
	output := bytes.Buffer{}
	logs := bytes.Buffer{}
	input := "1,+,2\n2,-,2\n2,*,2\n6,/,3"
	handle := NewCSVHandler(strings.NewReader(input), &output, &logs)
	err := handle.Handle()
	AssertError(t, err, nil)
	AssertEquals(t, output.String(), "1,+,2,3\n2,-,2,0\n2,*,2,4\n6,/,3,2\n")
	if logs.Len() > 0 {
		t.Logf("\n%s", logs.String())
	}
}

func TestNewCSVHandler_MultipleBadValues(t *testing.T) {
	output := bytes.Buffer{}
	logs := bytes.Buffer{}
	input := "1,+,2\nNaN,-,2\n2,NaN,2\n6,/,NaN"
	handle := NewCSVHandler(strings.NewReader(input), &output, &logs)
	err := handle.Handle()
	AssertError(t, err, nil)
	AssertEquals(t, output.String(), "1,+,2,3\n")
	if logs.Len() > 0 {
		t.Logf("\n%s", logs.String())
	}
}

func TestNewCSVHandler_TestReadError(t *testing.T) {
	output := bytes.Buffer{}
	logs := bytes.Buffer{}
	input := ReaderError{err: boink}
	handle := NewCSVHandler(&input, &output, &logs)
	err := handle.Handle()
	AssertError(t, err, boink)
	if logs.Len() > 0 {
		t.Logf("\n%s", logs.String())
	}
}

func TestNewCSVHandler_TestWriteError(t *testing.T) {
	output := WriterError{err: boink}
	logs := bytes.Buffer{}
	input := "1,+,2"
	handle := NewCSVHandler(strings.NewReader(input), &output, &logs)
	err := handle.Handle()
	AssertError(t, err, boink)
	if logs.Len() > 0 {
		t.Logf("\n%s", logs.String())
	}
}
