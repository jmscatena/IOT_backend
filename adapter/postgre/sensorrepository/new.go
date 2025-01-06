package sensorrepository

import (
	"backend/adapter/postgre"
	"backend/core/domain"
)

type repository struct {
	db postgre.PoolInterface
}

// New returns contract implementation of SensorRepository
func New(db postgre.PoolInterface) domain.SensorRepository {
	return &repository{
		db: db,
	}
}
