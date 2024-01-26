package tests

import (
	"bytes"
	"devices_crud/internal/devices"
	"devices_crud/internal/devices/model"
	"devices_crud/internal/drivers/rest"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	rest.BuildRoutes(router, devices.NewDevicesDependencies(
		&devices.DeviceDependencies{UseMocks: true, Logger: log.New(os.Stdout, "TEST: ", log.Ltime)}))
	return router
}

func TestShouldReturnEmptyListWhenThereIsNoDevices(t *testing.T) {
	router := setupRouter()
	expected := "[]"

	httpReq, _ := http.NewRequest("GET", "/v1/devices", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	responseData := w.Body.String()

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, expected, responseData)
}

func TestShouldCreateDevice(t *testing.T) {
	router := setupRouter()
	body := "{\"name\":\"test\",\"deviceBrand\":\"test\"}"
	bodyReader := bytes.NewReader([]byte(body))

	httpReq, _ := http.NewRequest("POST", "/v1/devices", bodyReader)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	responseData := w.Body.String()
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.NotEqual(t, "", responseData)
}

func TestShouldReturnDevice(t *testing.T) {
	router := setupRouter()
	body := "{\"name\":\"test\",\"deviceBrand\":\"brand\"}"
	bodyReader := bytes.NewReader([]byte(body))

	httpReq, _ := http.NewRequest("POST", "/v1/devices", bodyReader)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	responseData := w.Body.String()
	parsedResponse := model.NewDeviceResponse{}
	json.Unmarshal([]byte(responseData), &parsedResponse)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.NotEqual(t, "", responseData)

	httpReq, _ = http.NewRequest("GET", fmt.Sprintf("/v1/devices/%s", parsedResponse.UUID), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	responseData = w.Body.String()
	getDeviceResponse := model.Device{}
	json.Unmarshal([]byte(responseData), &getDeviceResponse)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.NotEqual(t, "", responseData)
	assert.Equal(t, parsedResponse.UUID, getDeviceResponse.ID)
	assert.Equal(t, "test", getDeviceResponse.Name)
	assert.Equal(t, "brand", getDeviceResponse.DeviceBrand)
}

func TestShouldReturnDeviceList(t *testing.T) {
	router := setupRouter()
	parsedResponse, parsedResponse2 := addTwoDevices(router)

	httpReq, _ := http.NewRequest("GET", "/v1/devices", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	responseData := w.Body.String()
	getDevicesResponse := []model.Device{}
	json.Unmarshal([]byte(responseData), &getDevicesResponse)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.NotEqual(t, "", responseData)
	assert.Equal(t, 2, len(getDevicesResponse))

	foundDevicesNames := make([]string, 0)
	foundDevicesBrands := make([]string, 0)
	fundDevicesIds := make([]string, 0)
	for _, device := range getDevicesResponse {
		foundDevicesNames = append(foundDevicesNames, device.Name)
		foundDevicesBrands = append(foundDevicesBrands, device.DeviceBrand)
		fundDevicesIds = append(fundDevicesIds, device.ID)
	}

	assert.Contains(t, foundDevicesNames, "test_1")
	assert.Contains(t, foundDevicesNames, "test_2")
	assert.Contains(t, foundDevicesBrands, "brand_1")
	assert.Contains(t, foundDevicesBrands, "brand_2")
	assert.Contains(t, fundDevicesIds, parsedResponse.UUID)
	assert.Contains(t, fundDevicesIds, parsedResponse2.UUID)
}

func TestShouldDeleteDevice(t *testing.T) {
	router := setupRouter()
	parsedResponse, _ := addTwoDevices(router)

	httpReq, _ := http.NewRequest("DELETE", fmt.Sprintf("/v1/devices/%s", parsedResponse.UUID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	assert.Equal(t, 204, w.Code)

	httpReq, _ = http.NewRequest("GET", fmt.Sprintf("/v1/devices/%s", parsedResponse.UUID), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	assert.Equal(t, 404, w.Code)
}

func TestShouldReplaceDevice(t *testing.T) {
	router := setupRouter()
	parsedResponse, _ := addTwoDevices(router)

	body := "{\"name\":\"test_3\",\"deviceBrand\":\"brand_3\"}"
	bodyReader := bytes.NewReader([]byte(body))

	httpReq, _ := http.NewRequest("PUT", fmt.Sprintf("/v1/devices/%s", parsedResponse.UUID), bodyReader)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	responseData := w.Body.String()
	getDeviceResponse := model.Device{}
	json.Unmarshal([]byte(responseData), &getDeviceResponse)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.NotEqual(t, "", responseData)
	assert.Equal(t, parsedResponse.UUID, getDeviceResponse.ID)
	assert.Equal(t, "test_3", getDeviceResponse.Name)
	assert.Equal(t, "brand_3", getDeviceResponse.DeviceBrand)
}

func TestShouldPatchNameOfDevice(t *testing.T) {
	router := setupRouter()
	parsedResponse, _ := addTwoDevices(router)

	body := "{\"name\":\"test_3\"}"
	bodyReader := bytes.NewReader([]byte(body))

	httpReq, _ := http.NewRequest("PATCH", fmt.Sprintf("/v1/devices/%s", parsedResponse.UUID), bodyReader)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	httpReq, _ = http.NewRequest("GET", fmt.Sprintf("/v1/devices/%s", parsedResponse.UUID), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)
	responseData := w.Body.String()
	getDeviceResponse := model.Device{}
	json.Unmarshal([]byte(responseData), &getDeviceResponse)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.NotEqual(t, "", responseData)
	assert.Equal(t, parsedResponse.UUID, getDeviceResponse.ID)
	assert.Equal(t, "test_3", getDeviceResponse.Name)
	assert.Equal(t, "brand_1", getDeviceResponse.DeviceBrand)
}

func TestShouldSearchDevices(t *testing.T) {
	router := setupRouter()
	dev1, _ := addTwoDevices(router)
	httpReq, _ := http.NewRequest("GET", "/v1/devices/search?q=brand_1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)
	responseData := w.Body.String()
	getDevicesResponse := []model.Device{}
	json.Unmarshal([]byte(responseData), &getDevicesResponse)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.NotEqual(t, "", responseData)
	assert.Equal(t, 1, len(getDevicesResponse))
	assert.Equal(t, dev1.UUID, getDevicesResponse[0].ID)
	assert.Equal(t, "test_1", getDevicesResponse[0].Name)
	assert.Equal(t, "brand_1", getDevicesResponse[0].DeviceBrand)
}

func TestShouldNotFoundAnyDevices(t *testing.T) {
	router := setupRouter()
	addTwoDevices(router)
	httpReq, _ := http.NewRequest("GET", "/v1/devices/search?q=brand_3", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)
	responseData := w.Body.String()
	getDevicesResponse := []model.Device{}
	json.Unmarshal([]byte(responseData), &getDevicesResponse)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.NotEqual(t, "", responseData)
	assert.Equal(t, 0, len(getDevicesResponse))
}

func TestShoudlReturn404WhenDeviceNotFound(t *testing.T) {
	router := setupRouter()
	httpReq, _ := http.NewRequest("GET", "/v1/devices/123", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	assert.Equal(t, 404, w.Code)
}

func addTwoDevices(router *gin.Engine) (model.NewDeviceResponse, model.NewDeviceResponse) {
	body := "{\"name\":\"test_1\",\"deviceBrand\":\"brand_1\"}"
	bodyReader := bytes.NewReader([]byte(body))

	httpReq, _ := http.NewRequest("POST", "/v1/devices", bodyReader)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	responseData := w.Body.String()
	parsedResponse := model.NewDeviceResponse{}
	json.Unmarshal([]byte(responseData), &parsedResponse)

	body = "{\"name\":\"test_2\",\"deviceBrand\":\"brand_2\"}"
	bodyReader = bytes.NewReader([]byte(body))

	httpReq, _ = http.NewRequest("POST", "/v1/devices", bodyReader)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	responseData = w.Body.String()
	parsedResponse2 := model.NewDeviceResponse{}
	json.Unmarshal([]byte(responseData), &parsedResponse2)

	return parsedResponse, parsedResponse2
}
