package external

import (
	"GoMetricsExporter/app/entities/mysql/BabyPetControl/model"
	"encoding/json"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	matricesGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "matrices",
			Help: "Número de matrizes",
		},
		[]string{"name"},
	) // Label para diferenciar as matrizes

	deviceCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "breedings",
		Help: "Number of active devices",
	})
)

func init() {
	prometheus.MustRegister(matricesGauge)
	prometheus.MustRegister(deviceCount)
}

func UpdateMetricsMatrices(matrices []model.Matrices) {
	fmt.Println(matrices)
	if len(matrices) > 0 {
		for _, matrix := range matrices {
			fmt.Println(matrix)
			jsonData, err := json.Marshal(matrix)
			if err != nil {
				fmt.Printf("Failed to marshal matrix to JSON: %v\n", err)
				continue
			}

			matricesGauge.WithLabelValues(string(jsonData)).Set(1)
		}

		if err := push.New("http://192.168.1.115:9091", "matrices_job").
			Collector(matricesGauge).
			Push(); err != nil {
			fmt.Printf("Unable to send matrix metrics to Pushgateway: %v\n", err)
		}
		fmt.Println("Matrix metrics updated.")
	}

}
