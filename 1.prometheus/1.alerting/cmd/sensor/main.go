package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	prometheus.MustRegister(cpuTempSensorReadingsTotal)
	prometheus.MustRegister(cpuTempSensor)
	prometheus.MustRegister(memSensorReadingsTotal)
	prometheus.MustRegister(memSensor)

	go cpuTemperatureSensor()
	go memTemperatureSensor()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8081", r)
}

const (
	cpuTemp  = 100
	memUsage = 100
)

func randomMax(max int) int {
	min := 90

	return rand.Intn(max-min) + min
}

var (
	cpuTempSensorReadingsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "cpu_temperature_sensor_readings_total",
		Help: "Total number of readings from the CPU temperature sensor.",
	})

	cpuTempSensor = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
)

func cpuTemperatureSensor() {
	for {
		cpuTempSensorReadingsTotal.Inc()
		cpuTempSensor.Set(float64(randomMax(cpuTemp)))

		time.Sleep(time.Second)
	}
}

var (
	memSensorReadingsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "mem_sensor_readings_total",
		Help: "Total number of readings from the CPU temperature sensor.",
	})

	memSensor = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "mem_usage_percentage",
		Help: "Current temperature of the CPU.",
	})
)

func memTemperatureSensor() {
	for {
		memSensorReadingsTotal.Inc()
		memSensor.Set(float64(randomMax(cpuTemp)))

		time.Sleep(time.Second)
	}
}
