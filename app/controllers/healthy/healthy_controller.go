package controllers

import (
	"log"
	"time"

	usecases "GoMetricsExporter/app/usecases"
	"GoMetricsExporter/infrastructure"
	"GoMetricsExporter/infrastructure/database/healthy"
	"GoMetricsExporter/infrastructure/database/mysql/repositories"
)

var (
	controlGateway = repositories.ControlRepository{} // Certifique-se de que ControlRepository implementa ControlGateway
	healthyRepo    = healthy.HealthyRepository{}
	env            = infrastructure.GetEnvironmentVariables()
)

func Run() {
	// Inicializa o ControlUseCase com o ControlGateway
	err := usecases.ControlUseCase(controlGateway)
	if err != nil {
		log.Fatalf("[CONTROLLER] Error initializing use case: %v\n", err.Error())
	}

	// Inicia um loop para manter o serviço ativo
	for {
		// Use o método HealthyCheck do ControlUseCase em um intervalo definido
		// Aqui você pode adicionar outras chamadas ou funcionalidades, se necessário
		time.Sleep(time.Second * time.Duration(env.ENV_TIME_KEEPALIVE_SECONDS))
	}
}
