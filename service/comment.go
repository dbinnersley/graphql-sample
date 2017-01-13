package service

import (
	"github.com/dbinnersley/graphql-sample/model"
	"github.com/gocql/gocql"
)

//This is the service to get access comments on posts

type CommentService interface {
	GetCommentById(string) (interface{}, error)	//This will get comments by comment id
	GetCommentsByPost(string) (interface{}, error)	//This will get all comments for a specific post
	GetCommentsByUser(string) (interface{}, error)	//this will get all comments for a specific user
}

func CreateCommentService(servicetype string, connstring string) CommentService{
	var commentservice CommentService
	switch servicetype {
	case "memory":
		commentservice = &MemoryCommentService{Comments:comments}
	case "cassandra":
		cluster  := gocql.NewCluster(connstring)
		cluster.Keyspace = "graphql_sample"
		session,err := cluster.CreateSession()
		if err != nil{
			panic(err)
		}
		commentservice= &CassandraCommentService{Conn:session}
	default:
		panic("Invalid service type specified")
	}

	return commentservice
}

////////////////////////////////////////
// Memory Comment Service Values
////////////////////////////////////////

var comments = []model.Comment{
	model.Comment{
		Id: "1",
		Content: "This is my first Comment!!!!",
		UserId: "1",
		PostId: "3",
	},
	model.Comment{
		Id: "2",
		Content: "This post sucks!!!!",
		UserId: "2",
		PostId: "3",
	},
	model.Comment{
		Id: "3",
		Content: "This is awesome!!!!",
		UserId: "4",
		PostId: "3",
	},
	model.Comment{
		Id: "4",
		Content: "Were you drunk when you wrote this!!!!",
		UserId: "4",
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

func (m* MemoryCommentService) GetCommentsByUser(userId string) (interface{}, error){
	results := make([]*model.Comment,0)
	for index, _ := range m.Comments {
		if m.Comments[index].UserId == userId {
			results = append(results, &m.Comments[index])
		}

	}
	return results, nil
}

////////////////////////////////////////
// Cassandra Comment Service Values
////////////////////////////////////////

type CassandraCommentService struct {
	Conn *gocql.Session
}

func (c* CassandraCommentService) GetCommentById(commentId string) (interface{}, error){

	comment := &model.Comment{}
	query := c.Conn.Query(`SELECT * FROM comment_by_id WHERE id = ?`, commentId)

	err := query.Scan(&comment.Id,&comment.Content,
		&comment.UserId,&comment.PostId)

	if err != nil{
		panic(err)
	}

	return comment, nil
}

func (c* CassandraCommentService) GetCommentsByPost(postId string) (interface{}, error){

	//Create a zero size length comment
	comments := make([]*model.Comment, 0)
	query := c.Conn.Query(`SELECT * FROM comment_by_id WHERE postid = ? allow filtering`, postId)

	iter := query.Iter()
	newcomment := &model.Comment{}
	for iter.Scan(&newcomment.Id, &newcomment.Content, &newcomment.UserId, &newcomment.PostId){
		comments = append(comments, newcomment)
		newcomment = &model.Comment{}
	}

	return comments, nil
}

func (c* CassandraCommentService) GetCommentsByUser(commentId string) (interface{}, error){

	//Create a zero size length comment
	comments := make([]*model.Comment, 0)
	query := c.Conn.Query(`SELECT * FROM comment_by_id WHERE commentid = ? allow filtering`, commentId)

	iter := query.Iter()
	newcomment := &model.Comment{}
	for iter.Scan(&newcomment.Id, &newcomment.Content, &newcomment.UserId, &newcomment.PostId){
		comments = append(comments, newcomment)
		newcomment = &model.Comment{}
	}

	return comments, nil
}