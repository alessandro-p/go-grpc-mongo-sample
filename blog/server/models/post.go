package models

import (
	"github.com/alessandro-p/go-grpc-mongo-sample/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Author  string             `bson:"author"`
	Title   string             `bson:"title"`
	Content string             `bson:"content"`
}

func (p *Post) ToProto() *proto.Post {
	return &proto.Post{
		Id:      p.ID.Hex(),
		Author:  p.Author,
		Content: p.Content,
		Title:   p.Title,
	}
}
