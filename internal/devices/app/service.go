package app

import (
	"context"
	"devices_crud/internal/devices/app/ports"
	"devices_crud/internal/devices/model"
	"log"
	"time"

	"github.com/google/uuid"
)

type DeviceService struct {
	DevicesRepository ports.DevicesRepository
	Logger            *log.Logger
}

func NewDeviceService(devicesRepository ports.DevicesRepository, logger *log.Logger) *DeviceService {
	return &DeviceService{
		DevicesRepository: devicesRepository,
		Logger:            logger,
	}
}

func (s *DeviceService) AddDevice(ctx context.Context, device *model.NewDeviceRequest) (*string, error) {
	newDevice := &model.Device{
		ID:          uuid.New().String(),
		Name:        device.Name,
		DeviceBrand: device.DeviceBrand,
		CreatedAt:   time.Now(),
	}

	id, err := s.DevicesRepository.Save(ctx, newDevice)
	if err != nil {
		return nil, err
	}

	return id, nil
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

func (s *DeviceService) PatchDevice(ctx context.Context, device *model.PatchDeviceRequest) (*string, error) {
	id, err := s.DevicesRepository.Patch(ctx, device)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (s *DeviceService) DeleteDevice(ctx context.Context, id string) error {
	return s.DevicesRepository.Delete(ctx, id)
}

func (s *DeviceService) SearchDevices(ctx context.Context, query string) ([]model.Device, error) {
	return s.DevicesRepository.Search(ctx, query)
}
