package handlers

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/swintch/mdw-smarty-calc-lib2/calc"
)

type CSVHandler struct {
	stdin       *csv.Reader
	stdout      *csv.Writer
	stderr      *log.Logger
	calculators map[string]calc.Calculator
}

func NewCSVHandler(stdin io.Reader, stdout, stderr io.Writer, calculators map[string]calc.Calculator) *CSVHandler {
	return &CSVHandler{
		stdin:       csv.NewReader(stdin),
		stdout:      csv.NewWriter(stdout),
		stderr:      log.New(stderr, "csv:", log.LstdFlags),
		calculators: calculators,
	}
}

func (this *CSVHandler) Handle() error {
	for {
		record, err := this.stdin.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			this.stderr.Println(err)
			continue
		}
		value1, err := strconv.Atoi(record[0])
		if err != nil {
			this.stderr.Println(err)
			continue
		}
		calculator, ok := calculators[record[1]]
		if !ok {
			this.stderr.Println(err)
			continue
		}
		value2, err := strconv.Atoi(record[2])
		if err != nil {
			this.stderr.Println(err)
			continue
		}
		calcResult := calculator.Calculate(value1, value2)
		err = this.stdout.Write(append(record, strconv.Itoa(calcResult)))
		if err != nil {
			this.stderr.Println(err)
		}

	}
	this.stdout.Flush()
	return nil
}

var calculators = map[string]calc.Calculator{
	"+": &calc.Addition{},
	"-": &calc.Subtraction{},
	"*": &calc.Multiplication{},
	"/": &calc.Division{},
}
