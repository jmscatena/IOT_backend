package dto

import (
	"encoding/json"
	"io"
)

// CreateSensorRequest is an representation request body to create a new Sensor
type CreateSensorRequest struct {
	Device      string  `json:"device"`
	Corporation float32 `json:"corp"`
	Description string  `json:"description"`
}

// FromJSONCreateSensorRequest converts json body request to a CreateSensorRequest struct
func FromJSONCreateSensorRequest(body io.Reader) (*CreateSensorRequest, error) {
	createSensorRequest := CreateSensorRequest{}
	if err := json.NewDecoder(body).Decode(&createSensorRequest); err != nil {
		return nil, err
	}
	return &createSensorRequest, nil
}
