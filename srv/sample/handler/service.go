package handler

import (
	"context"

	sample "github.com/ob-vss-ss19/sample-micro-tests/srv/sample/proto"
)

// CinemaService contains different cinemas and id where next cinema should be inserted
type SampleService struct {
	Items  []*sample.Item
	NextID int32
}

// List lists all cinemas
func (s *SampleService) List(ctx context.Context, req *sample.ListRequest, rsp *sample.ListResult) error {
	rsp.Items = s.Items
	return nil
}
