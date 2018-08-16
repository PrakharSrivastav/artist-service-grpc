package rpc

import (
	"context"
	"fmt"

	"github.com/PrakharSrivastav/artist-service-grpc/internal/service"
	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

// ArtistService implements the grpc interfaces
type ArtistService struct {
	service service.Service
}

// Get fetches an artist by id
func (f *ArtistService) Get(_ context.Context, req *pb.SimpleArtistRequest) (*pb.Artist, error) {
	return f.service.Get(req)
}

// GetAll fetches all the artists from database
func (f *ArtistService) GetAll(_ *empty.Empty, stream pb.ArtistService_GetAllServer) error {
	artists, err := f.service.GetAll()
	if err != nil {
		fmt.Println("Error ::", err.Error())
		return err
	}
	for _, a := range artists {
		fmt.Println("Iterating")
		if err := stream.Send(a); err != nil {
			fmt.Println("Error processing stream :: ", err.Error())
			return err
		}
	}
	return nil
}

// GetArtistByAlbum fetches all the artists for an album
func (f *ArtistService) GetArtistByAlbum(req *pb.SimpleArtistRequest, stream pb.ArtistService_GetArtistByAlbumServer) error {
	artists, err := f.service.GetArtistByAlbum(req.GetId())
	if err != nil {
		return err
	}
	for _, a := range artists {
		if err := stream.Send(a); err != nil {
			fmt.Println("Error processing stream :: ", err.Error())
			return err
		}
	}
	return nil
}

// GetArtistByTrack fetches all the artists for a track
func (f *ArtistService) GetArtistByTrack(req *pb.SimpleArtistRequest, stream pb.ArtistService_GetArtistByTrackServer) error {
	artists, err := f.service.GetArtistByTrack(req.GetId())
	if err != nil {
		return err
	}
	for _, a := range artists {
		if err := stream.Send(a); err != nil {
			fmt.Println("Error processing stream :: ", err.Error())
			return err
		}
	}
	return nil
}

// Register registers artist implementation in grpc
func (f *ArtistService) Register(server *grpc.Server) {
	pb.RegisterArtistServiceServer(server, f)
}

// New service return new instance of ArtistService
func New(service service.Service) *ArtistService {
	return &ArtistService{
		service: service,
	}
}
