package main

import (
	"context"
	pb "github.com/rasha-hantash/grpc-gateway/rewardsrefund"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	ctx, cancelTimeoutFunc := context.WithTimeout(context.Background(), 3*time.Second)
	conn, err := grpc.DialContext(ctx, "localhost:5566",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)

	cancelTimeoutFunc()
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	log.Printf("Dialed OK ...")
	defer conn.Close()
	c := pb.NewRefundServiceClient(conn)

	// Contact the server and print out its response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddUser(ctx, &pb.AddUserRequest{})
	if err != nil {
		log.Fatalf("could not add user: %v", err)
	}
	log.Printf("User Id: " + r.GetId())

}
