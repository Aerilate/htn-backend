package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Aerilate/htn-backend/repository"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	server := NewServer(repository.NewRepo(nil))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
