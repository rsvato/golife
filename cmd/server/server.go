package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	api "github.com/rsvato/golife/api"
	lib "github.com/rsvato/golife/lib"
)

type lifeServer struct {
	api.UnimplementedLifeServiceServer
}

func (s *lifeServer) StreamEvolution(stream api.LifeService_StreamEvolutionServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	width := int(req.InitialState.Width)
	height := int(req.InitialState.Width)
	field := lib.ReadRle(width, height, req.InitialState.GetRleString())
	generation := 0
	for {
		field = field.Step()
		generation++
		nextRle := lib.SaveRle(*field)
		fmt.Print(field.String())

		update := &api.SimulationUpdate{
			Generation: int32(generation),
			CurrentState: &api.Board{
				Width:      int32(width),
				Height:     int32(height),
				DataFormat: &api.Board_RleString{RleString: nextRle},
			},
		}

		if err := stream.Send(update); err != nil {
			return err
		}

		select {
		case <-time.After(500 * time.Millisecond):
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
