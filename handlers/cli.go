package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/calc"
)

type CLIHandler struct {
	calculator calc.Calculator
	stdout     io.Writer
}

func NewCLIHandler(calculator calc.Calculator, stdout io.Writer) *CLIHandler {
	return &CLIHandler{
		calculator: calculator,
		stdout:     stdout,
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
	result := this.calculator.Calculate(value1, value2)
	_, err = fmt.Fprint(this.stdout, result)
	if err != nil {
		return fmt.Errorf("%w: %w", ConsoleWriteError, err)
	}
	return nil
}

var (
	InvalidNumberOfArguments = errors.New("Two arguments must be provided")
	InvalidArgumentFormat    = errors.New("Invalid argument format. Must be in integer format")
	ConsoleWriteError        = errors.New("Write to console failed")
)
