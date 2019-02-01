package views

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/bakhtik/sms-docker/internal/pkg/testing"
)

type CacheMock struct{}

func (cm *CacheMock) Increment(key string) (result int64, err error) {
	return 42, nil
}

func TestIndexHandler(t *testing.T) {

	cache := &CacheMock{}
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	h := IndexHandler(cache)
	h.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response code is not 200: %v", resp.StatusCode)
	}
	if strings.Contains(string(body), "Welcome to Security Management System Docker edition!") == false {
		t.Errorf("Body does not contain 'Welcome to Security Management System Docker edition!'")
	}
	if strings.Contains(string(body), "Visits: 42") == false {
		t.Errorf("Body does not contain 'Visits: 42'")
	}
}
