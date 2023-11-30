package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "grpc-routes/routes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultDestination = "8.8.8.8"
)

var (
	addr        = flag.String("addr", "localhost:8088", "the address to connect to")
	destination = flag.String("destination", defaultDestination, "IP destination to get Route")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGetRoutesClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendRoutes(ctx, &pb.RoutesRequest{Destination: *destination})
	if err != nil {
		log.Fatalf("could not get routes: %v", err)
	}
	log.Printf("Sending route information request for route: " + *destination)
	log.Printf("Response: %s", r.GetRoutetable())
}
