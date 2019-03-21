package config

import (
	"github.com/Atluss/Go-Nats-Api-Example/lib"
	"log"
	"testing"
)

func TestConfig(t *testing.T) {

	path := "settings.json"

	cnf, err := Config(path)
	lib.FailOnError(err, "Test error")

	log.Printf("%+v", cnf)
	log.Printf("Name: %s", cnf.Name)
	log.Printf("Version: %s", cnf.Version)
	log.Printf("Nats host: %s", cnf.Nats.Host)
	log.Printf("Nats port: %s", cnf.Nats.Port)

}
