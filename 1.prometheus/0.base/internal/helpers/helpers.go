package helpers

import (
	"math/rand"
	"net/http"
	"time"
)

func RandomDurationMS(maxMS int) time.Duration {
	minMS := 1

	return time.Duration(rand.Intn(maxMS-minMS)+minMS) * time.Millisecond
}

func Random2xx() int {
	statuses := []int{
		http.StatusOK,
		http.StatusAccepted,
	}

	index := rand.Intn(len(statuses))

	return statuses[index]
}

func Random4xx() int {
	statuses := []int{
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusNotFound,
		http.StatusTooManyRequests,
	}

	index := rand.Intn(len(statuses))

	return statuses[index]
}

func Random5xx() int {
	statuses := []int{
		http.StatusInternalServerError,
		http.StatusNotImplemented,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
	}

	index := rand.Intn(len(statuses))

	return statuses[index]
}
