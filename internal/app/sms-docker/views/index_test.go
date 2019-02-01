package views

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/bakhtik/sms-docker/internal/pkg/testing"
	"github.com/go-redis/redis"
)

func TestIndexHandler(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{})
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	h := IndexHandler(redisClient)
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
