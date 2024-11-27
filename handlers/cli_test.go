package handlers

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/swintch/mdw-smarty-calc-lib2/calc"
)

func TestHandler_NotEnoughArgs(t1 *testing.T) {
	handlerObj := NewCLIHandler(nil, nil)
	err := handlerObj.Handle(nil)
	assertError(t1, err, requiresTwoArgs)
}

func TestHandler_InvalidFirstArgument(t1 *testing.T) {
	testValues := []string{"NaN", "3"}
	handlerObj := NewCLIHandler(nil, nil)
	err := handlerObj.Handle(testValues)
	assertError(t1, err, InvalidArg)
	assertError(t1, err, strconv.ErrSyntax)

}

func TestHandler_InvalidSecondArgument(t1 *testing.T) {
	testValues := []string{"1", "NaN"}
	handlerObj := NewCLIHandler(nil, nil)
	err := handlerObj.Handle(testValues)
	assertError(t1, err, InvalidArg)
	assertError(t1, err, strconv.ErrSyntax)
}

func TestHandler_ValidArguments(t1 *testing.T) {
	var output bytes.Buffer
	testValues := []string{"1", "3"}
	calculator := &calc.Addition{}
	handlerObj := NewCLIHandler(calculator, &output)
	err := handlerObj.Handle(testValues)
	assertError(t1, err, nil)
	assertEquals(t1, output.String(), "4")
}

func TestHandler_WriterError(t1 *testing.T) {
	output := WriterErr{err: boink}
	testValues := []string{"1", "3"}
	calculator := &calc.Addition{}
	handlerObj := NewCLIHandler(calculator, &output)
	err := handlerObj.Handle(testValues)
	assertError(t1, err, boink)
	assertError(t1, err, WriterError)
}
