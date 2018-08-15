package service

import (
	"context"

	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/jmoiron/sqlx"
)

type serviceImpl struct {
	repo *repository
}

func New(db *sqlx.DB) Service {
	return &serviceImpl{
		repo: &repository{db: db},
	}
}

func (f *serviceImpl) Get(req *pb.SimpleArtistRequest) (*pb.Artist, error) {
	artist, err := f.repo.get(req.GetId())
	if err != nil {
		return nil, err
	}
	return artist.ToProto(), err
}

func (f *serviceImpl) GetAll() ([]*pb.Artist, error) {
	artists, err := f.repo.getAll()
	if err != nil {
		return nil, err
	}

	var protoArtists []*pb.Artist
	for _, a := range artists {
		protoArtists = append(protoArtists, a.ToProto())
	}

	return protoArtists, nil
}

func (f *serviceImpl) GetArtistByAlbum(ctx context.Context, req *pb.SimpleArtistRequest) ([]*pb.Artist, error) {
	panic("not implemented")
}

func (f *serviceImpl) GetArtistByTrack(ctx context.Context, req *pb.SimpleArtistRequest) ([]*pb.Artist, error) {
	panic("not implemented")
}

func (f *serviceImpl) CleanupAndInit() error {
	return f.repo.setupDatabase()
}
