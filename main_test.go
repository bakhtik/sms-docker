package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexHandler(t *testing.T) {

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	h := indexHandler()
	h.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response code is not 200: %v", resp.StatusCode)
	}
	if strings.Contains(string(body), "Welcome to Security Management System Docker edition!") == false {
		t.Errorf("Body does not contain 'Welcome to Security Management System Docker edition!'")
	}
}
