package main

import (
	"encoding/json"
	"os"

	"github.com/bakhtik/sms-docker/internal/app/sms-docker/model"
	"github.com/bakhtik/sms-docker/internal/app/sms-docker/server"
	"github.com/bakhtik/sms-docker/internal/pkg/jsonconfig"
)

func main() {
	// Load the configuration file
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

	server.Run(config.Server, model.InitRedisCache(config.Cache))
}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Cache  model.CacheConfig   `json:"Cache"`
	Server server.ServerConfig `json:"Server"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
