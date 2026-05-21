package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid request",
			method:         http.MethodPost,
			body:           `{"name":"Somchai"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Hello Somchai"}`,
		},
		{
			name:           "Invalid method (GET)",
			method:         http.MethodGet,
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Invalid method (PUT)",
			method:         http.MethodPut,
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Invalid method (DELETE)",
			method:         http.MethodDelete,
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Invalid JSON",
			method:         http.MethodPost,
			body:           `{"name":"Somchai"`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Empty name",
			method:         http.MethodPost,
			body:           `{}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Empty body",
			method:         http.MethodPost,
			body:           ``,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Extra fields in JSON",
			method:         http.MethodPost,
			body:           `{"name":"Somchai","age":30}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(
				tt.method,
				"/hello",
				strings.NewReader(tt.body),
			)

			rec := httptest.NewRecorder()

			helloHandler(rec, req)

			res := rec.Result()

			if res.StatusCode != tt.expectedStatus {
				t.Errorf(
					"Expected status %d, got %d",
					tt.expectedStatus,
					res.StatusCode,
				)
			}

			body := strings.TrimSpace(
				rec.Body.String(),
			)

			if tt.expectedBody != "" &&
				body != tt.expectedBody {

				t.Errorf(
					"Expected body %s, got %s",
					tt.expectedBody,
					body,
				)
			}
		})
	}
}
