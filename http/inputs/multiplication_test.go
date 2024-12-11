package inputs

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestHttpMultiplicationBind(t *testing.T) {
	gunit.Run(new(HttpMultiplicationBindFixture), t)
}

type HttpMultiplicationBindFixture struct {
	*gunit.Fixture
	multiplication Multiplication
}

func (this *HttpMultiplicationBindFixture) TestMultiplicationBothValuesGood() {
	request := setRequest("1", "2")
	err := this.multiplication.Bind(request)
	this.So(err, should.BeNil)
	this.So(this.multiplication.A, should.Equal, 1)
	this.So(this.multiplication.B, should.Equal, 2)
}

func (this *HttpMultiplicationBindFixture) TestMultiplicationBadParamA() {
	request := setRequest("NaN", "2")
	err := this.multiplication.Bind(request)
	this.So(err, should.NotBeNil)
	this.So(this.multiplication, should.Equal, Multiplication{})
}

func (this *HttpMultiplicationBindFixture) TestMultiplicationBadParamB() {
	request := setRequest("10", "NaN")
	err := this.multiplication.Bind(request)
	this.So(err, should.NotBeNil)
	this.So(this.multiplication, should.Equal, Multiplication{})

}
