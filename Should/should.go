package should

import (
	"fmt"
	"reflect"
)

type testingT interface {
	Helper()
	Error(...any)
}

type assertion func(actual any, expected ...any) error

func So(t testingT, actual any, assert assertion, expected ...any) bool {
	t.Helper()
	err := assert(actual, expected...)
	if err != nil {
		t.Error(err)
		return false
	}
	return true

}

func Equal(actual any, expected ...any) error {
	if reflect.DeepEqual(actual, expected[0]) {
		return nil
	}
	return fmt.Errorf("expected value = %v, actual value = %v", expected, actual)
}
func BeTrue(actual any, _ ...any) error {
	return Equal(actual, true)
}
func BeFalse(actual any, _ ...any) error {
	return Equal(actual, false)
}
func BeNil(actual any, _ ...any) error {
	return Equal(actual, nil)
}

type negated struct{}

var NOT negated

func (negated) Equal(actual any, expected ...any) error {
	if !reflect.DeepEqual(actual, expected[0]) {
		return nil
	}
	return fmt.Errorf("expected value = %v, actual value = %v", expected, actual)
}
func (negated) BeNil(actual any, _ ...any) error {
	return NOT.Equal(actual, nil)
}
