package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	//structObject := testStruct{}
	//var testStruct1 *testStruct
	//instance := reflect.ValueOf(testStruct1)
	//value := instance.NumMethod()
	//fmt.Println(value)
	//testName := instance.Method(0)
	//fmt.Println(testName)
	//instance := reflect.New(fixtureType)
	//testMethod := instance.MethodByName("Function1")
	//testMethod.Call([]reflect.Value{})
	/*for i := 0; i < instance.NumMethod(); i++ {
		returnedMethod := instance.Method(i)
		returnedMethod.Call([]reflect.Value{})
	}*/
	Run(new(testStruct))

}

type testStruct struct{}

func (t *testStruct) Setup() {
	fmt.Println("setup")
}

func (t *testStruct) Test1() {
	fmt.Println("test1")
}

func (t *testStruct) Test2() {
	fmt.Println("test2")
}

func Run(fixture any) {
	fixtureTypeOf := reflect.TypeOf(fixture)
	for i := 0; i < fixtureTypeOf.NumMethod(); i++ {
		testMethodName := fixtureTypeOf.Method(i).Name
		instance := reflect.ValueOf(fixture)
		if strings.HasPrefix(testMethodName, "Setup") {
			setupMethod := instance.MethodByName("Setup")
			callableSetup := setupMethod.Interface().(func())
			callableSetup()
		}
		if strings.HasPrefix(testMethodName, "Test") {
			MethodToCall := instance.MethodByName(testMethodName)
			callableTest := MethodToCall.Interface().(func())
			callableTest()
		}
	}
}
