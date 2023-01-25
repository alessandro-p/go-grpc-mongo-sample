package api

import (
	"context"
	"fmt"
	"log"

	"github.com/alessandro-p/go-grpc-mongo-sample/blog/proto"
	"github.com/alessandro-p/go-grpc-mongo-sample/blog/server/models"
	"github.com/alessandro-p/go-grpc-mongo-sample/blog/server/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostController struct {
	proto.PostServiceServer
	PostRepository repository.IPostRepository
}

func (pc *PostController) CreatePost(
	ctx context.Context,
	request *proto.CreatePostRequest,
) (*proto.PostId, error) {
	log.Printf("CreatePost called with %v\n", request)

	oid, err := pc.PostRepository.CreatePost(
		ctx,
		request.Author,
		request.Title,
		request.Content,
	)

	if err != nil {
		return nil, err
	}

	return &proto.PostId{
		Id: oid,
	}, nil
}

func (pc *PostController) GetPost(ctx context.Context, request *proto.PostId) (*proto.Post, error) {
	log.Printf("GetPost called with %v\n", request)

	return pc.PostRepository.FindOne(ctx, request.Id)
}

func (pc *PostController) ListPosts(e *emptypb.Empty, stream proto.PostService_ListPostsServer) error {
	cursor, err := pc.PostRepository.FindAll(context.Background())

	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		data := &models.Post{}
		err := cursor.Decode(data)

		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Error while decoding data: %v", err))
		}

		stream.Send(data.ToProto())
	}

	if err = cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}

	return nil
}

func (pc *PostController) UpdatePost(ctx context.Context, request *proto.Post) (*emptypb.Empty, error) {
	res, err := pc.PostRepository.UpdatePost(
		ctx,
		request.Id,
		request.Author,
		request.Title,
		request.Content,
	)

	if err != nil {
		return &emptypb.Empty{}, err
	}

	if res == 0 {
		return &emptypb.Empty{}, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Unable to update the document %s. No document was updated", request.Id),
		)
	}

	return &emptypb.Empty{}, nil
}

func (pc *PostController) DeletePost(ctx context.Context, request *proto.PostId) (*emptypb.Empty, error) {
	log.Printf("DeletePost called with %v\n", request)

	res, err := pc.PostRepository.DeleteOne(ctx, request.Id)

	if res == 0 {
		return &emptypb.Empty{}, status.Error(
			codes.NotFound,
			fmt.Sprintf("Unable to delete document %s. Not found", request.Id),
		)
	}

	return &emptypb.Empty{}, err
}
