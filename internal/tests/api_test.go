package tests

import (
	"devices_crud/internal/devices"
	"devices_crud/internal/drivers/rest"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	rest.BuildRoutes(router, devices.NewDevicesDependencies(
		&devices.DeviceDependencies{UseMocks: true, Logger: log.New(os.Stdout, "TEST: ", log.Ltime)}))
	return router
}

func TestPing(t *testing.T) {
	router := setupRouter()
	expected := "{\"message\":\"Pong\"}"

	httpReq, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	responseData := w.Body.String()

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, expected, responseData)
}
