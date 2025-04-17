package main

import (
	"net/http"
	"prom/internal/helpers"
	"prom/internal/middleware"
	"runtime"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Config struct {
}

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	router := chi.NewRouter()
	router.Use(chimiddleware.Logger)
	router.Use(chimiddleware.Recoverer)
	router.Use(middleware.HttpMetrics)

	router.Get("/code/2xx/{duration}", func(w http.ResponseWriter, r *http.Request) {
		sleep(r)
		w.WriteHeader(helpers.Random2xx())
	})

	router.Get("/code/4xx/{duration}", func(w http.ResponseWriter, r *http.Request) {
		sleep(r)
		w.WriteHeader(helpers.Random4xx())
	})

	router.Get("/code/5xx/{duration}", func(w http.ResponseWriter, r *http.Request) {
		sleep(r)
		w.WriteHeader(helpers.Random5xx())
	})

	router.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", router)
}

func sleep(r *http.Request) {
	durationStr := chi.URLParam(r, "duration")
	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		panic(err)
	}
	time.Sleep(helpers.RandomDurationMS(duration))
}
