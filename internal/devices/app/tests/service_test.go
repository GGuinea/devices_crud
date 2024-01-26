package tests

import (
	"context"
	"devices_crud/internal/devices/app"
	"devices_crud/internal/devices/app/ports"
	"devices_crud/internal/devices/model"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, id)
}

func TestShouldGetDevice(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	device := &model.NewDeviceRequest{
		Name:        "Test Device",
		DeviceBrand: "Test Brand",
	}

	id, err := deviceService.AddDevice(ctx, device)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, id)

	deviceFound, err := deviceService.GetDevice(ctx, *id)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, deviceFound)
	assert.Equal(t, *id, deviceFound.ID)
	assert.Equal(t, "Test Device", deviceFound.Name)
	assert.Equal(t, "Test Brand", deviceFound.DeviceBrand)
}

func TestShouldGetAllDevices(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	device := &model.NewDeviceRequest{
		Name:        "Test Device",
		DeviceBrand: "Test Brand",
	}

	_, err := deviceService.AddDevice(ctx, device)
	assert.Equal(t, nil, err)

	device_2 := &model.NewDeviceRequest{
		Name:        "Test Device 2",
		DeviceBrand: "Test Brand 2",
	}

	_, err = deviceService.AddDevice(ctx, device_2)
	assert.Equal(t, nil, err)

	devices, err := deviceService.GetAllDevices(ctx)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, devices)
	assert.Equal(t, 2, len(devices))

	foundDevicesNames := make([]string, 0)
	foundDevicesBrands := make([]string, 0)
	for _, device := range devices {
		foundDevicesNames = append(foundDevicesNames, device.Name)
		foundDevicesBrands = append(foundDevicesBrands, device.DeviceBrand)
	}

	assert.Contains(t, foundDevicesNames, "Test Device")
	assert.Contains(t, foundDevicesNames, "Test Device 2")
	assert.Contains(t, foundDevicesBrands, "Test Brand")
	assert.Contains(t, foundDevicesBrands, "Test Brand 2")
}

func TestShouldReplaceDevice(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	device := &model.NewDeviceRequest{
		Name:        "Test Device",
		DeviceBrand: "Test Brand",
	}

	id, err := deviceService.AddDevice(ctx, device)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, id)

	deviceFound, err := deviceService.GetDevice(ctx, *id)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, deviceFound)
	assert.Equal(t, *id, deviceFound.ID)
	assert.Equal(t, "Test Device", deviceFound.Name)
	assert.Equal(t, "Test Brand", deviceFound.DeviceBrand)

	deviceFound.Name = "Test Device 2"
	deviceFound.DeviceBrand = "Test Brand 2"

	deviceReplaced, err := deviceService.ReplaceDevice(ctx, deviceFound)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, deviceReplaced)
	assert.Equal(t, *id, deviceReplaced.ID)
	assert.Equal(t, "Test Device 2", deviceReplaced.Name)
	assert.Equal(t, "Test Brand 2", deviceReplaced.DeviceBrand)
}

func TestShouldPatchDevice(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	device := &model.NewDeviceRequest{
		Name:        "Test Device",
		DeviceBrand: "Test Brand",
	}

	id, err := deviceService.AddDevice(ctx, device)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, id)

	deviceFound, err := deviceService.GetDevice(ctx, *id)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, deviceFound)
	assert.Equal(t, *id, deviceFound.ID)

	newName := "Test Device 2"
	newDeviceBrand := "Test Brand 2"

	patchDevice := &model.PatchDeviceRequest{
		ID:          *id,
		Name:        &newName,
		DeviceBrand: &newDeviceBrand,
	}

	idPatched, err := deviceService.PatchDevice(ctx, patchDevice)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, idPatched)

	devicePatched, err := deviceService.GetDevice(ctx, *idPatched)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, devicePatched)
	assert.Equal(t, *id, devicePatched.ID)
	assert.Equal(t, "Test Device 2", devicePatched.Name)
	assert.Equal(t, "Test Brand 2", devicePatched.DeviceBrand)
}

func TestShouldDeleteDevice(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	device := &model.NewDeviceRequest{
		Name:        "Test Device",
		DeviceBrand: "Test Brand",
	}

	id, err := deviceService.AddDevice(ctx, device)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, id)

	deviceFound, err := deviceService.GetDevice(ctx, *id)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, deviceFound)
	assert.Equal(t, *id, deviceFound.ID)

	err = deviceService.DeleteDevice(ctx, *id)
	assert.Equal(t, nil, err)

	deviceFound, err = deviceService.GetDevice(ctx, *id)
	assert.Equal(t, nil, err)
	assert.Nil(t, deviceFound)
}

func TestShouldNotGetDevice(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	deviceFound, err := deviceService.GetDevice(ctx, "invalid_id")
	assert.Equal(t, nil, err)
	assert.Nil(t, deviceFound)
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

	assert.Equal(t, nil, err)
	assert.Nil(t, idPatched)
}

func TestShouldSearchDevicesByBrand(t *testing.T) {
	deviceService := getDeviceService()
	ctx := context.Background()

	device := &model.NewDeviceRequest{
		Name:        "Test Device",
		DeviceBrand: "Test Brand",
	}

	_, err := deviceService.AddDevice(ctx, device)
	assert.Equal(t, nil, err)

	device_2 := &model.NewDeviceRequest{
		Name:        "Test Device 2",
		DeviceBrand: "Test Brand 2",
	}

	_, err = deviceService.AddDevice(ctx, device_2)
	assert.Equal(t, nil, err)

	devices, err := deviceService.SearchDevices(ctx, "Test Brand")

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, devices)
	assert.Equal(t, 2, len(devices))

	foundDevicesNames := make([]string, 0)
	foundDevicesBrands := make([]string, 0)
	for _, device := range devices {
		foundDevicesNames = append(foundDevicesNames, device.Name)
		foundDevicesBrands = append(foundDevicesBrands, device.DeviceBrand)
	}
	assert.Contains(t, foundDevicesNames, "Test Device")
	assert.Contains(t, foundDevicesNames, "Test Device 2")
	assert.Contains(t, foundDevicesBrands, "Test Brand")
	assert.Contains(t, foundDevicesBrands, "Test Brand 2")
}
