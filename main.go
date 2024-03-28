package main

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gofm/db"
	"io"
	"log"
	"net/http"
)

func isOK(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "ok")
	if err != nil {
		return
	}
}

var (
	opsProcessedPop = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_processed_ops_pop",
	})
)

var (
	opsProcessedRap = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_processed_ops_rap",
	})
)

var (
	opsProcessedRock = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_processed_ops_rock",
	})
)

var (
	opsProcessedSlow = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_processed_ops_slow",
	})
)

var (
	opsProcessedGen = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_processed_ops_gen",
	})
)

func main() {
	db.Database()
	myRouter := mux.NewRouter().StrictSlash(true)
	internRouter := mux.NewRouter().StrictSlash(true)

	hubRapRadio := newHub("rap")
	go hubRapRadio.run()
	go hubRapRadio.sendMusique()
	hubPopRadio := newHub("pop")
	go hubPopRadio.run()
	go hubPopRadio.sendMusique()
	hubRockRadio := newHub("rock")
	go hubRockRadio.run()
	go hubRockRadio.sendMusique()
	hubSlowRadio := newHub("slow")
	go hubSlowRadio.run()
	go hubSlowRadio.sendMusique()
	hubGenRadio := newHub("gen")
	go hubGenRadio.run()
	go hubGenRadio.sendMusique()

	myRouter.HandleFunc("/rap", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hubRapRadio, w, r)
	}).Methods(http.MethodGet)
	myRouter.HandleFunc("/pop", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hubPopRadio, w, r)
	}).Methods(http.MethodGet)
	myRouter.HandleFunc("/rock", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hubRockRadio, w, r)
	}).Methods(http.MethodGet)
	myRouter.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hubSlowRadio, w, r)
	}).Methods(http.MethodGet)
	myRouter.HandleFunc("/gen", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hubGenRadio, w, r)
	}).Methods(http.MethodGet)

	internRouter.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)
	internRouter.HandleFunc("/ping", isOK).Methods(http.MethodGet)

	go func() {
		log.Fatal(http.ListenAndServe(":2112", internRouter))
	}()
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
