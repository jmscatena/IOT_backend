package domain

import (
	"backend/core/dto"
	"net/http"
)

// Sensor is entity of table Sensor database column
type Sensor struct {
	ID          int32   `json:"id"`
	Device      string  `json:"device"`
	Corporation float32 `json:"corp"`
	Description string  `json:"description"`
}

// SensorService is a contract of http adapter layer
type SensorService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
}

// SensorUseCase is a contract of business rule layer
type SensorUseCase interface {
	Create(SensorRequest *dto.CreateSensorRequest) (*Sensor, error)
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination[[]Sensor], error)
}

// SensorRepository is a contract of database connection adapter layer
type SensorRepository interface {
	Create(SensorRequest *dto.CreateSensorRequest) (*Sensor, error)
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination[[]Sensor], error)
}
