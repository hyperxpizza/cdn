package impl

import (
	"bufio"
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"

	"github.com/hyperxpizza/cdn/pkg/compressor"
	"github.com/hyperxpizza/cdn/pkg/config"
	"github.com/hyperxpizza/cdn/pkg/customErrors"
	"github.com/hyperxpizza/cdn/pkg/database"
	"github.com/hyperxpizza/cdn/pkg/filebrowser"
	pb "github.com/hyperxpizza/cdn/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CDNServiceImpl struct {
	pb.UnimplementedCDNGrpcServiceServer
	cfg *config.Config
	db  *database.Database
	fb  *filebrowser.FileBrowser
}

func NewCDNService(c *config.Config) (*CDNServiceImpl, error) {
	db, err := database.NewDatabase(c)
	if err != nil {
		return nil, err
	}

	fb := filebrowser.NewFileBrowser(c)
	if err != nil {
		return nil, err
	}

	return &CDNServiceImpl{
		cfg: c,
		db:  db,
		fb:  fb,
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

func (c CDNServiceImpl) UploadFile(stream pb.CDNGrpcService_UploadFileServer) error {

	req, err := stream.Recv()
	if err != nil {
		return status.Error(
			codes.Internal,
			err.Error(),
		)
	}

	file := req.GetFile()
	if file.GetSize() > uint64(c.cfg.Uploader.MaxFileSize) {
		return status.Error(
			codes.FailedPrecondition,
			customErrors.ErrFileTooLarge,
		)
	}

	/*
		//check if file exists in the database
		err = c.db.CheckIfFileExists(file.Name, file.Bucket)
		if err != nil {

		}
	*/

	data := bytes.Buffer{}

	for {
		req, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return status.Error(
				codes.FailedPrecondition,
				err.Error(),
			)
		}

		chunk := req.GetChunkData()
		_, err = data.Write(chunk)
		if err != nil {
			return status.Error(
				codes.Canceled,
				err.Error(),
			)
		}
	}

	//compress file
	sizeAfterCompression, compressedData, err := compressor.CompressFile(data.Bytes())
	if err != nil {
		return status.Error(
			codes.Internal,
			err.Error(),
		)
	}

	//insert into the bucket
	err = c.fb.SaveFile(compressedData, file.Name, file.Bucket)
	if err != nil {
		return status.Error(
			codes.Internal,
			err.Error(),
		)
	}

	file.SizeAfterCompression = uint64(sizeAfterCompression)

	//insert into the database
	mappedFile := unmapFile(file)
	err = c.db.AddFile(mappedFile)
	if err != nil {
		return status.Error(
			codes.Internal,
			err.Error(),
		)
	}

	return nil
}

func (c CDNServiceImpl) DownloadFile(req *pb.DownloadFileRequest, stream pb.CDNGrpcService_DownloadFileServer) error {
	//check if file exists in the database
	err := c.db.CheckIfFileExists(req.Name, req.Bucket)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return status.Error(
				codes.NotFound,
				customErrors.ErrFileNotFound,
			)
		}
	}

	//check if file exists in the filebrowser
	compressedFile, err := c.fb.GetFile(req.Name, req.Bucket)
	if err != nil {
		if errors.Is(err, customErrors.Wrap(customErrors.ErrBucketNotFound)) || errors.Is(err, customErrors.Wrap(customErrors.ErrFileNotFound)) {
			return status.Error(
				codes.NotFound,
				err.Error(),
			)
		}

		return status.Error(
			codes.Internal,
			err.Error(),
		)
	}

	defer compressedFile.Close()

	data, err := ioutil.ReadAll(bufio.NewReader(compressedFile))
	if err != nil {
		return status.Error(
			codes.Internal,
			err.Error(),
		)
	}

	//decompress file
	decompressedData, err := compressor.DecompressFile(data)
	if err != nil {
		return status.Error(
			codes.Internal,
			err.Error(),
		)
	}

	chunk := 1024
	chunksSent := 0
	for {

		if chunksSent > len(decompressedData) {
			break
		}

		resp := pb.DownlaodFileRespose{
			ChunkData: decompressedData[:chunk],
		}

		err := stream.Send(&resp)
		if err != nil {
			return status.Error(
				codes.Internal,
				err.Error(),
			)
		}

		chunksSent = chunksSent + chunk
	}

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

func (c CDNServiceImpl) DeleteFile(ctx context.Context, req *pb.DeleteFileRequest) (*emptypb.Empty, error) {

	return &emptypb.Empty{}, nil
}
