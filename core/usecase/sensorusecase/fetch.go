package sensorusecase

import (
	"backend/core/domain"
	"backend/core/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination[[]domain.Sensor], error) {
	Sensors, err := usecase.repository.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return Sensors, nil
}
