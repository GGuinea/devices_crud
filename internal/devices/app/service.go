package app

import (
	"context"
	"devices_crud/internal/devices/app/ports"
	"devices_crud/internal/devices/model"
)

type DeviceService struct {
	DevicesRepository ports.DevicesRepository
}

func NewDeviceService(devicesRepository ports.DevicesRepository) *DeviceService {
	return &DeviceService{
		DevicesRepository: devicesRepository,
	}
}

func (s *DeviceService) AddDevice(ctx context.Context, device *model.Device) (*model.Device, error) {
	return s.DevicesRepository.Save(ctx, device)
}

func (s *DeviceService) GetDevice(ctx context.Context, id string) (*model.Device, error) {
	return s.DevicesRepository.FindByID(ctx, &id)
}

func (s *DeviceService) GetAllDevices(ctx context.Context) ([]model.Device, error) {
	return s.DevicesRepository.FindAll(ctx)
}

func (s *DeviceService) ReplaceDevice(ctx context.Context, device *model.Device) (*model.Device, error) {
	return s.DevicesRepository.Replace(ctx, device)
}

func (s *DeviceService) PatchDevice(ctx context.Context, device *model.Device) (*model.Device, error) {
	return s.DevicesRepository.Patch(ctx, device)
}

func (s *DeviceService) DeleteDevice(ctx context.Context, id string) error {
	return s.DevicesRepository.Delete(ctx, id)
}

func (s *DeviceService) SearchDevices(ctx context.Context, query string) ([]model.Device, error) {
	return s.DevicesRepository.Search(ctx, query)
}
