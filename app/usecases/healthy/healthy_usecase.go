package usecases

import (
	gateways "GoMetricsExporter/app/gateways/healthy"
	"fmt"
	"sync"
)

var (
	once          sync.Once
	healthy_check gateways.HealthyCheckGateway
)

const (
	MAX_WORKERS = 10
)

func HealthyCheckUseCase(
	healthy_check gateways.HealthyCheckGateway,
) error {
	once.Do(func() { initialize(healthy_check) })
	return nil
}

func HealthyCheck() {
	fmt.Println(healthy_check.HealthyCheck())

}

func initialize(healthy_check_repo gateways.HealthyCheckGateway) {
	healthy_check = healthy_check_repo

	for i := 0; i < 1; i++ {
		go HealthyCheck()

	}
	println("saiu do for")
}
