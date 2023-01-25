package api

import (
	"context"
	"io"
	"log"

	"github.com/alessandro-p/go-grpc-mongo-sample/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func CreatePost(
	client proto.PostServiceClient,
	author, title, content string,
) string {
	res, err := client.CreatePost(
		context.Background(), &proto.CreatePostRequest{
			Author:  author,
			Title:   title,
			Content: content,
		},
	)

	if err != nil {
		log.Fatalf("Could not create post: %v\n", err)
	}

	return res.Id
}

func GetPost(client proto.PostServiceClient, id string) *proto.Post {
	res, err := client.GetPost(
		context.Background(), &proto.PostId{
			Id: id,
		},
	)

	if err != nil {
		log.Fatalf("Could not get post: %v\n", err)
	}

	return res
}

func GetAllPosts(client proto.PostServiceClient) []*proto.Post {
	stream, err := client.ListPosts(
		context.Background(), &emptypb.Empty{},
	)

	if err != nil {
		log.Fatalf("Could not list posts: %v\n", err)
	}

	var result []*proto.Post
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		result = append(result, msg)
	}

	return result
}

func UpdatePost(client proto.PostServiceClient, id, author, title, content string) {
	_, err := client.UpdatePost(
		context.Background(), &proto.Post{
			Id:      id,
			Title:   title,
			Content: content,
			Author:  author,
		},
	)

	if err != nil {
		log.Fatalf("Could not update post: %v\n", err)
	}

	log.Printf("Post updated successfully")
}

func DeletePost(client proto.PostServiceClient, id string) {
	_, err := client.DeletePost(
		context.Background(), &proto.PostId{
			Id: id,
		},
	)

	if err != nil {
		log.Fatalf("Could not delete post: %v\n", err)
	}
}
