package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/micro/go-micro"
	"github.com/ob-vss-ss19/sample-micro-tests/srv/sample/handler"
	sample "github.com/ob-vss-ss19/sample-micro-tests/srv/sample/proto"
)

// RunService starts the sample service
func RunService(ctx context.Context, test bool) {
	var port int64
	port = 0
	if test {
		//Calculate a random port for starting the service.
		//By default you will get some strange behaviour like "port already in use"
		reader := rand.Reader
		rsp, _ := rand.Int(reader, big.NewInt(1000))
		fmt.Printf("Using port %v\n", rsp)
		port = 1024 + 4 + rsp.Int64()
	}

	service := micro.NewService(
		micro.Name("go.micro.srv.sample"),
		micro.Address(fmt.Sprintf(":%v", port)),
		micro.Context(ctx),
	)

	//Disable the service Init method. (Does command line parsing and breaks coverage checks) -
	if !test {
		service.Init()
	}
	data := []*sample.Item{}
	data = append(data, &sample.Item{Name: "Test"})
	data = append(data, &sample.Item{Name: "Eintrag"})
	var maxID int32
	maxID = 2
	r := sample.RegisterSampleHandler(service.Server(), &handler.SampleService{Items: data, NextID: maxID})
	if r != nil {
		log.Fatalf("Registering cinema handler failed! %v\n", r.Error())
	}
	r = service.Run()
	if r != nil {
		log.Fatalf("Running service failed! %v\n", r.Error())
	}
}
