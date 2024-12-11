package calculator

import (
	"context"
	"fmt"

	commandType "github.com/swintch/calc-apps/app/commands"
)

type Calculator interface{ Calculate(a, b int) int }

type Handler struct{ add, sub, mul, div Calculator }

func NewHandler(add, sub, mul, div Calculator) *Handler {
	return &Handler{
		add: add,
		sub: sub,
		mul: mul,
		div: div,
	}
}

func (this *Handler) Handle(ctx context.Context, commands ...any) {
	for _, command := range commands {
		switch commandType := command.(type) {
		case *commandType.Add:
			commandType.Result.C = this.add.Calculate(commandType.A, commandType.B)
		case *commandType.Subtraction:
			commandType.Result.C = this.sub.Calculate(commandType.A, commandType.B)
		case *commandType.Multiplication:
			commandType.Result.C = this.mul.Calculate(commandType.A, commandType.B)
		case *commandType.Division:
			commandType.Result.C = this.div.Calculate(commandType.A, commandType.B)
		default:
			panic(fmt.Sprintf("unsuported command: %T", commandType))
		}

	}

}
