package service

import (
	"context"
	"testing"
	"time"

	"github.com/micro/go-micro/client"
	sample "github.com/ob-vss-ss19/sample-micro-tests/srv/sample/proto"
	"github.com/stretchr/testify/assert"
)

func TestServiceStart(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	go RunService(ctx, true)
	time.Sleep(50 * time.Millisecond)
	var client client.Client
	srvClient := sample.NewSampleService("go.micro.srv.sample", client)
	var req sample.ListRequest
	rsp, err := srvClient.List(context.TODO(), &req)
	assert.Nil(t, err)
	assert.Len(t, rsp.Items, 2)
	cancel()
	time.Sleep(1 * time.Second)
}
