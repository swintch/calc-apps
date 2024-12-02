package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/swintch/calc"
)

type CLIHandler struct {
	operator string
	stdout   io.Writer
}

func NewCLIHandler(operator string, stdout io.Writer) *CLIHandler {
	return &CLIHandler{
		operator: operator,
		stdout:   stdout,
	}
}

func (this CLIHandler) Handler(args []string) error {

	if len(args) != 2 {
		return fmt.Errorf("%w", InvalidNumberOfArguments)
	}
	value1, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: %w", InvalidArgumentFormat, err)
	}
	value2, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: %w", InvalidArgumentFormat, err)
	}
	calculator, err := getCalculator(this.operator)
	if err != nil {
		return fmt.Errorf("%w", InvalidOperator)
	}
	result := calculator.Calculate(value1, value2)
	_, err = fmt.Fprint(this.stdout, result)
	if err != nil {
		return fmt.Errorf("%w: %w", ConsoleWriteError, err)
	}
	return nil
}

func getCalculator(s string) (calc.Calculator, error) {
	calculator, ok := calculators[s]
	if !ok {
		return nil, InvalidOperator
	}
	return calculator, nil
}

var (
	InvalidNumberOfArguments = errors.New("Two arguments must be provided")
	InvalidArgumentFormat    = errors.New("Invalid argument format. Must be in integer format")
	ConsoleWriteError        = errors.New("Write to console failed")
	InvalidOperator          = errors.New("Invalid operator. Expect +, -, *, or /")
)

var calculators = map[string]calc.Calculator{
	"+": &calc.Addition{},
	"-": &calc.Subtraction{},
	"*": &calc.Multiplication{},
	"/": &calc.Division{},
}
