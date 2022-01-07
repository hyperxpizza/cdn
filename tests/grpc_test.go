package tests

import (
	"bufio"
	"context"
	"flag"
	"io"
	"log"
	"net"
	"os"
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

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

/*
func TestSearchFile(t *testing.T) {
	flag.Parse()

	if *configPath == "" {
		t.Fail()
		return
	}

	c, err := config.NewConfig(*configPath)
	assert.NoError(t, err)

	go mockGrpcServer(c, *secure)

	ctx := context.Background()
	connection, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	assert.NoError(t, err)

	defer connection.Close()

	client := pb.NewCDNGrpcServiceClient(connection)

	req := pb.SearchRequest{
		Phrase: "some-search-phrase",
	}
	client.SearchFiles(ctx)
}
*/

/*
go test -v ./tests/ --run TestUploadFile --config=/home/hyperxpizza/dev/golang/cdn/config.json --secure=false --filePath=/home/hyperxpizza/dev/grafana_test.json --bucket=test-bucket --delete=true
*/
func TestUploadFile(t *testing.T) {
	flag.Parse()

	if *configPath == "" {
		t.Fail()
		return
	}

	if *filePath == "" {
		t.Fail()
		return
	}

	if *bucket == "" {
		t.Fail()
		return
	}

	printFlags()

	c, err := config.NewConfig(*configPath)
	assert.NoError(t, err)

	go mockGrpcServer(c, *secure)

	ctx := context.Background()
	connection, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	assert.NoError(t, err)

	defer connection.Close()

	client := pb.NewCDNGrpcServiceClient(connection)

	file, err := os.Open(*filePath)
	assert.NoError(t, err)

	defer file.Close()

	info, err := os.Stat(*filePath)
	assert.NoError(t, err)

	stream, err := client.UploadFile(ctx)
	assert.NoError(t, err)

	request := &pb.UploadFileRequest{
		Data: &pb.UploadFileRequest_File{
			File: &pb.File{
				Name: file.Name(),
				//Bucket: *bucket,
				Size: uint64(info.Size()),
			},
		},
	}

	err = stream.Send(request)
	assert.NoError(t, err)

	reader := bufio.NewReader(file)
	buf := make([]byte, 1024)

	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fail()
			return
		}

		request := &pb.UploadFileRequest{
			Data: &pb.UploadFileRequest_ChunkData{
				ChunkData: buf[:n],
			},
		}

		err = stream.Send(request)
		assert.NoError(t, err)
	}

	_, err = stream.CloseAndRecv()
	assert.NoError(t, err)

	if *delete {

	}
}

func TestDownloadFile(t *testing.T) {}

func TestDeleteFile(t *testing.T) {}
