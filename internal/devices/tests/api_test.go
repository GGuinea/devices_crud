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
)

func setupRouter() *gin.Engine {
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

	if w.Code != 200 {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	if responseData != expected {
		t.Errorf("Expected [], got %s", responseData)
	}
}

func TestShouldCreateDevice(t *testing.T) {
	router := setupRouter()
	body := "{\"name\":\"test\",\"deviceBrand\":\"test\"}"
	bodyReader := bytes.NewReader([]byte(body))

	httpReq, _ := http.NewRequest("POST", "/v1/devices", bodyReader)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	responseData := w.Body.String()

	if w.Code != 201 {
		t.Errorf("Expected 201, got %d", w.Code)
	}
	if responseData == "" {
		t.Errorf("Expected body, got %s", responseData)
	}
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

	if w.Code != 201 {
		t.Errorf("Expected 201, got %d", w.Code)
	}
	if responseData == "" {
		t.Errorf("Expected body, got %s", responseData)
	}

	httpReq, _ = http.NewRequest("GET", fmt.Sprintf("/v1/devices/%s", parsedResponse.UUID), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	responseData = w.Body.String()
	getDeviceResponse := model.Device{}
	json.Unmarshal([]byte(responseData), &getDeviceResponse)
	fmt.Println(getDeviceResponse)

	if w.Code != 200 {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	if getDeviceResponse.ID != parsedResponse.UUID {
		t.Errorf("Expected %s, got %s", parsedResponse.UUID, getDeviceResponse.ID)
	}

	if getDeviceResponse.Name != "test" {
		t.Errorf("Expected test, got %s", getDeviceResponse.Name)
	}

	if getDeviceResponse.DeviceBrand != "brand" {
		t.Errorf("Expected test, got %s", getDeviceResponse.DeviceBrand)
	}

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

	if w.Code != 200 {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	if len(getDevicesResponse) != 2 {
		t.Errorf("Expected 2 devices, got %d", len(getDevicesResponse))
	}

	if getDevicesResponse[0].ID != parsedResponse.UUID {
		t.Errorf("Expected %s, got %s", parsedResponse.UUID, getDevicesResponse[0].ID)
	}

	if getDevicesResponse[0].Name != "test_1" {
		t.Errorf("Expected test_1, got %s", getDevicesResponse[0].Name)
	}

	if getDevicesResponse[0].DeviceBrand != "brand_1" {
		t.Errorf("Expected brand_1, got %s", getDevicesResponse[0].DeviceBrand)
	}

	if getDevicesResponse[1].ID != parsedResponse2.UUID {
		t.Errorf("Expected %s, got %s", parsedResponse2.UUID, getDevicesResponse[1].ID)
	}

	if getDevicesResponse[1].Name != "test_2" {
		t.Errorf("Expected test_2, got %s", getDevicesResponse[1].Name)
	}

	if getDevicesResponse[1].DeviceBrand != "brand_2" {
		t.Errorf("Expected brand_2, got %s", getDevicesResponse[1].DeviceBrand)
	}
}

func TestShouldDeleteDevice(t *testing.T) {
	router := setupRouter()
	parsedResponse, _ := addTwoDevices(router)

	httpReq, _ := http.NewRequest("DELETE", fmt.Sprintf("/v1/devices/%s", parsedResponse.UUID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	if w.Code != 200 {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	httpReq, _ = http.NewRequest("GET", fmt.Sprintf("/v1/devices/%s", parsedResponse.UUID), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, httpReq)

	if w.Code != 404 {
		t.Errorf("Expected 404, got %d", w.Code)
	}
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

	if w.Code != 200 {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	if getDeviceResponse.ID != parsedResponse.UUID {
		t.Errorf("Expected %s, got %s", parsedResponse.UUID, getDeviceResponse.ID)
	}

	if getDeviceResponse.Name != "test_3" {
		t.Errorf("Expected test_3, got %s", getDeviceResponse.Name)
	}

	if getDeviceResponse.DeviceBrand != "brand_3" {
		t.Errorf("Expected brand_3, got %s", getDeviceResponse.DeviceBrand)
	}
}

func TestShouldPatchDevice(t *testing.T) {
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

	if w.Code != 200 {
		t.Errorf("Expected 200, got %d", w.Code)
	}

	if getDeviceResponse.ID != parsedResponse.UUID {
		t.Errorf("Expected %s, got %s", parsedResponse.UUID, getDeviceResponse.ID)
	}

	if getDeviceResponse.Name != "test_3" {
		t.Errorf("Expected test_3, got %s", getDeviceResponse.Name)
	}

	if getDeviceResponse.DeviceBrand != "brand_1" {
		t.Errorf("Expected brand_1, got %s", getDeviceResponse.DeviceBrand)
	}
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
	if w.Code != 200 {
		t.Errorf("Expected 200, got %d", w.Code)
	}
	if len(getDevicesResponse) != 1 {
		t.Errorf("Expected 1 devices, got %d", len(getDevicesResponse))
	}

	if getDevicesResponse[0].ID != dev1.UUID {
		t.Errorf("Expected %s, got %s", dev1.UUID, getDevicesResponse[0].ID)
	}
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
	if w.Code != 200 {
		t.Errorf("Expected 200, got %d", w.Code)
	}
	if len(getDevicesResponse) != 0 {
		t.Errorf("Expected 0 devices, got %d", len(getDevicesResponse))
	}
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
