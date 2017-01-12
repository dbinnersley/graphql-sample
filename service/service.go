package service

import "database/sql"
import "github.com/dbinnersley/graphql-sample/model"


var Users = []model.User{
	model.User{
		Id: "1",
		Name:"Derek",
		Height: 71,
		Weight: 155,
	},
	model.User{
		Id: "2",
		Name:"Derek2",
		Height: 70,
		Weight: 150,
	},
	model.User{
		Id: "3",
		Name:"Derek3",
		Height: 72,
		Weight: 145,
	},
}

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



type UserService interface{
	GetUserById(string) *model.User
}

type MemoryUserService struct{
	Users []model.User
}

func (m *MemoryUserService) GetUserById(userid string) (*model.User, error){
	for _,user := range m.Users{
		if user.Id == userid{
			return &user, nil
		}
	}
	return nil, nil
}

type MysqlUserService struct{
	DB *sql.DB
}

//This uses the sql to make the query for a user by id
func (m *MysqlUserService) GetUserById(userid string) (*model.User, error){
	prep,_ := m.DB.Prepare("Select * from user where id = ?")

	result, err := prep.Query(userid)

	defer result.Close()

	if err != nil{
		panic(err)
	}

	user := model.User{}

	hasnext := result.Next()
	if hasnext == true {
		err = result.Scan(&user.Id, &user.Name, &user.Height, &user.Weight)
		if err != nil {
			panic(err)
		}
	}

	return &user, nil

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

