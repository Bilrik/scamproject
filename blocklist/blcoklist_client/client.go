package main

import (
	"context"
	"fmt"
	"log"
	"scamProtec/blocklist/blocklistpb"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50052", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blocklistpb.NewBlocklistServiceClient(cc)

	req := &blocklistpb.ScoreEventRequest{
		EventId:   "001",
		ANumber:   "479",
		BNumber:   "501339",
		CallTime:  time.Now().UnixNano(),
		SipInvite: "whocares",
	}

	res, err := c.ScoreEvent(context.Background(), req)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(res)
}
