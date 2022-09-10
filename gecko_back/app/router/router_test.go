package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {

	e, err := Router()
	if err != nil {
		t.Fatal(err)
	}

	s := httptest.NewServer(e)

	r, err := http.Get(s.URL + "/api/v1/ping")
	if err != nil {
		t.Errorf("server initializaion error = %v", err)
		return
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		t.Errorf("server response error = %v", r.StatusCode)
		return
	}

}
