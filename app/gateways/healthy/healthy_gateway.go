package gateways

import entities "GoMetricsExporter/app/entities/healthy"

type HealthyCheckGateway interface {
	HealthyCheck() entities.HealthyCheck
}
