package main

import (
	"github.com/Atluss/Go-Nats-Api-Example/pkg/v1"
	"github.com/Atluss/Go-Nats-Api-Example/pkg/v1/api"
	"github.com/Atluss/Go-Nats-Api-Example/pkg/v1/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	settingPath := "settings.json"

	set := config.NewApiSetup(settingPath)
	set.Print()

	// do something if user close program (close DB, or wait running query)
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Exit program...")
		os.Exit(1)
	}()

	// setup nats queue for test request
	endP, err := api.NewEndPointV1Test(set)
	if v1.LogOnError(err, "warning") {
		log.Printf("Setup endpoint: %s", endP.Url)
	}

	log.Fatal(http.ListenAndServe(":8080", set.Route))

}
