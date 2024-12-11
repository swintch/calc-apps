package inputs

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestHttpDivisionBind(t *testing.T) {
	gunit.Run(new(HttpDivisionBindFixture), t)
}

type HttpDivisionBindFixture struct {
	*gunit.Fixture
	division Division
}

func (this *HttpDivisionBindFixture) TestDivisionBothValuesGood() {
	request := setRequest("10", "2")
	err := this.division.Bind(request)
	this.So(err, should.BeNil)
	this.So(this.division.A, should.Equal, 10)
	this.So(this.division.B, should.Equal, 2)
}

func (this *HttpDivisionBindFixture) TestDivisionBadParamA() {
	request := setRequest("NaN", "2")
	err := this.division.Bind(request)
	this.So(err, should.NotBeNil)
	this.So(this.division, should.Equal, Division{})
}

func (this *HttpDivisionBindFixture) TestDivisionBadParamB() {
	request := setRequest("10", "NaN")
	err := this.division.Bind(request)
	this.So(err, should.NotBeNil)
	this.So(this.division, should.Equal, Division{})

}
