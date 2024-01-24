package domain

type Device struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DeviceBrand string `json:"deviceBrand"`
	CreatedAt   string `json:"createdAt"`
}
