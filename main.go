package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"

	pb "github.com/mytaowu/proto"
)

var target = "127.0.0.1:9999"

func main() {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer conn.Close()

	c := pb.NewHelloGRPCClient(conn)
	for i := 0; i < 10; i++ {
		rsp, err := c.SayHi(context.Background(), &pb.Req{Message: "mytaowu " + string(i)})
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}

		fmt.Printf("rsp: %v\n", rsp)
		time.Sleep(time.Second * 3)
	}

}
