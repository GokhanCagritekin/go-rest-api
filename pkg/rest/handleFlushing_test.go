package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GokhanCagritekin/go-rest-api/pkg/flushing"
)

var fs = flushing.NewService(r)

func TestHandleFlushing(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:8080/flush", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleFlushing(fs))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
