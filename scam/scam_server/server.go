package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"scamProtec/scam/scampb"

	"gopkg.in/mgo.v2/bson"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"
)

var collection *mongo.Collection

type server struct {
	scampb.ScamDatabaseServer
}

type PhoneNumbers struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Number            string             `bson:"_number,omitempty"`
	CategoryCode      string             `bson:"_category_code,omitempty"`
	CategoryName      string             `bson:"_category_name,omitempty"`
	BlockedANumbers   []string           `bson:"_blocked_numbers,omitempty"`
	BlockedCategories []string           `bson:"_blocked_categories,omitempty"`
	Score             int32              `bson:"_score,omitempty"`
}

func (*server) AddPhoneNumber(ctx context.Context, req *scampb.AddPhoneNumberRequest) (*scampb.AddPhoneNumberResponse, error) {
	fmt.Println("Add PhoneNumber request")

	data := PhoneNumbers{
		Number:            req.GetNumber(),
		CategoryCode:      req.GetCategory_Code(),
		CategoryName:      req.GetCategory_Name(),
		BlockedANumbers:   req.GetBlockedANumbers(),
		BlockedCategories: req.GetBlockedCategories(),
		Score:             req.GetScore(),
	}

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return &scampb.AddPhoneNumberResponse{
				Success: false},
			status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error: %v", err),
			)
	}

	return &scampb.AddPhoneNumberResponse{
		Success: true,
	}, nil

}

func (*server) getPhoneCategory(ctx context.Context, req *scampb.GetPhoneCategoryRequest) (*scampb.GetPhoneCategoryResponse, error) {
	fmt.Println("category request")

	PhoneNumber := req.GetPhoneNumber()

	// create an empty struct
	data := &PhoneNumbers{}
	filter := bson.M{"_Number": PhoneNumber}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find Number: %v", err),
		)
	}

	return &scampb.GetPhoneCategoryResponse{
		CategoryName: data.CategoryName,
		CategoryCode: data.CategoryCode,
	}, nil
}

func (*server) GetScore(ctx context.Context, req *scampb.GetScoreRequest) (*scampb.GetScoreResponse, error) {
	fmt.Println("Score Request")

	PhoneNumber := req.GetPhoneNumber()

	// create an empty struct
	data := &PhoneNumbers{}
	filter := bson.M{"_Number": PhoneNumber}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find Number: %v", err),
		)
	}

	return &scampb.GetScoreResponse{
		Score: data.Score,
	}, nil
}

// func dataToBlogpb(data *blogItem) *scampb.Blog {
// 	return &scampb.Blog{
// 		Id:       data.ID.Hex(),
// 		AuthorId: data.AuthorID,
// 		Content:  data.Content,
// 		Title:    data.Title,
// 	}
// }

func (*server) GetPersonalBlocklist(ctx context.Context, req *scampb.GetPersonalBlocklistRequest) (*scampb.GetPersonalBlocklistResponse, error) {
	fmt.Println("blocklist request")

	PhoneNumber := req.GetPhoneNumber()

	// create an empty struct
	data := &PhoneNumbers{}
	filter := bson.M{"_number": PhoneNumber}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find number: %v", err),
		)
	}

	return &scampb.GetPersonalBlocklistResponse{
		BlockedANumbers:   data.BlockedANumbers,
		BlockedCategories: data.BlockedCategories,
	}, nil
}

// func (*server) UpdateBlog(ctx context.Context, req *scampb.UpdateBlogRequest) (*scampb.UpdateBlogResponse, error) {
// 	fmt.Println("Update blog request")
// 	blog := req.GetBlog()
// 	oid, err := primitive.ObjectIDFromHex(blog.GetId())
// 	if err != nil {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			fmt.Sprintf("Cannot parse ID"),
// 		)
// 	}

// 	// create an empty struct
// 	data := &blogItem{}
// 	filter := bson.M{"_number": N}

// 	res := collection.FindOne(ctx, filter)
// 	if err := res.Decode(data); err != nil {
// 		return nil, status.Errorf(
// 			codes.NotFound,
// 			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
// 		)
// 	}

// 	// we update our internal struct
// 	data.AuthorID = blog.GetAuthorId()
// 	data.Content = blog.GetContent()
// 	data.Title = blog.GetTitle()

// 	_, updateErr := collection.ReplaceOne(context.Background(), filter, data)
// 	if updateErr != nil {
// 		return nil, status.Errorf(
// 			codes.Internal,
// 			fmt.Sprintf("Cannot update object in MongoDB: %v", updateErr),
// 		)
// 	}

// 	return &scampb.UpdateBlogResponse{
// 		Blog: dataToscampb(data),
// 	}, nil

// }

// func (*server) DeleteBlog(ctx context.Context, req *scampb.DeleteBlogRequest) (*scampb.DeleteBlogResponse, error) {
// 	fmt.Println("Delete blog request")
// 	oid, err := primitive.ObjectIDFromHex(req.GetBlogId())
// 	if err != nil {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			fmt.Sprintf("Cannot parse ID"),
// 		)
// 	}

// 	filter := bson.M{"_id": oid}

// 	res, err := collection.DeleteOne(ctx, filter)

// 	if err != nil {
// 		return nil, status.Errorf(
// 			codes.Internal,
// 			fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
// 		)
// 	}

// 	if res.DeletedCount == 0 {
// 		return nil, status.Errorf(
// 			codes.NotFound,
// 			fmt.Sprintf("Cannot find blog in MongoDB: %v", err),
// 		)
// 	}

// 	return &scampb.DeleteBlogResponse{BlogId: req.GetBlogId()}, nil
// }

// func (*server) ListBlog(_ *scampb.ListBlogRequest, stream scampb.BlogService_ListBlogServer) error {
// 	fmt.Println("List blog request")

// 	cur, err := collection.Find(context.Background(), primitive.D{{}})
// 	if err != nil {
// 		return status.Errorf(
// 			codes.Internal,
// 			fmt.Sprintf("Unknown internal error: %v", err),
// 		)
// 	}
// 	defer cur.Close(context.Background()) // Should handle err
// 	for cur.Next(context.Background()) {
// 		data := &blogItem{}
// 		err := cur.Decode(data)
// 		if err != nil {
// 			return status.Errorf(
// 				codes.Internal,
// 				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
// 			)

// 		}
// 		stream.Send(&scampb.ListBlogResponse{Blog: dataToscampb(data)}) // Should handle err
// 	}
// 	if err := cur.Err(); err != nil {
// 		return status.Errorf(
// 			codes.Internal,
// 			fmt.Sprintf("Unknown internal error: %v", err),
// 		)
// 	}
// 	return nil
// }

func main() {
	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Connecting to MongoDB")
	// connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scam Service Started")
	collection = client.Database("mydb").Collection("PhoneNumbers")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	scampb.RegisterScamDatabaseServer(s, &server{})
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
	// First we close the connection with MongoDB:
	fmt.Println("Closing MongoDB Connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error on disconnection with MongoDB : %v", err)
	}

	// Finally, we stop the server
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("End of Program")
}
