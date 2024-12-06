package calculator

import (
	"context"

	"github.com/swintch/calc"
	commands1 "github.com/swintch/calc-apps/app/commands"
)

type Calculator interface{ Calculate(a, b int) int }

type Handler struct{ add, sub, mul, div Calculator }

func (this *Handler) Handle(ctx context.Context, commands ...any) {

	switch typed := commands[0].(type) {
	case *commands1.Add:
		this.ProcessAddition(ctx, typed)
	case *commands1.Subtraction:

	case *commands1.Multiplication:

	case *commands1.Division:
	}
}

func (this *Handler) ProcessAddition(ctx context.Context, command *commands1.Add) {
	calculator := calc.Addition{}
	result := calculator.Calculate(command.A, command.B)
	command.Result.C = result
}
