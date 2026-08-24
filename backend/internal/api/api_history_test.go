package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHistoryByMACRejectsInvalidLimit(t *testing.T) {
	router := setupTestRouter(t)

	tests := []string{
		"/api/history/AA:BB:CC:DD:EE:FF",
		"/api/history/AA:BB:CC:DD:EE:FF?num=abc",
		"/api/history/AA:BB:CC:DD:EE:FF?num=0",
		"/api/history/AA:BB:CC:DD:EE:FF?num=-1",
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
