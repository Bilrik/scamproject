package main

import (
	"context"
	"fmt"
	"log"
	"myapp/Golang/grpc/Test/testpb"
	"net"

	"github.com/nyaruka/phonenumbers"
	"google.golang.org/grpc"
)

type server struct {
	testpb.UnimplementedTestServiceServer
}

func (*server) Valid(ctx context.Context, req *testpb.IsValidrequest) (*testpb.IsValidresponse, error) {
	fmt.Printf("IsValid function was invoked with %v\n", req)
	Number := req.GetValid().GetNumber()
	num, err := phonenumbers.Parse(Number, "US")
	if err != nil {
		log.Fatal(err)
	}

	res := &testpb.IsValidresponse{
		Result: phonenumbers.IsValidNumber(num),
	}
	return res, nil
}

/*func (*server) ValidList(stream testpb.TestService_ValidListServer) error {
	fmt.Printf("IsValidList function was invoked with \n")
	result := false
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&testpb.IsValidListresponse{
				Valid: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		Number := req.GetValid().GetNumber()
		valid := true
		var Numbers []string
		Numbers = append(Numbers, Number)

		for _, num := range Numbers {
			num, err := phonenumbers.Parse(num, "US")
			if err != nil {
				log.Fatal(err)
			}
			if !phonenumbers.IsValidNumber(num) {
				valid = false
				break
			}
		}

		return nil
	}
}*/

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	testpb.RegisterTestServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
