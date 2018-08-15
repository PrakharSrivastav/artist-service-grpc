package rpc

import (
	"context"
	"fmt"

	"github.com/PrakharSrivastav/artist-service-grpc/internal/model"
	"github.com/PrakharSrivastav/artist-service-grpc/internal/service"
	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

// ArtistService implements the grpc interfaces
type ArtistService struct {
	service service.Service
}

func (f *ArtistService) Get(_ context.Context, req *pb.SimpleArtistRequest) (*pb.Artist, error) {
	return f.service.Get(req)
}

// GetAll fetches all the artists from database
func (f *ArtistService) GetAll(_ *empty.Empty, stream pb.ArtistService_GetAllServer) error {
	fmt.Println("Inside the function")
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

func (f *ArtistService) GetArtistByAlbum(req *pb.SimpleArtistRequest, stream pb.ArtistService_GetArtistByAlbumServer) error {
	fmt.Println("Inside the function")
	var artists []*model.Artist
	artists = append(artists, model.NewArtist())
	artists = append(artists, model.NewArtist())
	artists = append(artists, model.NewArtist())
	artists = append(artists, model.NewArtist())
	artists = append(artists, model.NewArtist())
	artists = append(artists, model.NewArtist())
	fmt.Println("Added to the list")
	for _, a := range artists {
		fmt.Println("Iterating")
		if err := stream.Send(a.ToProto()); err != nil {
			fmt.Println("Error processing stream :: ", err.Error())
			return err
		}
	}
	return nil
}

func (f *ArtistService) GetArtistByTrack(req *pb.SimpleArtistRequest, stream pb.ArtistService_GetArtistByTrackServer) error {
	fmt.Println("Inside the function")
	var artists []*model.Artist
	artists = append(artists, model.NewArtist())
	artists = append(artists, model.NewArtist())
	artists = append(artists, model.NewArtist())
	artists = append(artists, model.NewArtist())
	artists = append(artists, model.NewArtist())
	artists = append(artists, model.NewArtist())
	fmt.Println("Added to the list")
	for _, a := range artists {
		fmt.Println("Iterating")
		if err := stream.Send(a.ToProto()); err != nil {
			fmt.Println("Error processing stream :: ", err.Error())
			return err
		}
	}
	return nil
}

func (f *ArtistService) Register(server *grpc.Server) {
	pb.RegisterArtistServiceServer(server, f)
}

func New(service service.Service) *ArtistService {
	return &ArtistService{
		service: service,
	}
}
