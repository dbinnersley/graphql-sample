package service

import (
	"github.com/dbinnersley/graphql-sample/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PostService interface{
	GetPostById(string) (interface{}, error)
	GetPostsByUser(string) (interface{}, error)
}

func CreatePostService(servicetype string, connstring string) PostService{
	var postservice PostService
	switch servicetype {
	case "memory":
		postservice = &MemoryPostService{Posts:posts}
	case "mongodb":
		session,error := mgo.Dial(connstring)
		if error != nil{
			panic(error)
		}
		postservice = &MongoPostService{Conn:session}
	default:
		panic("Invalid service type specified")
	}

	return postservice
}


////////////////////////////////////////
// Memory Post Services
////////////////////////////////////////

var posts = []model.Post{
	model.Post{
		Id: "1",
		Content: "This is the content of the first post",
		UserId: "1",
	},
	model.Post{
		Id: "2",
		Content: "This is the content of the second post",
		UserId: "2",
	},
	model.Post{
		Id: "3",
		Content: "What are you looking at me for!!!",
		UserId: "2",
	},
}


type MemoryPostService struct{
	Posts []model.Post
}

func (m* MemoryPostService) GetPostById(postId string) (interface{}, error){
	for _, post := range m.Posts {
		if post.Id == postId{
			return &post, nil
		}
	}
	return nil, nil
}

func (m* MemoryPostService) GetPostsByUser(userId string) (interface{}, error){
	results := make([]*model.Post,0)
	for index, _ := range m.Posts {
		if m.Posts[index].UserId == userId{
			results = append(results, &m.Posts[index])
		}

	}
	return results, nil
}

////////////////////////////////////////
// Mongo Post Services
////////////////////////////////////////

type MongoPostService struct{
	Conn *mgo.Session
}

func (m* MongoPostService) GetPostById(postId string) (interface{}, error){
	conn := m.Conn.Copy()
	defer conn.Close()

	coll := conn.DB("graphql_sample").C("post")

	post := &model.Post{}
	err := coll.Find(bson.M{"_id":postId}).One(post)
	if err != nil{
		return nil, nil
	}

	return post, nil
}

func (m* MongoPostService) GetPostsByUser(userId string) (interface{}, error){
	conn := m.Conn.Copy()
	defer conn.Close()

	coll := conn.DB("graphql_sample").C("post")

	posts := []*model.Post{}
	err := coll.Find(bson.M{"userid":userId}).All(&posts)
	if err != nil{
		panic(err)
	}

	return posts,nil
}