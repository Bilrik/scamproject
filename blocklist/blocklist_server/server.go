package blocklist

import (
	"blocklistProtec/blocklsit/blocklistpb"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"scamProtec/blocklist/blocklistpb"
	"scamProtec/scam/scampb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	blocklistpb.BlocklistServiceServer
}

func (*server) ScoreEvent(ctx context.Context, req *blocklistpb.ScoreEventRequest) (*blocklistpb.ScoreEventResponse, error) {
	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close() // Maybe this should be in a separate function and the error handled?

	c := scampb.NewScamDatabaseClient(cc)

	getBlocked, err := c.GetPersonalBlocklist(ctx, &scampb.GetPersonalBlocklistRequest{PhoneNumber: req.GetBNumber})
	if err != nil {
		log.Printf("Could not get blocklist: %v\n", err)
		return nil, err
	}

	score, err := c.GetScore(ctx, &scampb.GetScoreRequest{PhoneNumber: req.GetANumber()})
	if err != nil {
		log.Printf("Unable to get Score: %v\n", err)
		return &blocklistpb.ScoreEventResponse{
			EventId:   req.EventId,
			BlockCall: ,
			IsScam:    false,
			Score:     score.Score,
			Metadata:  &blocklistpb.Metadata{Category: category},
		}, err
	}

	getCategoryResponse, err := c.GetPhoneCategory(ctx, &scampb.GetPhoneCategoryRequest{PhoneNumber: req.GetANumber()})
	if err != nil {
		log.Printf("Could not get phone category: %v\n", err)
		return nil, err
	}

	category := getCategoryResponse.GetCategoryCode()

	if score.GetScore() >= 75 {
		return &blocklistpb.ScoreEventResponse{
			EventId:   req.EventId,
			BlockCall: true,
			IsScam:    true,
			Score:     score.Score,
			Metadata:  &blocklistpb.Metadata{Category: category},
		}, nil
	}

	blockedCategories := getBlocked.GetBlockedCategories()

	for _, cat := range blockedCategories {
		if cat == category {
			return &blocklistpb.ScoreEventResponse{
				EventId:   req.EventId,
				BlockCall: true,
				IsScam:    false,
				Score:     score.Score,
				Metadata:  &blocklistpb.Metadata{Category: category},
			}, nil
		}
	}

	blockedANumbers := getBlocked.GetBlockedANumbers()

	for _, num := range blockedANumbers {
		if num == req.GetANumber() {
			return &blocklistpb.ScoreEventResponse{
				EventId:   req.EventId,
				BlockCall: true,
				IsScam:    false,
				Score:     score.Score,
				Metadata:  &blocklistpb.Metadata{Category: category},
			}, nil
		}
	}

	return &blocklistpb.ScoreEventResponse{
		EventId:   req.EventId,
		BlockCall: false,
		IsScam:    false,
		Score:     score.Score,
		Metadata:  &blocklistpb.Metadata{Category: category},
	}, nil
}

func (*server) ScoreEvents(stream blocklistpb.BlocklistService_ScoreEventsServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}

		res, err := (*server).ScoreEvent(&server{}, context.Background(), req)
		if err != nil {
			log.Println(err)
		}

		stream.send(res)
	}
	return nil
}

func main() {
	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Blocklist Service Started")

	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	blocklistpb.RegisterBlocklistServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch

	// Finally, we stop the server
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of Program")
}
