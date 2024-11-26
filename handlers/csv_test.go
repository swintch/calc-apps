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

func TestCSVHandler_MultiWithBadArguments(t *testing.T) {
	rawInput := "1,+,2\n5,*,5\nNAN,-,2\n6,NaN,3\n5,+,NaN"
	var output bytes.Buffer
	var logs bytes.Buffer
	handler := NewCSVHandler(strings.NewReader(rawInput), &output, &logs, calculators)

	err := handler.Handle()

	assertError(t, err, nil)
	assertEquals(t, output.String(), "1,+,2,3\n5,*,5,25\n")
	if logs.Len() > 0 {
		t.Logf("\n%s", logs.String())
	}
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

func TestCSVHandler_WithMultiplication(t *testing.T) {
	rawInput := "2,*,2"
	var output bytes.Buffer
	var logs bytes.Buffer
	handler := NewCSVHandler(strings.NewReader(rawInput), &output, &logs, calculators)
	err := handler.Handle()
	assertError(t, err, nil)
	assertEquals(t, output.String(), "2,*,2,4\n")
}

func TestCSVHandler_WithSubtraction(t *testing.T) {
	rawInput := "2,-,2"
	var output bytes.Buffer
	var logs bytes.Buffer
	handler := NewCSVHandler(strings.NewReader(rawInput), &output, &logs, calculators)

	err := handler.Handle()

	assertError(t, err, nil)
	assertEquals(t, output.String(), "2,-,2,0\n")
	if logs.Len() > 0 {
		t.Logf("\n%s", logs.String())
	}
}

func TestCSVHandler_ReaderError(t *testing.T) {
	rawInput := &ReaderErr{err: boink}
	var output bytes.Buffer
	var logs bytes.Buffer
	handler := NewCSVHandler(rawInput, &output, &logs, calculators)

	err := handler.Handle()

	assertError(t, err, boink)
	if logs.Len() > 0 {
		t.Logf("\n%s", logs.String())
	}
}

func TestCSVHandler_WriterError(t *testing.T) {
	rawInput := "2,-,2"
	var output = &WriterErr{err: boink}
	var logs bytes.Buffer
	handler := NewCSVHandler(strings.NewReader(rawInput), output, &logs, calculators)

	err := handler.Handle()

	assertError(t, err, boink)
	if logs.Len() > 0 {
		t.Logf("\n%s", logs.String())
	}

}
