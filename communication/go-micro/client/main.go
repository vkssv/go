package main

import (
	"communications/go-micro/proto"
	"context"
	"fmt"
	"time"

	proto "communication/go-micro/proto"

	micro "go-micro.dev/v4"
)

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	service.Init()
	// Create new greeter client and call hello
	greeter := proto.NewGreeterClient("greeter", service.Client())
	callEvery(3*time.Second, greeter, hello)
}

func hello(t time.Time, greeter proto.GreeterClient) {
	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "V, calling at " + t.String()})
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Printf("%s\n", rsp.Greeting)
}

func callEvery(d time.Duration, greeter proto.GreeterClient, f func(time.Time, proto.GreeterClient)) {
	for x := range time.Tick(d) {
		f(x, greeter)
	}
}
