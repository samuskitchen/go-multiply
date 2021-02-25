package integration

import (
	"encoding/json"
	"fmt"
	"go-multiply/infrastructure/config"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestIntegration_GetMultiplyNumbers(t *testing.T) {

	tests := []struct {
		Name          string
		Number1       string
		Number2       string
		ExpectedValue float64
		ExpectedCode  int
	}{
		{
			Name:          "Get Multiply Successful",
			Number1:       "5",
			Number2:       "2",
			ExpectedValue: 10,
			ExpectedCode:  200,
		},
		{
			Name:          "Error Get Number1",
			Number1:       "",
			Number2:       "",
			ExpectedValue: 0,
			ExpectedCode:  400,
		},
		{
			Name:          "Error Get Number2",
			Number1:       "1",
			Number2:       "",
			ExpectedValue: 0,
			ExpectedCode:  400,
		},
		{
			Name:          "Error number1 is not a numeric",
			Number1:       "TEST",
			Number2:       "TEST",
			ExpectedValue: 0,
			ExpectedCode:  400,
		},
		{
			Name:          "Error number2 is not a numeric",
			Number1:       "1",
			Number2:       "TEST",
			ExpectedValue: 0,
			ExpectedCode:  400,
		},
	}

	for _, test := range tests {
		fn := func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/multiply?number1=%s&number2=%s", test.Number1, test.Number2), nil)
			if err != nil {
				t.Errorf("error multiply request: %v", err)
			}

			w := httptest.NewRecorder()
			config.NewServerTest(os.Getenv("API_PORT")).ServeHTTP(w, req)

			if e, a := test.ExpectedCode, w.Code; e != a {
				t.Errorf("expected status code: %v, got status code: %v", e, a)
			}

			if test.ExpectedCode != http.StatusBadRequest {
				var valueResponse float64

				if err := json.NewDecoder(w.Body).Decode(&valueResponse); err != nil {
					t.Errorf("error decoding valueResponse body: %v", err)
				}

				if e, a := test.ExpectedValue, valueResponse; e != a {
					t.Errorf("expected total value: %v, got total value: %v", e, a)
				}
			}
		}

		t.Run(test.Name, fn)
	}

}
