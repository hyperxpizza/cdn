package impl

import (
	"context"

	"github.com/hyperxpizza/cdn/pkg/config"
	pb "github.com/hyperxpizza/cdn/pkg/grpc"
)

type CDNServiceImpl struct {
	pb.UnimplementedCDNGrpcServiceServer
	cfg *config.Config
}

func NewCDNService(c *config.Config) CDNServiceImpl {
	return CDNServiceImpl{
		cfg: c,
	}
}

func (c *CDNServiceImpl) UploadFile(req *pb.UploadFileRequest, stream pb.CDNGrpcService_UploadFileServer) error {
	return nil
}

func (c *CDNServiceImpl) DownloadFile(req *pb.DownloadFileRequest, stream pb.CDNGrpcService_DownloadFileServer) error {
	return nil
}

func (c *CDNServiceImpl) SearchFiles(ctx context.Context, req *pb.SearchResponse) (*pb.SearchResponse, error) {
	var resp pb.SearchResponse
	return &resp, nil
}
