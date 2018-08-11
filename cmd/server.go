package cmd

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/joerx/grpc_routeguide/server"
	"google.golang.org/grpc"

	pb "github.com/joerx/grpc_routeguide/routeguide"
)

var db string
var listen string

func init() {
	flag.StringVar(&db, "db", "./testdata/route_guide_db.json", "Path to data file for server")
	flag.StringVar(&listen, "listen", ":10000", "Address to listen on for server")
	Commands["server"] = Server
}

func newGRPCServer() (*grpc.Server, error) {
	srv, err := server.New(db)
	if err != nil {
		return nil, err
	}
	grpcSrv := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcSrv, srv)
	return grpcSrv, nil
}

func runServer(srv *grpc.Server) (<-chan bool, error) {
	sigs := make(chan os.Signal)
	done := make(chan bool)

	l, err := net.Listen("tcp", listen)
	if err != nil {
		return nil, err
	}

	// wait for signal, gracefully stop server if we receive one
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		log.Println("Received signal, terminating")
		srv.Stop()
		done <- true
	}()

	// start the server, print a warm and fuzzy welcome message
	go func() {
		srv.Serve(l)
	}()
	log.Printf("gRPC server is listing on %s, Ctrl-C to stop", listen)

	return done, nil
}

// Server command starts the gRPC server and waits for SIGINT or SIGINT
func Server() error {
	srv, err := newGRPCServer()
	if err != nil {
		return err
	}
	done, err := runServer(srv)
	if err != nil {
		return err
	}
	<-done
	return nil
}
