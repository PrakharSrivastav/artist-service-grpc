package service

import (
	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"

	"github.com/PrakharSrivastav/artist-service-grpc/internal/client"
	"github.com/jmoiron/sqlx"
)

type serviceImpl struct {
	repo   *repository
	client *client.Client
}

func New(db *sqlx.DB) Service {
	return &serviceImpl{
		repo:   &repository{db: db},
		client: client.NewClient(),
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

func (f *serviceImpl) GetArtistByAlbum(albumID string) ([]*pb.Artist, error) {

	// get all the artists in the album (grpc)
	var artistIds []string
	if album, err := f.client.GetAlbum(albumID); err != nil {
		return nil, err
	} else {
		artistIds = append(artistIds, album.GetArtistId())
	}
	// query for these ids in database
	artists, err := f.repo.getIn(artistIds)
	if err != nil {
		return nil, err
	}

	var protoArtists []*pb.Artist
	for _, a := range artists {
		protoArtists = append(protoArtists, a.ToProto())
	}

	return protoArtists, nil
}

func (f *serviceImpl) GetArtistByTrack(trackId string) ([]*pb.Artist, error) {
	// Get the artists for a track
	var ids []string
	if track, err := f.client.GetTrack(trackId); err != nil {
		return nil, err
	} else {
		ids = append(ids, track.GetArtistId())
	}

	// query for these ids in database
	artists, err := f.repo.getIn(ids)
	if err != nil {
		return nil, err
	}

	var protoArtists []*pb.Artist
	for _, a := range artists {
		protoArtists = append(protoArtists, a.ToProto())
	}

	return protoArtists, nil
}

func (f *serviceImpl) CleanupAndInit() error {
	return f.repo.setupDatabase()
}
