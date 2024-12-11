package http

import (
	"net/http"

	"github.com/smarty/shuttle"
	shuttlelocal "github.com/swintch/calc-apps/externals/shuttle"
	"github.com/swintch/calc-apps/http/inputs"
)

func Router(calculator Handler) http.Handler {
	h := http.NewServeMux()
	processor := func() shuttlelocal.Processor { return NewProcessor(calculator) }
	h.Handle("/add", shuttlelocal.NewHandler(func() shuttlelocal.InputModel { return inputs.NewAddition() }, processor))
	h.Handle("/sub", shuttlelocal.NewHandler(func() shuttlelocal.InputModel { return inputs.NewSubtraction() }, processor))
	h.Handle("/mul", shuttlelocal.NewHandler(func() shuttlelocal.InputModel { return inputs.NewMultiplication() }, processor))
	h.Handle("/div", shuttlelocal.NewHandler(func() shuttlelocal.InputModel { return inputs.NewDivision() }, processor))
	return h
}

func SmartyShuttleRouter(calculator Handler) http.Handler {
	h := http.NewServeMux()
	processor := func() shuttle.Processor { return NewProcessor(calculator) }
	h.Handle("/add", shuttle.NewHandler(shuttle.Options.InputModel(func() shuttle.InputModel { return inputs.NewAddition() }), shuttle.Options.Processor(processor)))
	h.Handle("/sub", shuttle.NewHandler(shuttle.Options.InputModel(func() shuttle.InputModel { return inputs.NewSubtraction() }), shuttle.Options.Processor(processor)))
	h.Handle("/mul", shuttle.NewHandler(shuttle.Options.InputModel(func() shuttle.InputModel { return inputs.NewMultiplication() }), shuttle.Options.Processor(processor)))
	h.Handle("/div", shuttle.NewHandler(shuttle.Options.InputModel(func() shuttle.InputModel { return inputs.NewDivision() }), shuttle.Options.Processor(processor)))
	return h
}
