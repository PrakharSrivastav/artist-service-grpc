package client

import (
	"context"
	pb "github.com/PrakharSrivastav/gql-grpc-defintions/go/schema"
)

// GetAlbum gets the album details by Id
func (c *Client) GetAlbum(ctx context.Context, albumID string) (*pb.Album, error) {
	return c.albumClientRPC.Get(ctx, &pb.SimpleAlbumRequest{Id: albumID})
}

// GetTrack gets the track details by Id
func (c *Client) GetTrack(ctx context.Context, trackID string) (*pb.Track, error) {
	return c.trackClientRPC.Get(ctx, &pb.SimpleTrackRequest{Id: trackID})
}
