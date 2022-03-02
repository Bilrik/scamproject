package main

import (
	"context"
	"fmt"
	"log"
	"myapp/Golang/grpc/Test/testpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := testpb.NewTestServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	doUnary(c)
}

func doUnary(c testpb.TestServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &testpb.IsValidrequest{
		Valid: &testpb.IsValid{
			Number: "5013398428",
		},
	}
	res, err := c.Valid(context.Background(), req)
	if err != nil {
		fmt.Print("Error")
		log.Fatalf("error while calling Test RPC: %v", err)
	}
	log.Printf("Phone Number Valid testing (%s): %v", req.Valid.Number, res.Result)
}

/*func doClientStreaming(c testpb.TestServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*testpb.IsValidListrequest{
		&testpb.IsValidListrequest{
			Valid: &testpb.IsValidrequest{
				Number: "5013398428",
			},
		},
		&testpb.IsValidListrequest{
			Valid: &testpb.IsValid{
				Number: "50188880007",
			},
		},
		&testpb.IsValidListrequest{
			Valid: &testpb.IsValidrequest{
				Number: "0013398428",
			},
		},
		&testpb.IsValidListrequest{
			Valid: &testpb.IsValidrequest{
				Number: "555",
			},
		},
	}

	stream, err := c.ValidList(context.Background())
	if err != nil {
		log.Fatalf("error while calling ValidList: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from ValidList: %v", err)
	}
	fmt.Printf("ValidList Response: %v\n", res)
}*/
