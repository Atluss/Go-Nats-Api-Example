package config

import (
	"github.com/Atluss/Go-Nats-Api-Example/lib"
	"github.com/nats-io/go-nats"
	"time"
)

func NewApiSetup(settings string) *setup {

	cnf, err := Config(settings)
	lib.FailOnError(err, "error config file")

	set, err := Setup(cnf)
	lib.FailOnError(err, "error setup")

	return set
}

func Setup(cnf *config) (*setup, error) {

	set := setup{}

	if err := cnf.validate(); err != nil {
		return &set, err
	}

	set.Config = cnf

	if err := set.natsConnection(); err != nil {
		return &set, err
	}

	return &set, nil
}

// setup main setup api struct
type setup struct {
	Config *config    // api setting
	Nats   *nats.Conn // nats
}

// natsConnection setup nats
func (obj *setup) natsConnection() error {

	var err error

	if obj.Nats, err = nats.Connect(obj.Config.Nats.Address[0].Address, nats.MaxReconnects(-1), nats.ReconnectWait(time.Second*5)); err != nil {
		return err
	}

	return nil
}
