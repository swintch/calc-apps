package handlers

import (
	"bytes"
	"os"
	"strconv"
	"testing"
)

func TestCLIHandler_InvalidNumberOfArguments(t *testing.T) {
	handle := NewCLIHandler("+", os.Stdout)
	err := handle.Handler([]string{"1"})
	AssertError(t, err, InvalidNumberOfArguments)
}

func TestCLIHandler_InvalidFirstArgument(t *testing.T) {
	handle := NewCLIHandler("+", os.Stdout)
	err := handle.Handler([]string{"a", "2"})
	AssertError(t, err, InvalidArgumentFormat)
	AssertError(t, err, strconv.ErrSyntax)
}

func TestCLIHandler_InvalidSecondArgument(t *testing.T) {
	handle := NewCLIHandler("+", os.Stdout)
	err := handle.Handler([]string{"1", "b"})
	AssertError(t, err, InvalidArgumentFormat)
	AssertError(t, err, strconv.ErrSyntax)
}

func TestCLIHandler_TestOutputToConsole(t *testing.T) {
	output := bytes.Buffer{}
	handle := NewCLIHandler("+", &output)
	err := handle.Handler([]string{"1", "2"})
	AssertError(t, err, nil)
	AssertEquals(t, output.String(), "3")

}

func TestCLIHandler_TestOutputToConsoleError(t *testing.T) {
	output := &WriterError{err: boink}
	handle := NewCLIHandler("+", output)
	err := handle.Handler([]string{"1", "2"})
	AssertError(t, err, boink)
	AssertError(t, err, ConsoleWriteError)

}

func TestCLIHandler_TestBadOperator(t *testing.T) {
	output := bytes.Buffer{}
	handle := NewCLIHandler("adff", &output)
	err := handle.Handler([]string{"1", "2"})
	AssertError(t, err, InvalidOperator)
}
