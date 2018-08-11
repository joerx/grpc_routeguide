package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	pb "github.com/joerx/grpc_routeguide/routeguide"
)

type routeGuideServer struct {
	features []*pb.Feature
}

// New creates a new RouteGuideServer instance, loading data from given path
func New(dbPath string) (pb.RouteGuideServer, error) {
	s := &routeGuideServer{}
	if err := s.loadFeatures(dbPath); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *routeGuideServer) loadFeatures(path string) error {
	log.Printf("Loading features from %s", path)

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, &s.features); err != nil {
		return err
	}

	log.Printf("Loaded %d features", len(s.features))
	return nil
}

func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	log.Printf("Incoming call to GetFeature")
	for _, f := range s.features {
		if proto.Equal(f.Location, point) {
			return f, nil
		}
	}
	return &pb.Feature{Name: "", Location: point}, nil
}

func (s *routeGuideServer) ListFeatures(rect *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
	log.Printf("Incoming call to ListFeatures")
	return fmt.Errorf("Not implemented")
}

func (s *routeGuideServer) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {
	log.Printf("Incoming call to RecordRoute")
	return fmt.Errorf("Not implemented")
}

func (s *routeGuideServer) RouteChat(stream pb.RouteGuide_RouteChatServer) error {
	log.Printf("Incoming call to RouteChat")
	return fmt.Errorf("Not implemented")
}
