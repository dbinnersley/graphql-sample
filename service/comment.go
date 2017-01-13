package service

import "github.com/dbinnersley/graphql-sample/model"

//This is the service to get access comments on posts

type CommentService interface {
	GetCommentById(string) (interface{}, error)	//This will get comments by comment id
	GetCommentsByPost(string) (interface{}, error)	//This will get all comments for a specific post
	GetCommentsByUser(string) (interface{}, error)	//this will get all comments for a specific user
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