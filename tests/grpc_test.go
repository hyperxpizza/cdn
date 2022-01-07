package tests

import (
	"flag"
	"log"
	"testing"

	"github.com/hyperxpizza/cdn/pkg/config"
	pb "github.com/hyperxpizza/cdn/pkg/grpc"
	"github.com/hyperxpizza/cdn/pkg/grpc/impl"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const buffer = 1024 * 1024

var lis *bufconn.Listener

func mockGrpcServer(c *config.Config, secure bool) error {
	lis = bufconn.Listen(buffer)
	server := grpc.NewServer()

	cdnServiceServer, err := impl.NewCDNService(c, secure)
	if err != nil {
		return err
	}

	pb.RegisterCDNGrpcServiceServer(server, cdnServiceServer)

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func TestSearchFile(t *testing.T) {
	flag.Parse()

	if *configPath == "" {
		t.Fail()
		return
	}

	c, err := config.NewConfig(*configPath)
	assert.NoError(t, err)

	go mockGrpcServer(c, *secure)

}

func TestUploadFile(t *testing.T) {}

func TestDownloadFile(t *testing.T) {}

func TestDeleteFile(t *testing.T) {}
