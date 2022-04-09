package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/AndreyGuznov/Avito_task2/internal/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	server := apiserver.NewServer(config)
	if err := server.StartServer(); err != nil {
		log.Fatal(err)
	}
}
