package main

import (
	"flag"
	"log"
	"sync"

	"github.com/hyperxpizza/cdn/pkg/config"
	"github.com/hyperxpizza/cdn/pkg/grpc/impl"
	"github.com/hyperxpizza/cdn/pkg/router"
)

var grpcFlag = flag.Bool("grpc", false, "set true to run a grpc server")
var configPath = flag.String("config", "", "path to config.json file")
var restFlag = flag.Bool("rest", true, "set true to run a rest server")
var secure = flag.Bool("secure", false, "secure grpc connection witr tls")

func main() {
	flag.Parse()

	if *configPath == "" {
		log.Fatal("Config path was not set, aborting...")
		return
	}

	c, err := config.NewConfig(*configPath)
	if err != nil {
		log.Fatalf("Could not load config from file: %s , error: %s", *configPath, err.Error())
		return
	}

	var wg sync.WaitGroup

	if *grpcFlag {
		wg.Add(1)
		go func() {
			server, err := impl.NewCDNService(c, *secure)
			if err != nil {
				wg.Done()
				return
			}

			server.Run()
		}()
	}

	if *restFlag {
		wg.Add(1)
		go func() {
			err := router.Run(c)
			if err != nil {
				wg.Done()
				return
			}
		}()
	}

	wg.Wait()
}
