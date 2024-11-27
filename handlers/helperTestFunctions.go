package handlers

import (
	"errors"
	"reflect"
	"testing"
)

var boink = errors.New("boink")

type WriterErr struct {
	err error
}

func (this *WriterErr) Write(p []byte) (n int, err error) {
	return 0, this.err
}

////////////////////////////

type ReaderErr struct {
	err error
}

func (this *ReaderErr) Read(p []byte) (n int, err error) { return 0, this.err }

//////////////////////

func assertEquals(t *testing.T, actual, expected any) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %s, expected = %s", actual, expected)
	}
}

func assertError(t *testing.T, error, expected error) {
	t.Helper()
	if !errors.Is(error, expected) {
		t.Errorf("Expected error %v, but got error %v", expected, error)
	}
}
