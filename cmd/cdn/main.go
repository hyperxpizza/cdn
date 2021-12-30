package main

import (
	"flag"
	"log"

	"github.com/hyperxpizza/cdn/pkg/config"
	"github.com/hyperxpizza/cdn/pkg/grpc/impl"
	"github.com/hyperxpizza/cdn/pkg/router"
)

var grpcFlag = flag.Bool("grpc", false, "set true to run a grpc server")
var configPath = flag.String("config", "", "path to config.json file")

func main() {
	flag.Parse()

	if *configPath == "" {
		log.Fatal("Config path was not set, aborting...")
		return
	}

	c, err := config.NewConfig(*configPath)
	if err != nil {
		log.Fatalf("Could not load config from file: %s , error: %s", *configPath, err.Error())
	}

	if *grpcFlag {
		go func() {
			server, err := impl.NewCDNService(c)
			if err != nil {
				return
			}

			server.Run()
		}()
	}

	router.Run(c)

}
