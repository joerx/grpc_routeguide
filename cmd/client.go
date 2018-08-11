package cmd

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/joerx/grpc_routeguide/routeguide"
	"google.golang.org/grpc"
)

var clientAddr string

func init() {
	flag.StringVar(&clientAddr, "server", "localhost:10000", "Server address for client")
	Commands["client"] = Client
}

// Client cmd runs a simple demo client
func Client() error {
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)

	pt := &pb.Point{Latitude: 407838351, Longitude: -746143763}
	hi := &pb.Point{Latitude: 407838351, Longitude: -746143763}
	lo := &pb.Point{Latitude: 407838351, Longitude: -746143763}
	rect := &pb.Rectangle{Hi: hi, Lo: lo}

	fmt.Printf("GetFeature(%v)\n", pt)
	if err := getFeature(client, pt); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println()

	fmt.Printf("ListFeatures(%v)\n", rect)
	if err := listFeatures(client, rect); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return nil
}

func listFeatures(client pb.RouteGuideClient, rect *pb.Rectangle) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := client.ListFeatures(ctx, rect)
	if err != nil {
		return fmt.Errorf(`rpc call error "%v"`, grpc.ErrorDesc(err))
	}

	for {
		ft, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf(`error in stream "%s"`, grpc.ErrorDesc(err))
		}
		log.Println(ft)
	}

	return nil
}

func getFeature(client pb.RouteGuideClient, point *pb.Point) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ft, err := client.GetFeature(ctx, point)
	if err != nil {
		return fmt.Errorf(`error getting feature "%s"`, grpc.ErrorDesc(err))
	}
	fmt.Println(ft)

	return nil
}
