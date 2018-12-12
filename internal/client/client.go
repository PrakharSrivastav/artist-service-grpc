package client

import (
	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
	"github.com/opentracing/opentracing-go"
)
import "google.golang.org/grpc"

type Client struct {
	albumClientRPC pb.AlbumServiceClient
	trackClientRPC pb.TrackServiceClient
}

func NewClient(tracer opentracing.Tracer) *Client {
	c := Client{}
	connArtist := getConnection("album-service-client", "localhost:6564" ,tracer)//"artist-service:6565")
	connTracks := getConnection("track-service-client", "localhost:6560", tracer)//"track-service:6565")
	c.albumClientRPC = pb.NewAlbumServiceClient(connArtist)
	c.trackClientRPC = pb.NewTrackServiceClient(connTracks)
	return &c
}

func getConnection(name string, addr string , tracer opentracing.Tracer) *grpc.ClientConn {
	connection := New(name, addr , tracer)
	return connection.Dial()
}
