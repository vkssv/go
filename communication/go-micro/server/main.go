package main

import (
	proto "communication/go-micro/proto"
	"context"
	"fmt"

	micro "go-micro.dev/v4"
)

// The Greeter API
type Greeter struct{}

// Hello is a Greeter API method
func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello" + req.Name
	fmt.Printf("Responding with %s\n", rsp.Greeting)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("1.0.1"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
