package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/swintch/mdw-smarty-calc-lib2/calc"
)

type CLIHandler struct {
	calculator calc.Calculator
	stdout     io.Writer
}

func (this *CLIHandler) Handle(args []string) error {
	if len(args) != 2 {
		return requiresTwoArgs
	}
	integerValue1, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: %w", InvalidArg, err)
	}
	integerValue2, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: %w", InvalidArg, err)
	}
	result := this.calculator.Calculate(integerValue1, integerValue2)
	_, err = fmt.Fprintf(this.stdout, "%d", result)
	if err != nil {
		return fmt.Errorf("%w: %w", WriterError, err)
	}
	return nil

}

func NewCLIHandler(calculator calc.Calculator, stdout io.Writer) *CLIHandler {
	return &CLIHandler{
		calculator: calculator,
		stdout:     stdout,
	}
}

var (
	requiresTwoArgs = errors.New("Requires two arguments!")
	InvalidArg      = errors.New("Invalid argument!")
	WriterError     = errors.New("Writer error!")
)
