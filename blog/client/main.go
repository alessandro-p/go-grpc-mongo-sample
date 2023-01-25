package main

import (
	"fmt"
	"log"

	"github.com/alessandro-p/go-grpc-mongo-sample/blog/client/services/api"
	"github.com/alessandro-p/go-grpc-mongo-sample/blog/client/services/utils"
	"github.com/alessandro-p/go-grpc-mongo-sample/blog/proto"
	"google.golang.org/grpc"
)

const address string = "localhost:50051"

func main() {
	tlsEnabled := false
	opts := utils.GetGrpcOptions(tlsEnabled)
	connection, err := grpc.Dial(
		address,
		opts...,
	)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer connection.Close()

	blogServiceClient := proto.NewPostServiceClient(connection)

	log.Println("Creating posts...")
	for i := 0; i < 5; i++ {
		api.CreatePost(
			blogServiceClient,
			fmt.Sprintf("John Doe #%d", i),
			fmt.Sprintf("Test Post #%d", i),
			fmt.Sprintf("Test Content #%d", i),
		)
	}

	log.Println("Retrieving post list:")
	posts := api.GetAllPosts(blogServiceClient)

	for i, post := range posts {
		log.Printf("Post %d: %v\n", i, post)
	}

	log.Println("Updating first post...")
	api.UpdatePost(
		blogServiceClient,
		posts[0].Id,
		"Updated Author",
		"Updated Title",
		"Updated Content",
	)

	log.Println("Getting first post...")
	post := api.GetPost(
		blogServiceClient,
		posts[0].Id,
	)

	log.Printf("Post now is: %v\n", post)

	log.Println("Clean up")
	for _, post := range posts {
		api.DeletePost(blogServiceClient, post.Id)
	}

	posts = api.GetAllPosts(blogServiceClient)
	if len(posts) != 0 {
		log.Fatal("Error, posts have not successfully been cleaned up")
	}

	log.Println("Clean up successful")
}
