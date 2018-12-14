package service

import (
	"context"
	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
)

type Service interface {
	Get(*pb.SimpleArtistRequest) (*pb.Artist, error)
	GetAll() ([]*pb.Artist, error)
	GetArtistByAlbum(ctx context.Context, albumID string) ([]*pb.Artist, error)
	GetArtistByTrack(ctx context.Context, trackID string) ([]*pb.Artist, error)
	CleanupAndInit() error
}
