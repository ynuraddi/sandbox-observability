package main

import (
	"math/rand/v2"
	"net/http"

	"golang.org/x/sync/errgroup"
)

const target = "http://localhost:8080"

var (
	endpoints = []string{
		"code/2xx",
		"code/4xx",
		"code/5xx",
	}

	durations = []string{
		"100",
		"250",
		"500",
		"1000",
	}
)

const (
	success = 95
	short   = 999
)

func main() {
	g := errgroup.Group{}
	g.SetLimit(2)

	for {
		seed := rand.IntN(100)
		if seed <= success {
			g.Go(func() error {
				return sendRequest(endpoints[0])
			})
		} else {
			g.Go(func() error {
				return sendRequest(endpoints[rand.IntN(2)+1])
			})
		}
	}
}

func sendRequest(path string) error {
	var duration string
	seed := rand.IntN(1000)
	if seed <= short {
		duration = durations[rand.IntN(2)]
	} else {
		duration = durations[rand.IntN(2)+2]
	}

	_, err := http.Get(target + "/" + path + "/" + duration)
	if err != nil {
		panic(err)
	}
	return nil
}
