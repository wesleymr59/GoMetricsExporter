package controllers

import (
	usecases "GoMetricsExporter/app/usecases/healthy"
	"GoMetricsExporter/infrastructure"
	"GoMetricsExporter/infrastructure/database/healthy"
	"log"
	"time"
)

var (
	healthy_repo = healthy.HealthyRepository{}
	env          = infrastructure.GetEnvironmentVariables()
)

func Run() {
	var err error
	for {
		err = usecases.HealthyCheckUseCase(healthy_repo)

		if err != nil {
			log.Printf("[CONSUMER] Error: %v\n", err.Error())
		}

		time.Sleep(time.Second * time.Duration(env.ENV_TIME_KEEPALIVE_SECONDS))
	}
}
