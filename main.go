package main

import (
	controllers "GoMetricsExporter/app/controllers/healthy"
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	log.Println("[CONSUMER] Starting Consumer MetricsExporter...")
	controllers.Run()
}
