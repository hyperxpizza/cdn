package tests

import (
	"log"
	"testing"

	"github.com/hyperxpizza/cdn/pkg/config"
	pb "github.com/hyperxpizza/cdn/pkg/grpc"
	"github.com/hyperxpizza/cdn/pkg/grpc/impl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const buffer = 1024 * 1024

var lis *bufconn.Listener

func mockGrpcServer(c *config.Config) error {
	lis = bufconn.Listen(buffer)
	server := grpc.NewServer()

	cdnServiceServer, err := impl.NewCDNService(c)
	if err != nil {
		return err
	}

	pb.RegisterCDNGrpcServiceServer(server, cdnServiceServer)

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatal(err)
			return
		}
	}()

	return nil
}

func TestSearchFile(t *testing.T) {}

func TestUploadFile(t *testing.T) {}

func TestDownloadFile(t *testing.T) {}
