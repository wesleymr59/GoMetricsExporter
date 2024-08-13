package main

import (
	controllers "GoMetricsExporter/app/controllers/healthy"
	"GoMetricsExporter/infrastructure/database/mysql/config"
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	log.Println("[CONSUMER] Starting Consumer MetricsExporter...")
	config.InitDB()
	controllers.Run()
}
