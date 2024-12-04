package handlers

import (
	"bytes"
	"strings"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestCSVHandler(t *testing.T) {
	gunit.Run(new(CSVHandlerFixture), t)
}

type CSVHandlerFixture struct {
	*gunit.Fixture
}

func (this *CSVHandlerFixture) TestAddition() {
	output := bytes.Buffer{}
	input := "1,+,2"
	handle := NewCSVHandler(strings.NewReader(input), &output, this)
	err := handle.Handle()
	this.So(err, should.Equal, nil)
	this.So(output.String(), should.Equal, "1,+,2,3\n")
}

func (this *CSVHandlerFixture) TestMultipleValues() {
	output := bytes.Buffer{}
	input := "1,+,2\n2,-,2\n2,*,2\n6,/,3"
	handle := NewCSVHandler(strings.NewReader(input), &output, this)
	err := handle.Handle()
	this.So(err, should.Equal, nil)
	this.So(output.String(), should.Equal, "1,+,2,3\n2,-,2,0\n2,*,2,4\n6,/,3,2\n")
}

func (this *CSVHandlerFixture) TestMultipleBadValues() {
	output := bytes.Buffer{}
	input := "1,+,2\nNaN,-,2\n2,NaN,2\n6,/,NaN"
	handle := NewCSVHandler(strings.NewReader(input), &output, this)
	err := handle.Handle()
	this.So(err, should.Equal, nil)
	this.So(output.String(), should.Equal, "1,+,2,3\n")
}

func (this *CSVHandlerFixture) TestReadError() {
	output := bytes.Buffer{}
	input := ReaderError{err: boink}
	handle := NewCSVHandler(&input, &output, this)
	err := handle.Handle()
	this.So(err, should.Equal, boink)
}

func (this *CSVHandlerFixture) TestWriteError() {
	output := WriterError{err: boink}
	input := "1,+,2"
	handle := NewCSVHandler(strings.NewReader(input), &output, this)
	err := handle.Handle()
	this.So(err, should.Equal, boink)
}
