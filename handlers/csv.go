package handlers

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

type CSVHandler struct {
	stdin  *csv.Reader
	stdout *csv.Writer
	stderr *log.Logger
}

func NewCSVHandler(stdin io.Reader, stdout io.Writer, stderr io.Writer) *CSVHandler {
	return &CSVHandler{
		stdin:  csv.NewReader(stdin),
		stdout: csv.NewWriter(stdout),
		stderr: log.New(stderr, "csv:", log.LstdFlags),
	}
}

func (this *CSVHandler) Handle() error {
	for {
		record, err := this.stdin.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			this.stderr.Println("Reader Error: ", err)
			return err
		}
		value1, err := strconv.Atoi(record[0])
		if err != nil {
			this.stderr.Printf("Invalid Operand [%v]. Error : %v", record[0], err)
			continue
		}
		calculator, err := getCalculator(record[1])
		if err != nil {
			this.stderr.Printf("Invalid Operator [%v]. Error : %v", record[1], err)
			continue
		}
		value2, err := strconv.Atoi(record[2])
		if err != nil {
			this.stderr.Printf("Invalid Operand [%v]. Error : %v", record[0], err)
			continue
		}
		result := calculator.Calculate(value1, value2)
		err = this.stdout.Write(append(record, strconv.Itoa(result)))

		if err != nil {
			this.stderr.Printf("Write Error: %v", err)
			return err
		}
	}
	this.stdout.Flush()
	return this.stdout.Error()
}
