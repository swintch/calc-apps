package gunit

import (
	"reflect"
	"strings"
	"testing"
)

func Run(t *testing.T, fixture any) {
	fixtureTypeOf := reflect.TypeOf(fixture)
	for i := 0; i < fixtureTypeOf.NumMethod(); i++ {
		testMethodName := fixtureTypeOf.Method(i).Name
		if strings.HasPrefix(testMethodName, "Test") {
			t.Run(testMethodName, func(t *testing.T) {
				instance := reflect.New(fixtureTypeOf.Elem())
				instance.Elem().FieldByName("Fixture").Set(reflect.ValueOf(&Fixture{T: t}))
				setupMethod := instance.MethodByName("Setup")
				if setupMethod.IsValid() {
					callableSetupMethod := setupMethod.Interface().(func())
					callableSetupMethod()
				}
				MethodToCall := instance.MethodByName(testMethodName)
				callableTestMethod := MethodToCall.Interface().(func())
				callableTestMethod()

			})
		}
	}
}

type Fixture struct{ *testing.T }

func (this *Fixture) So(actual any, assert assertion, expected ...any) {
	failure := assert(actual, expected...)
	if failure != nil {
		this.T.Error(failure)
	}

}

type assertion func(actual any, expected ...any) error
