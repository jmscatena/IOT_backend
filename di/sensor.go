package di

import (
	"backend/adapter/http/sensorservice"
	"backend/adapter/postgre"
	"backend/adapter/postgre/sensorrepository"
	"backend/core/domain"
	"backend/core/usecase/sensorusecase"
)

// ConfigSensorDI return a SensorService abstraction with dependency injection configuration
func ConfigSensorDI(conn postgre.PoolInterface) domain.SensorService {
	SensorRepository := sensorrepository.New(conn)
	SensorUseCase := sensorusecase.New(SensorRepository)
	SensorService := sensorservice.New(SensorUseCase)

	return SensorService
}
