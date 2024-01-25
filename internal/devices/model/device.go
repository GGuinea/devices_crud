package model

import "time"

type Device struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	DeviceBrand string    `json:"deviceBrand"`
	CreatedAt   time.Time `json:"createdAt"`
}

type NewDeviceRequest struct {
	Name        string `json:"name"`
	DeviceBrand string `json:"deviceBrand"`
}

type NewDeviceResponse struct {
	UUID string `json:"uuid"`
}
