package usecases

import (
	"GoMetricsExporter/app/gateways"
	external "GoMetricsExporter/infrastructure/external/prometheus"
	"fmt"
	"sync"
)

var (
	controlGateway gateways.ControlGateway
	once           sync.Once
)

func ControlUseCase(gateway gateways.ControlGateway) error {
	once.Do(func() { initialize(gateway) })
	return nil
}

func GetMatrizes() {
	matrices, err := controlGateway.GetMatrizes()
	if err != nil {
		fmt.Println("Error fetching matrizes:", err)
		return
	}
	fmt.Println(matrices)

	external.UpdateMetricsMatrices(matrices)
}

func GetPadreadores() {
	matrices, err := controlGateway.GetPadreadores()
	if err != nil {
		fmt.Println("Error fetching matrizes:", err)
		return
	}
	fmt.Println(matrices)

}

// initialize configura o gateway e inicia as goroutines necessárias
func initialize(gateway gateways.ControlGateway) {
	controlGateway = gateway

	go GetMatrizes()
	go GetPadreadores()

	println("saiu do for")
}
