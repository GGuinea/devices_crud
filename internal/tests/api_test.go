package tests

import (
	"devices_crud/internal/drivers/rest"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	rest.BuildRoutes(router)
	return router
}

func TestPing(t *testing.T) {
	router := setupRouter()
	expected := "{\"message\":\"Pong\"}"

	httpReq, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	responseData := w.Body.String()

	if w.Code != 200 {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	if responseData != expected {
		t.Errorf("Expected {\"message\":\"Pong\"}, got %s", responseData)
	}
}
