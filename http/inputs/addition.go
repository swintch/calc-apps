package inputs

import (
	"net/http"

	"github.com/smarty/shuttle"
)

type Addition struct {
	shuttle.BaseInputModel
	A int
	B int
}

func NewAddition() *Addition {
	return &Addition{}
}

func (this *Addition) Bind(request *http.Request) error {
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
