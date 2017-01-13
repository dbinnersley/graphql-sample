package service

import (
	"github.com/dbinnersley/graphql-sample/model"

)

//This is the service to get access comments on posts

type CommentService interface {
	GetCommentById(string) (interface{}, error)	//This will get comments by comment id
	GetCommentsByPost(string) (interface{}, error)	//This will get all comments for a specific post
	GetCommentsByAuthor(string) (interface{}, error)	//this will get all comments for a specific user
}

func CreateCommentService(servicetype string, connstring string) CommentService{
	var commentservice CommentService
	switch servicetype {
	case "memory":
		commentservice = &MemoryCommentService{Comments:comments}
	default:
		panic("Invalid service type specified")
	}

	return commentservice
}


var comments = []model.Comment{
	model.Comment{
		Id: "1",
		Content: "This is my first Comment!!!!",
		AuthorId: "1",
		PostId: "3",
	},
	model.Comment{
		Id: "2",
		Content: "This post sucks!!!!",
		AuthorId: "2",
		PostId: "3",
	},
	model.Comment{
		Id: "3",
		Content: "This is awesome!!!!",
		AuthorId: "4",
		PostId: "3",
	},
	model.Comment{
		Id: "4",
		Content: "Were you drunk when you wrote this!!!!",
		AuthorId: "4",
		PostId: "2",
	},
}


type MemoryCommentService struct{
	Comments []model.Comment
}

func (m* MemoryCommentService) GetCommentById(commentId string) (interface{}, error){
	for _, comment := range m.Comments {
		if comment.Id == commentId {
			return &comment, nil
		}
	}
	return nil, nil
}

func (m* MemoryCommentService) GetCommentsByPost(postId string) (interface{}, error){
	results := make([]*model.Comment,0)
	for index, _ := range m.Comments {
		if m.Comments[index].PostId == postId {
			results = append(results, &m.Comments[index])
		}
	}
	return results, nil
}

func (m* MemoryCommentService) GetCommentsByAuthor(authorId string) (interface{}, error){
	results := make([]*model.Comment,0)
	for index, _ := range m.Comments {
		if m.Comments[index].AuthorId == authorId {
			results = append(results, &m.Comments[index])
		}

	}
	return results, nil
}