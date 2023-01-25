package repository

import (
	"context"
	"fmt"

	"github.com/alessandro-p/go-grpc-mongo-sample/blog/proto"
	"github.com/alessandro-p/go-grpc-mongo-sample/blog/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IPostRepository interface {
	CreatePost(ctx context.Context, author, title, content string) (string, error)
	FindOne(ctx context.Context, postId string) (*proto.Post, error)
	FindAll(ctx context.Context) (*mongo.Cursor, error)
	DeleteOne(ctx context.Context, postId string) (int64, error)
	UpdatePost(ctx context.Context, postId, author, title, content string) (int64, error)
}

type PostRepository struct {
	*mongo.Client
}

func (r *PostRepository) CreatePost(
	ctx context.Context,
	author, title, content string,
) (string, error) {
	res, err := r.Client.Database(DefaultDatabase).Collection(PostsCollection).InsertOne(
		ctx, &models.Post{
			Author:  author,
			Title:   title,
			Content: content,
		},
	)

	if err != nil {
		return "",
			status.Errorf(codes.Internal, fmt.Sprintf("Unable to insert document, %v\n", err))
	}

	// casting the interface to ObjectId
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "",
			status.Errorf(codes.Internal, fmt.Sprintf("Unable to parse object id, %v\n", err))
	}

	return oid.Hex(), nil
}

func (r *PostRepository) FindOne(ctx context.Context, postId string) (*proto.Post, error) {
	oid, err := primitive.ObjectIDFromHex(postId)

	if err != nil {
		return nil,
			status.Errorf(codes.Internal, fmt.Sprintf("Unable to convert HEX to Object Id. Error: %v\n", err))
	}

	filter := bson.M{"_id": oid}
	res := r.Client.Database(DefaultDatabase).Collection(PostsCollection).FindOne(ctx, filter)

	data := &models.Post{}
	if err := res.Decode(data); err != nil {
		return nil,
			status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find post with id provided: %v", err))
	}

	return data.ToProto(), err
}

func (r *PostRepository) FindAll(ctx context.Context) (*mongo.Cursor, error) {
	cursor, err := r.Client.Database(DefaultDatabase).Collection(PostsCollection).Find(ctx, primitive.D{{}})

	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Unable to complete findAll: %v", err))
	}

	return cursor, nil
}

func (r *PostRepository) UpdatePost(ctx context.Context, postId, author, title, content string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(postId)

	if err != nil {
		return 0,
			status.Errorf(codes.Internal, fmt.Sprintf("Unable to convert HEX to Object Id. Error: %v\n", err))
	}

	filter := bson.M{"_id": oid}
	post := &models.Post{
		Author:  author,
		Title:   title,
		Content: content,
	}

	res, err := r.Client.Database(DefaultDatabase).Collection(PostsCollection).UpdateOne(
		ctx,
		filter,
		bson.M{"$set": post},
	)

	if err != nil {
		return 0, status.Errorf(codes.Internal, fmt.Sprintf("Could not update document: %v\n", err))
	}

	return res.ModifiedCount, nil
}

func (r *PostRepository) DeleteOne(ctx context.Context, postId string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(postId)

	if err != nil {
		return 0, status.Errorf(codes.Internal, fmt.Sprintf("Unable to convert HEX to Object Id. Error: %v\n", err))
	}

	filter := bson.M{"_id": oid}
	res, err := r.Client.Database(DefaultDatabase).Collection(PostsCollection).DeleteOne(ctx, filter)

	if err != nil {
		return 0, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find post with id provided: %v", err))
	}

	return res.DeletedCount, nil
}
