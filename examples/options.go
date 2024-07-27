package main

import (
	"flag"
	"fmt"

	"github.com/sondalex/gottings"
)

type Config struct {
	Host string `json:"host" env:"APP_HOST"`
	Port int    `json:"port" env:"APP_PORT"`
}

var port int

var flags map[string]interface{} = map[string]interface{}{
	"Port": &port,
}

func init() {
	flag.IntVar(flags["Port"].(*int), "port", 1312, "Port to be used by the Server")
}

func NewConfig() (*Config, error) {
	config := Config{}
	err := gottings.LoadOptions(flags, &config)
	if err != nil {
		return nil, nil
	}
	err = gottings.LoadConfiguration([]byte(`{"host": "127.0.0.1"}`), &config)
	if err != nil {
		return nil, nil
	}
	return &config, nil
}
func main() {
	config, err := NewConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}
