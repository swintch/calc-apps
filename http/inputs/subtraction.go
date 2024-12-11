package inputs

import (
	"net/http"

	"github.com/smarty/shuttle"
)

type Subtraction struct {
	shuttle.BaseInputModel
	A int
	B int
}

func NewSubtraction() *Subtraction {
	return &Subtraction{}
}

func (this *Subtraction) Bind(request *http.Request) error {
	query := request.URL.Query()
	valueRequestA, err := parseInteger(query, "a")
	if err != nil {
		return err
	}
	valueRequestB, err := parseInteger(query, "b")
	if err != nil {
		return err
	}
	this.A = valueRequestA
	this.B = valueRequestB
	return nil
}
