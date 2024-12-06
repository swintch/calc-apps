package handlers

import (
	"bytes"
	"os"
	"strconv"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestCLIHandler(t *testing.T) {
	gunit.Run(new(CLIHandlerFixture), t)
}

type CLIHandlerFixture struct {
	*gunit.Fixture
}

func (this *CLIHandlerFixture) TestInvalidNumberOfArguments() {
	handle := NewCLIHandler("+", os.Stdout)
	err := handle.Handler([]string{"1"})
	this.So(err, should.Wrap, InvalidNumberOfArguments)

}

func (this *CLIHandlerFixture) InvalidFirstArgument() {
	handle := NewCLIHandler("+", os.Stdout)
	err := handle.Handler([]string{"a", "2"})
	this.So(err, should.Wrap, InvalidArgumentFormat)
	this.So(err, should.Wrap, strconv.ErrSyntax)
}

func (this *CLIHandlerFixture) InvalidSecondArgument() {
	handle := NewCLIHandler("+", os.Stdout)
	err := handle.Handler([]string{"1", "b"})
	this.So(err, should.Wrap, InvalidArgumentFormat)
	this.So(err, should.Wrap, strconv.ErrSyntax)
}

func (this *CLIHandlerFixture) TestOutputToConsole() {
	output := bytes.Buffer{}
	handle := NewCLIHandler("+", &output)
	err := handle.Handler([]string{"1", "2"})
	this.So(err, should.Equal, nil)
	this.So(output.String(), should.Equal, "3")
}

func (this *CLIHandlerFixture) TestOutputToConsoleError() {
	output := &WriterError{err: boink}
	handle := NewCLIHandler("+", output)
	err := handle.Handler([]string{"1", "2"})
	this.So(err, should.Wrap, boink)
	this.So(err, should.Wrap, ConsoleWriteError)
}

func (this *CLIHandlerFixture) TestBadOperator() {
	output := bytes.Buffer{}
	handle := NewCLIHandler("adff", &output)
	err := handle.Handler([]string{"1", "2"})
	this.So(err, should.Wrap, InvalidOperator)
}

func (this *CLIHandlerFixture) TestAddition() {
	output := bytes.Buffer{}
	handle := NewCLIHandler("+", &output)
	err := handle.Handler([]string{"1", "2"})
	this.So(err, should.Equal, nil)
}

func (this *CLIHandlerFixture) TestSubtraction() {
	output := bytes.Buffer{}
	handle := NewCLIHandler("-", &output)
	err := handle.Handler([]string{"1", "2"})
	this.So(err, should.Equal, nil)
}

func (this *CLIHandlerFixture) TestMultiplication() {
	output := bytes.Buffer{}
	handle := NewCLIHandler("*", &output)
	err := handle.Handler([]string{"1", "2"})
	this.So(err, should.Equal, nil)
}

func (this *CLIHandlerFixture) TestDivision() {
	output := bytes.Buffer{}
	handle := NewCLIHandler("/", &output)
	err := handle.Handler([]string{"6", "2"})
	this.So(err, should.Equal, nil)
}
