package main

import (
	"github.com/gorilla/mux"
	"gofm/db"
	"log"
	"net/http"
)

func main() {
	db.Database()
	myRouter := mux.NewRouter().StrictSlash(true)
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
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
