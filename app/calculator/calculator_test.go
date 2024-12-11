package calculator

import (
	"context"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
	"github.com/swintch/calc"
	"github.com/swintch/calc-apps/app/commands"
)

func TestCalculatorHandler(t *testing.T) {
	gunit.Run(new(CalculatorHandlerFixture), t)
}

type CalculatorHandlerFixture struct {
	*gunit.Fixture
	handler Handler
}

func (this *CalculatorHandlerFixture) Setup() {
	this.handler.add = calc.Addition{}
	this.handler.sub = calc.Subtraction{}
	this.handler.div = calc.Division{}
	this.handler.mul = calc.Multiplication{}
}

func (this *CalculatorHandlerFixture) TestAdditionHandler() {
	command := &commands.Add{A: 1, B: 2}
	this.handler.Handle(context.Background(), command)
	this.So(command.Result.C, should.Equal, 3)
	this.So(command.Result.Error, should.BeNil)

}

func (this *CalculatorHandlerFixture) TestSubtractionHandler() {
	command := &commands.Subtraction{A: 6, B: 2}
	this.handler.Handle(context.Background(), command)
	this.So(command.Result.C, should.Equal, 4)
	this.So(command.Result.Error, should.BeNil)

}

func (this *CalculatorHandlerFixture) TestMultiplicationHandler() {
	command := &commands.Multiplication{A: 10, B: 10}
	this.handler.Handle(context.Background(), command)
	this.So(command.Result.C, should.Equal, 100)
	this.So(command.Result.Error, should.BeNil)

}

func (this *CalculatorHandlerFixture) TestDivisionHandler() {
	command := &commands.Division{A: 10, B: 2}
	this.handler.Handle(context.Background(), command)
	this.So(command.Result.C, should.Equal, 5)
	this.So(command.Result.Error, should.BeNil)

}
