package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "grpc-routes/routes"

	"github.com/libp2p/go-netroute"
	"google.golang.org/grpc"
)

// server is used to implement GetRoutesServer
type server struct {
	pb.UnimplementedGetRoutesServer
}

var (
	port = flag.Int("port", 8088, "GRPC Server port")
)

// SendRoutes implements grpc-routes.
func (s *server) SendRoutes(ctx context.Context, in *pb.RoutesRequest) (*pb.RoutesReply, error) {
	log.Printf("Received route request to: %v", in.GetDestination())

	routeResponse, err := determineRoute(in.GetDestination())
	if err != nil {
		return nil, err
	}

	return &pb.RoutesReply{Routetable: routeResponse}, nil
}

// deter
func determineRoute(destAddr string) (string, error) {
	r, err := netroute.New()
	if err != nil {
		panic(err)
	}
	iface, gw, sc, err := r.Route(net.ParseIP(destAddr))
	return ("Interface: " + iface.Name + " Gateway: " + gw.String() + " Source: " + sc.String()), nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetRoutesServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
