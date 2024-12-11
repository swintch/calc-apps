package inputs

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestHttpSubtractionBind(t *testing.T) {
	gunit.Run(new(HttpSubtractionBindFixture), t)
}

type HttpSubtractionBindFixture struct {
	*gunit.Fixture
	subtraction Subtraction
}

func (this *HttpSubtractionBindFixture) TestSubtractionBothValuesGood() {
	request := setRequest("1", "2")
	err := this.subtraction.Bind(request)
	this.So(err, should.BeNil)
	this.So(this.subtraction.A, should.Equal, 1)
	this.So(this.subtraction.B, should.Equal, 2)
}

func (this *HttpSubtractionBindFixture) TestSubtractionBadParamA() {
	request := setRequest("NaN", "2")
	err := this.subtraction.Bind(request)
	this.So(err, should.NotBeNil)
	this.So(this.subtraction, should.Equal, Subtraction{})
}

func (this *HttpSubtractionBindFixture) TestSubtractionBadParamB() {
	request := setRequest("10", "NaN")
	err := this.subtraction.Bind(request)
	this.So(err, should.NotBeNil)
	this.So(this.subtraction, should.Equal, Subtraction{})

}
