package inputs

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestHttpAdditionBind(t *testing.T) {
	gunit.Run(new(HttpAdditionBindFixture), t)
}

type HttpAdditionBindFixture struct {
	*gunit.Fixture
	addition Addition
}

func setRequest(queryStringValueA string, queryStringValueB string) *http.Request {
	request := httptest.NewRequest("GET", "/", nil)
	query := request.URL.Query()
	query.Set("a", queryStringValueA)
	query.Set("b", queryStringValueB)
	request.URL.RawQuery = query.Encode()
	return request
}

func (this *HttpAdditionBindFixture) TestAdditionBothValuesGood() {
	request := setRequest("1", "2")
	err := this.addition.Bind(request)
	this.So(err, should.BeNil)
	this.So(this.addition.A, should.Equal, 1)
	this.So(this.addition.B, should.Equal, 2)
}

func (this *HttpAdditionBindFixture) TestAdditionBadParamA() {
	request := setRequest("NaN", "2")
	err := this.addition.Bind(request)
	this.So(err, should.NotBeNil)
	this.So(this.addition, should.Equal, Addition{})
}

func (this *HttpAdditionBindFixture) TestAdditionBadParamB() {
	request := setRequest("1", "NaN")
	err := this.addition.Bind(request)
	this.So(err, should.NotBeNil)
	this.So(this.addition, should.Equal, Addition{})

}
