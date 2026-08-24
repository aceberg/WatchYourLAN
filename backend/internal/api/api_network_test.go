package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPortEndpointRejectsInvalidPort(t *testing.T) {
	router := setupTestRouter(t)

	tests := []string{
		"/api/port/127.0.0.1/not-a-port",
		"/api/port/127.0.0.1/0",
		"/api/port/127.0.0.1/-1",
		"/api/port/127.0.0.1/65536",
	}

	for _, path := range tests {
		t.Run(path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, path, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want %d; body: %s", rec.Code, http.StatusBadRequest, rec.Body.String())
			}
		})
	}
}
