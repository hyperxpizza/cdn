package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/hyperxpizza/cdn/pkg/config"
	"github.com/hyperxpizza/cdn/pkg/database"
	pb "github.com/hyperxpizza/cdn/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CDNServiceImpl struct {
	pb.UnimplementedCDNGrpcServiceServer
	cfg *config.Config
	db  *database.Database
}

func NewCDNService(c *config.Config) (*CDNServiceImpl, error) {
	db, err := database.NewDatabase(c)
	if err != nil {
		return nil, err
	}

	return &CDNServiceImpl{
		cfg: c,
		db:  db,
	}, nil
}

func (cdn *CDNServiceImpl) Run() {
	grpcServer := grpc.NewServer()
	pb.RegisterCDNGrpcServiceServer(grpcServer, cdn)

	addr := fmt.Sprintf(":%d", cdn.cfg.Grpc.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("grpc server running on %s:%d", cdn.cfg.Grpc.Host, cdn.cfg.Grpc.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Println("Failed to serve:", err.Error())
		return
	}

}

func (c CDNServiceImpl) UploadFile(rstream pb.CDNGrpcService_UploadFileServer) error {
	return nil
}

func (c CDNServiceImpl) DownloadFile(req *pb.DownloadFileRequest, stream pb.CDNGrpcService_DownloadFileServer) error {
	return nil
}

func (c CDNServiceImpl) SearchFiles(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	var resp pb.SearchResponse

	files, err := c.db.SearchFiles(req.GetPhrase())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &resp, nil
		}

		return nil, status.Error(
			codes.Internal,
			err.Error(),
		)
	}

	for _, f := range files {
		mappedFile := mapFile(f)
		resp.Files = append(resp.Files, &mappedFile)
	}

	return &resp, nil
}
