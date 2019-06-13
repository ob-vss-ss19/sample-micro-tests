package handler

import (
	"context"
	"testing"

	sample "github.com/ob-vss-ss19/sample-micro-tests/srv/sample/proto"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	service := &SampleService{}
	rsp := sample.ListResult{}
	ctx := context.TODO()
	err := service.List(ctx, nil, &rsp)
	assert.Nil(t, err)
	assert.Equal(t, len(rsp.Items), 0)
}
