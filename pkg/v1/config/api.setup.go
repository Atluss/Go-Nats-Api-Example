package config

import (
	"github.com/Atluss/Go-Nats-Api-Example/pkg/v1"
	"github.com/gorilla/mux"
	"github.com/nats-io/go-nats"
	"log"
	"time"
)

// NewApiSetup return new setup struct with router and nats
func NewApiSetup(settings string) *Setup {

	cnf, err := Config(settings)
	v1.FailOnError(err, "error config file")

	set, err := newSetup(cnf)
	v1.FailOnError(err, "error setup")

	return set
}

func newSetup(cnf *config) (*Setup, error) {

	set := Setup{}

	if err := cnf.validate(); err != nil {
		return &set, err
	}

	set.Config = cnf

	if err := set.natsConnection(); err != nil {
		return &set, err
	}

	set.Route = mux.NewRouter().StrictSlash(true)

	return &set, nil
}

// Setup main setup api struct
type Setup struct {
	Config *config     // api setting
	Nats   *nats.Conn  // nats
	Route  *mux.Router // mux frontend
}

// natsConnection setup nats
func (obj *Setup) natsConnection() error {

	var err error

	log.Println(obj.Config.Nats.Address[0].Address)

	if obj.Nats, err = nats.Connect(obj.Config.Nats.Address[0].Address, nats.MaxReconnects(-1), nats.ReconnectWait(time.Second*5)); err != nil {
		return err
	}

	return nil
}

func (obj *Setup) Print() {
	log.Printf("Name: %s", obj.Config.Name)
	log.Printf("Version: %s", obj.Config.Version)
	log.Printf("Nats version: %s", obj.Config.Nats.Version)
	log.Printf("Nats ReconnectedWait: %d", obj.Config.Nats.ReconnectedWait)
	log.Printf("Nats host: %s", obj.Config.Nats.Address[0].Host)
	log.Printf("Nats port: %s", obj.Config.Nats.Address[0].Port)
	log.Printf("Nats address: %s", obj.Config.Nats.Address[0].Address)
	log.Printf("Nats address(multi): %s", obj.Config.GetNatsAddresses())
}
