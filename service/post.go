package service

import "github.com/dbinnersley/graphql-sample/model"

var Posts = []model.Post{
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


//////////////////////////////////////////////
//Post services into the actual data retrieval
//////////////////////////////////////////////

type PostService interface{
	GetPostById(string) *model.Post
	GetPostsByUser(string) []*model.Post
}

type MemoryPostService struct{
	Posts []model.Post
}

func (m* MemoryPostService) GetPostById(postId string) (*model.Post, error){
	for _, post := range m.Posts{
		if post.Id == postId{
			return &post, nil
		}
	}
	return nil, nil
}

func (m* MemoryPostService) GetPostsByUser(userId string) ([]*model.Post, error){
	results := make([]*model.Post,0)
	for index, _ := range m.Posts{
		if m.Posts[index].UserId == userId{
			results = append(results, &m.Posts[index])
		}

	}
	return results, nil
}