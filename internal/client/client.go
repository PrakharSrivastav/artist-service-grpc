package client

import pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
import "google.golang.org/grpc"

type Client struct {
	albumClientRPC pb.AlbumServiceClient
	trackClientRPC pb.TrackServiceClient
}

func NewClient() *Client {
	c := Client{}
	connArtist := getConnection("album-service-client", "localhost:6566")
	connTracks := getConnection("track-service-client", "localhost:6567")
	c.albumClientRPC = pb.NewAlbumServiceClient(connArtist)
	c.trackClientRPC = pb.NewTrackServiceClient(connTracks)
	return &c
}

func getConnection(name string, addr string) *grpc.ClientConn {
	connection := New(name, addr)
	return connection.Dial()
}
