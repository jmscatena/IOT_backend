package sensorrepository

import (
	"backend/adapter/postgres"
	"backend/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of SensorRepository
func New(db postgres.PoolInterface) domain.SensorRepository {
	return &repository{
		db: db,
	}
}
