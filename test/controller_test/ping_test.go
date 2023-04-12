package controller_test

import (
	"github.com/stretchr/testify/assert"
	"go-mongosteen/internal/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandle(t *testing.T) {
	r := router.New()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello World", w.Body.String())
}
