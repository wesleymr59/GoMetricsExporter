package healthy

import entities "GoMetricsExporter/app/entities/healthy"

type HealthyRepository struct{}

func (x HealthyRepository) HealthyCheck() entities.HealthyCheck {
	return entities.HealthyCheck{
		Message: "OK",
	}
}
