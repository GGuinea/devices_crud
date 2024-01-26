package tests

import (
	"context"
	"devices_crud/internal/devices/app"
	"devices_crud/internal/devices/app/ports"
	"devices_crud/internal/devices/model"
	"log"
	"os"
	"testing"
)

func getDeviceService() *app.DeviceService {
	deviceRepositoryMock := ports.NewDevicesRepositoryMock()
	logger := log.New(os.Stdout, "TEST: ", log.Ltime)

	return app.NewDeviceService(deviceRepositoryMock, logger)
}

func TestShouldAddDevice(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	device := &model.NewDeviceRequest{
		Name:        "Test Device",
		DeviceBrand: "Test Brand",
	}

	id, err := deviceService.AddDevice(ctx, device)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if id == nil {
		t.Errorf("Expected id, got nil")
	}
}

func TestShouldGetDevice(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	device := &model.NewDeviceRequest{
		Name:        "Test Device",
		DeviceBrand: "Test Brand",
	}

	id, err := deviceService.AddDevice(ctx, device)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if id == nil {
		t.Errorf("Expected id, got nil")
	}

	deviceFound, err := deviceService.GetDevice(ctx, *id)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if deviceFound == nil {
		t.Errorf("Expected device, got nil")
	}
	if deviceFound.ID != *id {
		t.Errorf("Expected %s, got %s", *id, deviceFound.ID)
	}
}

