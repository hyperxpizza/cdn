package impl

import (
	"github.com/hyperxpizza/cdn/pkg/filebrowser"
	pb "github.com/hyperxpizza/cdn/pkg/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapFile(file *filebrowser.File) pb.File {
	return pb.File{
		Id:                   int64(file.ID),
		Name:                 file.Name,
		Bucket:               file.Bucket,
		Size:                 file.Size,
		SizeAfterCompression: file.SizeAfterCompression,
		Extension:            file.Extension,
		MimeType:             file.MimeType,
		Created:              timestamppb.New(file.Created),
		Updated:              timestamppb.New(file.Updated),
	}

}
