// Package main implements a server for Greeter service.
package rpc

import (
	"context"
	"main/rpc/definitions"
)

// GetFeature returns the feature at the given point.
func (s *RPCServer) GetFeature(ctx context.Context, point *definitions.Point) (*definitions.Feature, error) {
	resPoint := &definitions.Point{
		Latitude: 12, Longitude: 34,
	}

	return &definitions.Feature{Location: resPoint}, nil
}
