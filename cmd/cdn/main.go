package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/hyperxpizza/cdn/pkg/config"
	"github.com/hyperxpizza/cdn/pkg/database"
	pb "github.com/hyperxpizza/cdn/pkg/grpc"
	"github.com/hyperxpizza/cdn/pkg/grpc/impl"
	"github.com/hyperxpizza/cdn/pkg/router"
	"google.golang.org/grpc"
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
			grpcServer := grpc.NewServer()

			db, err := database.NewDatabase(c)
			if err != nil {
				log.Println(err)
				return
			}

			server := impl.NewCDNService(c, db)

			pb.RegisterCDNGrpcServiceServer(grpcServer, server)

			addr := fmt.Sprintf(":%d", c.Grpc.Port)
			lis, err := net.Listen("tcp", addr)
			if err != nil {
				log.Println(err)
				return
			}

			log.Printf("grpc server running on %s:%d", c.Grpc.Host, c.Grpc.Port)
			if err := grpcServer.Serve(lis); err != nil {
				log.Println("Failed to serve:", err.Error())
				return
			}
		}()
	}

	router.Run(c)

}
