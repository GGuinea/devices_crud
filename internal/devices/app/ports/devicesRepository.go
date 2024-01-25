package ports

import (
	"context"
	"devices_crud/internal/devices/model"

	"strings"
)

type DevicesRepository interface {
	Save(ctx context.Context, device *model.Device) (*model.Device, error)
	FindByID(ctx context.Context, id *string) (*model.Device, error)
	FindAll(ctx context.Context) ([]model.Device, error)
	Replace(ctx context.Context, device *model.Device) (*model.Device, error)
	Patch(ctx context.Context, device *model.Device) (*model.Device, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, query string) ([]model.Device, error)
}

type devicesRepositoryMock struct {
	Devices map[string]model.Device
}

func NewDevicesRepositoryMock() DevicesRepository {
	return &devicesRepositoryMock{
		Devices: make(map[string]model.Device),
	}
}

func (r *devicesRepositoryMock) Save(ctx context.Context, device *model.Device) (*model.Device, error) {
	r.Devices[device.ID] = *device
	return device, nil
}

func (r *devicesRepositoryMock) FindByID(ctx context.Context, id *string) (*model.Device, error) {
	device, ok := r.Devices[*id]
	if !ok {
		return nil, nil
	}
	return &device, nil
}

func (r *devicesRepositoryMock) FindAll(ctx context.Context) ([]model.Device, error) {
	devices := make([]model.Device, 0)
	for _, device := range r.Devices {
		devices = append(devices, device)
	}
	return devices, nil
}

func (r *devicesRepositoryMock) Replace(ctx context.Context, device *model.Device) (*model.Device, error) {
	r.Devices[device.ID] = *device
	return device, nil
}

func (r *devicesRepositoryMock) Patch(ctx context.Context, device *model.Device) (*model.Device, error) {
	r.Devices[device.ID] = *device
	return device, nil
}

func (r *devicesRepositoryMock) Delete(ctx context.Context, id string) error {
	delete(r.Devices, id)
	return nil
}

func (r *devicesRepositoryMock) Search(ctx context.Context, query string) ([]model.Device, error) {
	var devices []model.Device
	for _, device := range r.Devices {
		if strings.ToLower(device.DeviceBrand) == strings.ToLower(query) {
			devices = append(devices, device)
		}
	}
	return devices, nil
}
