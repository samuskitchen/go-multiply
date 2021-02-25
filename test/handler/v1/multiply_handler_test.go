package v1

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/mock"
	servMock "go-multiply/domain/multiply/application/command/mocks"
	"go-multiply/domain/multiply/application/handler/v1"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewMultiplyHandler(t *testing.T) {
	tests := []struct {
		name string
		want *v1.MultiplyRouter
	}{
		{
			name: "Constructor Ok",
			want: v1.NewMultiplyHandler(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := v1.NewMultiplyHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMultiplyHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplyRouter_GetMultiplyNumbers(t *testing.T) {

	t.Run("Error Not Get Number1 Handler", func(tt *testing.T) {

		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/multiply/", nil)
		newRecorder := httptest.NewRecorder()
		mockService := &servMock.MultiplyCommand{}

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("number1", "1")

		queryParam := newRequest.URL.Query()
		queryParam.Add("number1", "")

		newRequest.URL.RawQuery = queryParam.Encode()
		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))

		testMultiplyHandler := &v1.MultiplyRouter{Serv: mockService}

		testMultiplyHandler.GetMultiplyNumbers(newRecorder, newRequest)
		mockService.AssertExpectations(tt)
	})

	t.Run("Error Not Get Number2 Handler", func(tt *testing.T) {

		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/multiply/", nil)
		newRecorder := httptest.NewRecorder()
		mockService := &servMock.MultiplyCommand{}

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("number1", "1")

		queryParam := newRequest.URL.Query()
		queryParam.Add("number1", "5")
		queryParam.Add("number2", "")

		newRequest.URL.RawQuery = queryParam.Encode()
		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))

		testMultiplyHandler := &v1.MultiplyRouter{Serv: mockService}

		testMultiplyHandler.GetMultiplyNumbers(newRecorder, newRequest)
		mockService.AssertExpectations(tt)
	})

	t.Run("Error Number1 Is No Numeric Handler", func(tt *testing.T) {

		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/multiply/", nil)
		newRecorder := httptest.NewRecorder()
		mockService := &servMock.MultiplyCommand{}

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("number1", "1")

		queryParam := newRequest.URL.Query()
		queryParam.Add("number1", "TEST")
		queryParam.Add("number2", "TEST")

		newRequest.URL.RawQuery = queryParam.Encode()
		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))

		testMultiplyHandler := &v1.MultiplyRouter{Serv: mockService}

		testMultiplyHandler.GetMultiplyNumbers(newRecorder, newRequest)
		mockService.AssertExpectations(tt)
	})

	t.Run("Error Number2 Is No Numeric Handler", func(tt *testing.T) {

		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/multiply/", nil)
		newRecorder := httptest.NewRecorder()
		mockService := &servMock.MultiplyCommand{}

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("number1", "1")

		queryParam := newRequest.URL.Query()
		queryParam.Add("number1", "10")
		queryParam.Add("number2", "TEST")

		newRequest.URL.RawQuery = queryParam.Encode()
		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))

		testMultiplyHandler := &v1.MultiplyRouter{Serv: mockService}

		testMultiplyHandler.GetMultiplyNumbers(newRecorder, newRequest)
		mockService.AssertExpectations(tt)
	})

	t.Run("Get Multiply Numeric Handler", func(tt *testing.T) {

		newRequest := httptest.NewRequest(http.MethodGet, "/api/v1/multiply/", nil)
		newRecorder := httptest.NewRecorder()
		mockService := &servMock.MultiplyCommand{}

		newRequestCtx := chi.NewRouteContext()
		newRequestCtx.URLParams.Add("number1", "1")

		queryParam := newRequest.URL.Query()
		queryParam.Add("number1", "10")
		queryParam.Add("number2", "2")

		newRequest.URL.RawQuery = queryParam.Encode()
		newRequest = newRequest.WithContext(context.WithValue(newRequest.Context(), chi.RouteCtxKey, newRequestCtx))
		mockService.On("MultiplyNumbers", mock.Anything, mock.Anything).Return(rand.Float64()).Once()

		testMultiplyHandler := &v1.MultiplyRouter{Serv: mockService}

		testMultiplyHandler.GetMultiplyNumbers(newRecorder, newRequest)
		mockService.AssertExpectations(tt)
	})
}
