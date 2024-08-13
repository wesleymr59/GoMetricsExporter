package gateways

import "GoMetricsExporter/app/entities/mysql/BabyPetControl/model"

type ControlGateway interface {
	GetMatrizes() ([]model.Matrices, error)
	GetPadreadores() ([]model.Matrices, error)
}
