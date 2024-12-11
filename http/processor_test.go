package http

import (
	"context"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
	"github.com/swintch/calc"
	"github.com/swintch/calc-apps/app/calculator"
	"github.com/swintch/calc-apps/http/inputs"
	"github.com/swintch/calc-apps/http/views"
)

type FakeInput struct {
}

func TestProcessorHandler(t *testing.T) {
	gunit.Run(new(ProcessorFixture), t)
}

type ProcessorFixture struct {
	*gunit.Fixture
	processor *Processor
	handler   *calculator.Handler
}

func (this *ProcessorFixture) Setup() {
	this.handler = calculator.NewHandler(calc.Addition{}, calc.Subtraction{}, calc.Multiplication{}, calc.Division{})
	this.processor = NewProcessor(this.handler)
}

func (this *ProcessorFixture) TestAdditionProcessor() {
	additionInput := inputs.Addition{A: 1, B: 2}
	result := this.processor.Process(context.Background(), &additionInput)
	this.So(result, should.Equal, views.Addition{A: 1, B: 2, C: 3})
}

func (this *ProcessorFixture) TestMultiplicationProcessor() {
	multiplicationInput := inputs.Multiplication{A: 1, B: 2}
	result := this.processor.Process(context.Background(), &multiplicationInput)
	this.So(result, should.Equal, views.Multiplication{A: 1, B: 2, C: 2})
}

func (this *ProcessorFixture) TestSubtractionProcessor() {
	subtractionInput := inputs.Subtraction{A: 10, B: 2}
	result := this.processor.Process(context.Background(), &subtractionInput)
	this.So(result, should.Equal, views.Subtraction{A: 10, B: 2, C: 8})
}

func (this *ProcessorFixture) TestDivisionProcessor() {
	divisionInput := inputs.Division{A: 6, B: 2}
	result := this.processor.Process(context.Background(), &divisionInput)
	this.So(result, should.Equal, views.Division{A: 6, B: 2, C: 3})
}

func (this *ProcessorFixture) TestBadInputModelProcessor() {
	fakeInput := FakeInput{}
	result := this.processor.Process(context.Background(), &fakeInput)
	this.So(result, should.NotBeNil)
	this.So(result, should.Equal, internalServerError)
}
