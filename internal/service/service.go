package service

import pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"

type Service interface {
	Get(*pb.SimpleArtistRequest) (*pb.Artist, error)
	GetAll() ([]*pb.Artist, error)
	GetArtistByAlbum(albumID string) ([]*pb.Artist, error)
	GetArtistByTrack(trackID string) ([]*pb.Artist, error)
	CleanupAndInit() error
}
