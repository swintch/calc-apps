package should

import (
	"testing"
)

type testStruct struct {
	a int
	b int
}

func pass(t *testing.T, actual any, assertion assertion, expected ...any) {
	t.Helper()
	f := NewFakeT()
	So(f, actual, assertion, expected...)
	if f.failure != nil {
		t.Error("Expected a passing assertion, but got failure: ", f.failure)
	}
}

func fail(t *testing.T, actual any, assertion assertion, expected ...any) {
	t.Helper()
	f := NewFakeT()
	So(f, actual, assertion, expected...)
	if f.failure == nil {
		t.Error("Expected a failing assertion, but the test passed!")
	}
}

func TestEqual(t *testing.T) {
	So(t, 1, Equal, 1)
	So(t, 1.23, Equal, 1.23)
	So(t, -1, Equal, -1)
	So(t, "hello", Equal, "hello")
	So(t, []int{1, 2}, Equal, []int{1, 2})
	So(t, testStruct{1, 2}, Equal, testStruct{1, 2})
}

func TestBeTrue(t *testing.T) {
	So(t, 1 == 1, BeTrue)
}

func TestBeFalse(t *testing.T) {
	So(t, 1 == 2, BeFalse)
}

func TestBeNil(t *testing.T) {
	So(t, nil, BeNil)
}

func TestNegated_BeNil(t *testing.T) {
	So(t, 1, NOT.BeNil)
}

func TestNegated_Equal(t *testing.T) {
	So(t, 1.23, NOT.Equal, 1.23)
}

/////////////////////////////////////////////////////////////

type FakeT struct{ failure error }

func NewFakeT() *FakeT { return &FakeT{} }

func (this *FakeT) Helper() {}
func (this *FakeT) Error(a ...any) {
	this.failure = a[0].(error)
}
