package main

import (
	"context"
	"github.com/ob-vss-ss19/sample-micro-tests/srv/sample/service"
)

func main() {
	service.RunService(context.Background(), false)
}
