package sensorservice

import (
	"encoding/json"
	"net/http"

	"backend/core/dto"
)

func (service service) Create(response http.ResponseWriter, request *http.Request) {
	SensorRequest, err := dto.FromJSONCreateSensorRequest(request.Body)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	Sensor, err := service.usecase.Create(SensorRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(Sensor)
}
