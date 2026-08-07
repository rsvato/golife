package main

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	api "github.com/rsvato/golife/api"
	lib "github.com/rsvato/golife/lib"
)

type lifeServer struct {
	api.UnimplementedLifeServiceServer
}

const maxSize = 1024

func (s *lifeServer) StreamEvolution(stream api.LifeService_StreamEvolutionServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	if req.InitialState == nil {
		return status.Error(codes.InvalidArgument, "initial_state is required")
	}

	width := int(req.InitialState.Width)
	height := int(req.InitialState.Height)

	if width <= 0 || height <= 0 {
		return status.Error(codes.InvalidArgument, "width and height must be positive")
	}

	if width > maxSize || height > maxSize {
		return status.Error(codes.InvalidArgument, "dimension is too large (max 1024)")
	}

	delayMs := int(req.DelayMs)
	if delayMs <= 0 {
		log.Printf("delay set to 1000, was %d", delayMs)
		delayMs = 1000
	}

	field := lib.ReadRle(width, height, req.InitialState.GetRleString())
	generation := 0
	for {
		field = field.Step()
		generation++
		nextRle := lib.SaveRle(*field)

		update := &api.SimulationUpdate{
			Generation: int32(generation),
			CurrentState: &api.Board{
				Width:     int32(width),
				Height:    int32(height),
				RleString: nextRle,
			},
		}

		if err := stream.Send(update); err != nil {
			return err
		}

		select {
		case <-time.After(time.Duration(delayMs) * time.Millisecond):
		case <-stream.Context().Done():
			return stream.Context().Err()
		}

	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to start listener: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterLifeServiceServer(s, &lifeServer{})
	reflection.Register(s)

	log.Printf("Server started at %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
