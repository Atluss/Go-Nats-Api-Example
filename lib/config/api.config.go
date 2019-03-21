package config

import (
	"encoding/json"
	"fmt"
	"github.com/Atluss/Go-Nats-Api-Example/lib"
	"io/ioutil"
	"os"
)

// Config load new config for API
func Config(path string) (*config, error) {

	conf := config{}

	if err := lib.CheckFileExist(path); err != nil {
		return &conf, err
	}

	conf.FilePath = path

	if err := conf.load(); err != nil {
		return &conf, err
	}

	return &conf, nil
}

type natsConfig struct {
	Host string `json:"Host"`
	Port string `json:"Port"`
}

type config struct {
	Name     string     `json:"name"`     // API name
	Version  string     `json:"Version"`  // API version
	FilePath string     `json:"FilePath"` // path to Json settings file
	Nats     natsConfig `json:"Nats"`
}

// load all settings
func (obj *config) load() error {

	jsonSet, err := os.Open(obj.FilePath)

	defer lib.LogOnError(jsonSet.Close(), "warning: Can't close json settings file.")

	if !lib.LogOnError(err, "error: Can't open config file") {
		return err
	}

	bytesVal, _ := ioutil.ReadAll(jsonSet)
	err = json.Unmarshal(bytesVal, obj)

	if !lib.LogOnError(err, "error: Can't unmarshal json file") {
		return err
	}

	return obj.validate()
}

// validate it
func (obj *config) validate() error {

	if obj.Name == "" {
		return fmt.Errorf("error: config miss name")
	}

	if obj.Version == "" {
		return fmt.Errorf("error: config miss version")
	}

	if obj.Nats.Host == "" {
		return fmt.Errorf("error: config miss Nats host")
	}

	if obj.Nats.Port == "" {
		return fmt.Errorf("error: config miss Nats port")
	}

	return nil
}
