package resolver

import (
	"devices_crud/internal/devices/app"
	"log"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DeviceService *app.DeviceService
	Logger        *log.Logger
}
