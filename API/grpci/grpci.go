package grpci

import (
	"context"
	"fmt"
	"log"
	"scamProtec/invalid/testpb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Calls(c *gin.Context) {
	PhoneNumber := c.Param("Number")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50053", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	s := testpb.NewTestServiceClient(cc)

	req := &testpb.IsValidrequest{
		Valid: &testpb.IsValid{
			Number: PhoneNumber,
		},
	}

	res, err := s.Valid(context.Background(), req)
	if err != nil {
		fmt.Print("Error")
		log.Fatalf("error while calling Test RPC: %v", err)
	}

	fmt.Println(res)
}

func Batch(c *gin.Context) {
	PhoneNumber := c.Param("Number")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50053", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	s := testpb.NewTestServiceClient(cc)

	req := &testpb.IsValidrequest{
		Valid: &testpb.IsValid{
			Number: PhoneNumber,
		},
	}

	res, err := s.Valid(context.Background(), req)
	if err != nil {
		fmt.Print("Error")
		log.Fatalf("error while calling Test RPC: %v", err)
	}

	fmt.Println(res)
}
