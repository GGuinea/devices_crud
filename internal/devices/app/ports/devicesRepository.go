package ports

import (
	"context"
	"devices_crud/internal/devices/domain"
)

type DevicesRepository interface {
	Save(ctx context.Context, device *domain.Device) (*domain.Device, error)
	FindByID(ctx context.Context, id *string) (*domain.Device, error)
	FindAll(ctx context.Context) ([]domain.Device, error)
	Replace(ctx context.Context, device *domain.Device) (*domain.Device, error)
	Patch(ctx context.Context, device *domain.Device) (*domain.Device, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, query string) ([]domain.Device, error)
}
