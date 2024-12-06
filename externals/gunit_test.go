package gunit

import (
	"fmt"
	"testing"

	should "github.com/swintch/calc-apps/Should"
)

func TestGunitTestHandler(t *testing.T) {
	Run(t, new(GunitHandlerFixture))
}

type GunitHandlerFixture struct {
	*Fixture
}

func (this *GunitHandlerFixture) Setup() {

	fmt.Println("Setup")
}

func (this *GunitHandlerFixture) Test1() {

	this.So(1, should.Equal, 1)
}

func (this *GunitHandlerFixture) Test2() {

	fmt.Println("Test2")
}
