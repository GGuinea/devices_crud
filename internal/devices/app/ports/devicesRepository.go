package ports

import (
	"context"
	"devices_crud/internal/devices/model"

	"strings"
)

type DevicesRepository interface {
	Save(ctx context.Context, device *model.Device) (*string, error)
	FindByID(ctx context.Context, id *string) (*model.Device, error)
	FindAll(ctx context.Context) ([]model.Device, error)
	Replace(ctx context.Context, device *model.Device) (*model.Device, error)
	Patch(ctx context.Context, device *model.PatchDeviceRequest) (*string, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, query string) ([]model.Device, error)
}

var DevicesContainer map[string]model.Device

type devicesRepositoryMock struct {
}

func NewDevicesRepositoryMock() DevicesRepository {
	DevicesContainer = make(map[string]model.Device)
	return &devicesRepositoryMock{}
}

func (r *devicesRepositoryMock) Save(ctx context.Context, device *model.Device) (*string, error) {
	DevicesContainer[device.ID] = *device
	return &device.ID, nil
}

func (r *devicesRepositoryMock) FindByID(ctx context.Context, id *string) (*model.Device, error) {
	device, ok := DevicesContainer[*id]
	if !ok {
		return nil, nil
	}
	return &device, nil
}

func (r *devicesRepositoryMock) FindAll(ctx context.Context) ([]model.Device, error) {
	devices := make([]model.Device, 0)
	for _, device := range DevicesContainer {
		devices = append(devices, device)
	}
	return devices, nil
}

func (r *devicesRepositoryMock) Replace(ctx context.Context, device *model.Device) (*model.Device, error) {
	DevicesContainer[device.ID] = *device
	return device, nil
}

func (r *devicesRepositoryMock) Patch(ctx context.Context, device *model.PatchDeviceRequest) (*string, error) {
	deviceToPatch, ok := DevicesContainer[device.ID]
	if !ok {
		return nil, nil
	}

	if device.Name != nil {
		deviceToPatch.Name = *device.Name
	}
	if device.DeviceBrand != nil {
		deviceToPatch.DeviceBrand = *device.DeviceBrand
	}

	DevicesContainer[device.ID] = deviceToPatch
	return &device.ID, nil
}

func (r *devicesRepositoryMock) Delete(ctx context.Context, id string) error {
	delete(DevicesContainer, id)
	return nil
}

func (r *devicesRepositoryMock) Search(ctx context.Context, query string) ([]model.Device, error) {
	var devices []model.Device
	for _, device := range DevicesContainer {
		if strings.ToLower(device.DeviceBrand) == strings.ToLower(query) {
			devices = append(devices, device)
		}
	}
	return devices, nil
}