func TestShouldGetAllDevices(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	device := &model.NewDeviceRequest{
		Name:        "Test Device",
		DeviceBrand: "Test Brand",
	}

	_, err := deviceService.AddDevice(ctx, device)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	device_2 := &model.NewDeviceRequest{
		Name:        "Test Device 2",
		DeviceBrand: "Test Brand 2",
	}

	_, err = deviceService.AddDevice(ctx, device_2)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	devices, err := deviceService.GetAllDevices(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if len(devices) != 2 {
		t.Errorf("Expected 2 devices, got %d", len(devices))
	}

	if devices[0].Name != "Test Device" {
		t.Errorf("Expected Test Device, got %s", devices[0].Name)
	}

	if devices[1].Name != "Test Device 2" {
		t.Errorf("Expected Test Device 2, got %s", devices[1].Name)
	}
}

func TestShouldReplaceDevice(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	device := &model.NewDeviceRequest{
		Name:        "Test Device",
		DeviceBrand: "Test Brand",
	}

	id, err := deviceService.AddDevice(ctx, device)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if id == nil {
		t.Errorf("Expected id, got nil")
	}

	deviceFound, err := deviceService.GetDevice(ctx, *id)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if deviceFound == nil {
		t.Errorf("Expected device, got nil")
	}
	if deviceFound.ID != *id {
		t.Errorf("Expected %s, got %s", *id, deviceFound.ID)
	}

	deviceFound.Name = "Test Device 2"
	deviceFound.DeviceBrand = "Test Brand 2"

	deviceReplaced, err := deviceService.ReplaceDevice(ctx, deviceFound)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if deviceReplaced == nil {
		t.Errorf("Expected device, got nil")
	}
	if deviceReplaced.ID != *id {
		t.Errorf("Expected %s, got %s", *id, deviceReplaced.ID)
	}
	if deviceReplaced.Name != "Test Device 2" {
		t.Errorf("Expected Test Device 2, got %s", deviceReplaced.Name)
	}
	if deviceReplaced.DeviceBrand != "Test Brand 2" {
		t.Errorf("Expected Test Brand 2, got %s", deviceReplaced.DeviceBrand)
	}
}

func TestShouldPatchDevice(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	device := &model.NewDeviceRequest{
		Name:        "Test Device",
		DeviceBrand: "Test Brand",
	}

	id, err := deviceService.AddDevice(ctx, device)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if id == nil {
		t.Errorf("Expected id, got nil")
	}

	deviceFound, err := deviceService.GetDevice(ctx, *id)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if deviceFound == nil {
		t.Errorf("Expected device, got nil")
	}
	if deviceFound.ID != *id {
		t.Errorf("Expected %s, got %s", *id, deviceFound.ID)
	}

	newName := "Test Device 2"
	newDeviceBrand := "Test Brand 2"

	patchDevice := &model.PatchDeviceRequest{
		ID:          *id,
		Name:        &newName,
		DeviceBrand: &newDeviceBrand,
	}

	idPatched, err := deviceService.PatchDevice(ctx, patchDevice)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if idPatched == nil {
		t.Errorf("Expected id, got nil")
	}

	devicePatched, err := deviceService.GetDevice(ctx, *idPatched)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if devicePatched == nil {
		t.Errorf("Expected device, got nil")
	}
	if devicePatched.ID != *id {
		t.Errorf("Expected %s, got %s", *id, devicePatched.ID)
	}
	if devicePatched.Name != newName {
		t.Errorf("Expected Test Device 2, got %s", devicePatched.Name)
	}
	if devicePatched.DeviceBrand != newDeviceBrand {
		t.Errorf("Expected Test Brand 2, got %s", devicePatched.DeviceBrand)
	}
}

func TestShouldDeleteDevice(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	device := &model.NewDeviceRequest{
		Name:        "Test Device",
		DeviceBrand: "Test Brand",
	}

	id, err := deviceService.AddDevice(ctx, device)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if id == nil {
		t.Errorf("Expected id, got nil")
	}

	deviceFound, err := deviceService.GetDevice(ctx, *id)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if deviceFound == nil {
		t.Errorf("Expected device, got nil")
	}
	if deviceFound.ID != *id {
		t.Errorf("Expected %s, got %s", *id, deviceFound.ID)
	}

	err = deviceService.DeleteDevice(ctx, *id)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	deviceFound, err = deviceService.GetDevice(ctx, *id)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if deviceFound != nil {
		t.Errorf("Expected nil, got %s", deviceFound.ID)
	}
}

func TestShouldNotGetDevice(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	deviceFound, err := deviceService.GetDevice(ctx, "invalid_id")
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if deviceFound != nil {
		t.Errorf("Expected nil, got %s", deviceFound.ID)
	}
}

func TestShouldNotPatchDevice(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	newName := "Test Device 2"
	newDeviceBrand := "Test Brand 2"

	patchDevice := &model.PatchDeviceRequest{
		ID:          "invalid_id",
		Name:        &newName,
		DeviceBrand: &newDeviceBrand,
	}

	idPatched, err := deviceService.PatchDevice(ctx, patchDevice)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if idPatched != nil {
		t.Errorf("Expected nil, got %s", *idPatched)
	}
}

func TestShoudlSearchDevicr(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	device := &model.NewDeviceRequest{
		Name:        "Test Device",
		DeviceBrand: "Test Brand",
	}

	id, err := deviceService.AddDevice(ctx, device)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if id == nil {
		t.Errorf("Expected id, got nil")
	}

	deviceFound, err := deviceService.GetDevice(ctx, *id)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if deviceFound == nil {
		t.Errorf("Expected device, got nil")
	}
	if deviceFound.ID != *id {
		t.Errorf("Expected %s, got %s", *id, deviceFound.ID)
	}

	newName := "Test Device 2"
	newDeviceBrand := "Test Brand 2"

	patchDevice := &model.PatchDeviceRequest{
		ID:          *id,
		Name:        &newName,
		DeviceBrand: &newDeviceBrand,
	}

	idPatched, err := deviceService.PatchDevice(ctx, patchDevice)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if idPatched == nil {
		t.Errorf("Expected id, got nil")
	}

	devicePatched, err := deviceService.GetDevice(ctx, *idPatched)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if devicePatched == nil {
		t.Errorf("Expected device, got nil")
	}
	if devicePatched.ID != *id {
		t.Errorf("Expected %s, got %s", *id, devicePatched.ID)
	}
	if devicePatched.Name != newName {
		t.Errorf("Expected Test Device 2, got %s", devicePatched.Name)
	}
	if devicePatched.DeviceBrand != newDeviceBrand {
		t.Errorf("Expected Test Brand 2, got %s", devicePatched.DeviceBrand)
	}
}
