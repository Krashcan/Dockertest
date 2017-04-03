package main

import(
	"log"
	"os"
	"os/signal"
	"syscall"
	"net/http"

	"github.com/krashcan/dockertest/router"
)

func main(){
	router := router.NewRouter()

	log.Println("Now Listening ...... ")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<- c
		log.Println("Stopping Server ..... ")
		os.Exit(1)
	}()

	log.Fatal(http.ListenAndServe(":8080", router))
}