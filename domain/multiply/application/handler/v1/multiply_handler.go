package v1

import (
	"errors"
	"go-multiply/domain/multiply/application/command"
	"go-multiply/domain/multiply/infrastructure/service"
	"go-multiply/infrastructure/config/middleware"
	"net/http"
	"strconv"
)

type MultiplyRouter struct {
	Serv command.MultiplyCommand
}

func NewMultiplyHandler() *MultiplyRouter {
	return &MultiplyRouter{
		Serv: service.NewService(),
	}
}

func (mr *MultiplyRouter) GetMultiplyNumbers(w http.ResponseWriter, r *http.Request) {

	number1Str := r.URL.Query().Get("number1")
	if number1Str == "" {
		_ = middleware.HTTPError(w, r, http.StatusBadRequest, errors.New("cannot get number1").Error())
		return
	}

	number2Str := r.URL.Query().Get("number2")
	if number2Str == "" {
		_ = middleware.HTTPError(w, r, http.StatusBadRequest, errors.New("cannot get number2").Error())
		return
	}

	number1, err := strconv.ParseFloat(number1Str, 64)
	if err != nil {
		_ = middleware.HTTPError(w, r, http.StatusBadRequest, errors.New("the number1 is not a numeric").Error())
		return
	}

	number2, err := strconv.ParseFloat(number2Str, 64)
	if err != nil {
		_ = middleware.HTTPError(w, r, http.StatusBadRequest, errors.New("the number2 is not a numeric").Error())
		return
	}

	result := mr.Serv.MultiplyNumbers(number1, number2)

	_ = middleware.JSON(w, r, http.StatusOK, result)
}
