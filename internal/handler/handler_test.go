package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	reponseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(HomeHandler)
	handle.ServerHTTP(responseRecorder, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", status)
	}

	expected := "Hello worldies"
	if responseRecorder.Body.String() != expected {
		t.Errorf("Expected %s, got %s", expected, responseRecorder.Body.String())
	}
}
