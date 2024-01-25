package devices

import (
	"devices_crud/internal/devices/app"
	"devices_crud/internal/devices/app/ports"
)

type DeviceDependencies struct {
	UseMocks bool
}

type DependencyTree struct {
	DeviceSerivce *app.DeviceService
}

func NewDevicesDependencies(deps *DeviceDependencies) *DependencyTree {
	if deps == nil {
		panic("devices dependencies cannot be nil")
	}

	var service *app.DeviceService

	if deps.UseMocks {
		service = app.NewDeviceService(ports.NewDevicesRepositoryMock())
	} else {
		panic("We don't have a real implementation yet")
	}

	return &DependencyTree{
		DeviceSerivce: service,
	}

}
