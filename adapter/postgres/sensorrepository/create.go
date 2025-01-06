package sensorrepository

import (
	"backend/core/domain"
	"backend/core/dto"
	"context"
)

func (repository repository) Create(SensorRequest *dto.CreateSensorRequest) (*domain.Sensor, error) {
	ctx := context.Background()
	Sensor := domain.Sensor{}

	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO Sensor (name, price, description) VALUES ($1, $2, $3) returning *",
		SensorRequest.Name,
		SensorRequest.Price,
		SensorRequest.Description,
	).Scan(
		&Sensor.ID,
		&Sensor.Name,
		&Sensor.Price,
		&Sensor.Description,
	)

	if err != nil {
		return nil, err
	}

	return &Sensor, nil
}
