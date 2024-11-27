package handlers

import (
	"bytes"
	"strings"
	"testing"
)

func TestCSVHandler_withAddition(t *testing.T) {
	rawInput := "1,+,2"
	var output bytes.Buffer
	var logs bytes.Buffer
	handler := NewCSVHandler(strings.NewReader(rawInput), &output, &logs, calculators)
	err := handler.Handle()
	assertError(t, err, nil)
	assertEquals(t, output.String(), "1,+,2,3\n")
}

func TestCSVHandler_MultipleArguments(t *testing.T) {
	rawInput := "1,+,2\n5,*,5\n2,-,2\n2,/,2"
	var output bytes.Buffer
	var logs bytes.Buffer
	handler := NewCSVHandler(strings.NewReader(rawInput), &output, &logs, calculators)
	err := handler.Handle()
	assertError(t, err, nil)
	assertEquals(t, output.String(), "1,+,2,3\n5,*,5,25\n2,-,2,0\n2,/,2,1\n")
}

func TestCSVHandler_WithBadArguments(t *testing.T) {
	rawInput := "1,+,2\n5,*,5\nNAN,-,2\n6,NaN,3\n5,+,NaN"
	var output bytes.Buffer
	var logs bytes.Buffer
	handler := NewCSVHandler(strings.NewReader(rawInput), &output, &logs, calculators)
	err := handler.Handle()
	assertError(t, err, nil)
	assertEquals(t, output.String(), "1,+,2,3\n5,*,5,25\n")
}

func TestCSVHandler_WithDivision(t *testing.T) {
	rawInput := "2,/,2"
	var output bytes.Buffer
	var logs bytes.Buffer
	handler := NewCSVHandler(strings.NewReader(rawInput), &output, &logs, calculators)
	err := handler.Handle()
	assertError(t, err, nil)
	assertEquals(t, output.String(), "2,/,2,1\n")
}
