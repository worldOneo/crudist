package crudist

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/worldOneo/crudist/core"
	"github.com/worldOneo/crudist/errors"
)

// ReadConfig reads the configuration file
func ReadConfig(file string) (core.CrudConfig, error) {
	content, err := os.ReadFile(file)
	conf := core.CrudConfig{}
	if err != nil {
		return conf, errors.NewConfigReadError(file, err)
	}
	err = json.NewDecoder(bytes.NewReader(content)).Decode(&conf)
	if err != nil {
		return conf, errors.NewConfigParseError(err)
	}
	return conf, nil
}
