package service

import "context"
import pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"

type Service interface {
	Get(*pb.SimpleArtistRequest) (*pb.Artist, error)
	GetAll() ([]*pb.Artist, error)
	GetArtistByAlbum(ctx context.Context, req *pb.SimpleArtistRequest) ([]*pb.Artist, error)
	GetArtistByTrack(ctx context.Context, req *pb.SimpleArtistRequest) ([]*pb.Artist, error)
	CleanupAndInit() error
}
