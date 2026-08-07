package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/rsvato/golife/api"
	"github.com/rsvato/golife/lib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const connectionTimeout = 5 * time.Second

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Cannot dial: %v", err)
	}
	defer conn.Close()

	conn.Connect()

	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	defer cancel()

	client := api.NewLifeServiceClient(conn)
	req := &api.SimulationRequest{
		InitialState: &api.Board{
			Width:     5,
			Height:    5,
			RleString: "7.1#4.1#4.1#7.",
		},
		DelayMs: 500,
	}

	stream, err := client.StreamEvolution(ctx)
	if err != nil {
		log.Fatalf("Couldn't open stream: %v", err)
	}

	if err := stream.Send(req); err != nil {
		log.Fatalf("Couldn't send initial state: %v", err)
	}

	fmt.Printf("Simulation started.")

	for {
		update, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("Stream closed by server")
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive: %v", err)
		}

		clearScreen()
		fmt.Printf("Generation %d\n", update.Generation)
		width := int(update.CurrentState.Width)
		height := int(update.CurrentState.Height)
		field := lib.ReadRle(width, height, update.CurrentState.GetRleString())

		renderField(field)
	}
}

func renderField(f *lib.Field) {
	fmt.Print(f.String())
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
