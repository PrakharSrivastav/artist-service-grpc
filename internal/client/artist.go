package client

import (
	"context"
	"fmt"

	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
)

// GetAlbum gets the album details by Id
func (c *Client) GetAlbum(albumID string) (*pb.Album, error) {
	return c.albumClientRPC.Get(context.Background(), &pb.SimpleAlbumRequest{Id: albumID})
}

// GetTrack gets the track details by Id
func (c *Client) GetTrack(trackID string) (*pb.Track, error) {
	fmt.Println("Client code")
	return c.trackClientRPC.Get(context.Background(), &pb.SimpleTrackRequest{Id: trackID})
}
