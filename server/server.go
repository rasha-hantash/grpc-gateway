package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/rasha-hantash/grpc-gateway/rewardsrefund"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
)

// Backend implements the protobuf interface
type Backend struct {
	pb.UnimplementedRefundServiceServer
	mu    *sync.RWMutex
	users []*pb.User
}

// New initializes a new Backend struct
func New() *Backend {
	return &Backend{
		mu: &sync.RWMutex{},
	}
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "www/swagger.json")
}

func startHTTP() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Connect to the GRPC server
	conn, err := grpc.Dial(":5566", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// Register grpc-gateway
	rmux := runtime.NewServeMux()
	client := pb.NewRefundServiceClient(conn)
	err = pb.RegisterRefundServiceHandlerClient(ctx, rmux, client)
	if err != nil {
		log.Fatal(err)
	}

	// Serve the swagger,
	mux := http.NewServeMux()
	mux.Handle("/", rmux)
	mux.HandleFunc("/swagger.json", serveSwagger)
	fs := http.FileServer(http.Dir("www/swagger-ui"))
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", fs))

	log.Println("REST server ready...")
	log.Println("Serving Swagger at: http://localhost:8080/swagger-ui/")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func startGRPC() {
	// TODO to run locally without container change code to
	// localhost:5566
	lis, err := net.Listen("tcp", ":5566")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRefundServiceServer(grpcServer, &Backend{})
	log.Println("gRPC server ready...")
	grpcServer.Serve(lis)
}

func main() {
	go startGRPC()
	go startHTTP()
	go New()

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
