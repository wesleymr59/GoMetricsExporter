package infrastructure

import (
	"os"
	"strconv"
	"sync"
)

type EnvironmentVariables struct {
	ENV_TIME_KEEPALIVE_SECONDS int
}

var (
	env  *EnvironmentVariables = nil
	once sync.Once
)

func loadEnviromentVariables() {
	ENV_TIME_KEEPALIVE_SECONDS, _ := strconv.Atoi(os.Getenv("ENV_TIME_KEEPALIVE_SECONDS"))
	env = &EnvironmentVariables{

		ENV_TIME_KEEPALIVE_SECONDS: ENV_TIME_KEEPALIVE_SECONDS,
	}
}

func GetEnvironmentVariables() *EnvironmentVariables {
	once.Do(loadEnviromentVariables)
	return env
}
