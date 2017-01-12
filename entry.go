package main

import (
	"net/http"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)


type User struct{
	Id string	`json:"id"`//Id of the user
	Name string	`json:"name"`//Name of the user
	Height int	`json:"height"`//Height of the user
	Weight int	`json:"weight"`//Weight of the user
}

type Post struct{
	Id string	`json:"id" bson:"_id"`		//Id of the user
	Content string	`json:"content" bson:"content"` //Name of the user
	UserId string	`json:"userid" bson:"userid"`   //UserId of which the user belongs
}

//////////////////////////////////////////////
//User services into the actual data retrieval
//////////////////////////////////////////////

type UserService interface{
	GetUserById(string) *User
}

type MemoryUserService struct{
	users []User
}

func (m *MemoryUserService) GetUserById(userid string) (*User, error){
	for _,user := range m.users{
		if user.Id == userid{
			return &user, nil
		}
	}
	return nil, nil
}

//////////////////////////////////////////////
//Post services into the actual data retrieval
//////////////////////////////////////////////

type PostService interface{
	GetPostById(string) *Post
	GetPostsByUser(string) []*Post
}

type MemoryPostService struct{
	posts []Post
}

func (m* MemoryPostService) GetPostById(postId string) (*Post, error){
	for _, post := range m.posts{
		if post.Id == postId{
			return &post, nil
		}
	}
	return nil, nil
}

func (m* MemoryPostService) GetPostsByUser(userId string) ([]*Post, error){
	results := make([]*Post,0)
	for index, _ := range m.posts{
		if m.posts[index].UserId == userId{
			results = append(results, &m.posts[index])
		}

	}
	return results, nil
}



//////////////////////////////////////////////
//Post services into the actual data retrieval
//////////////////////////////////////////////

func main(){

	users := []User{
		User{
			Id: "1",
			Name:"Derek",
			Height: 71,
			Weight: 155,
		},
		User{
			Id: "2",
			Name:"Derek2",
			Height: 70,
			Weight: 150,
		},
		User{
			Id: "3",
			Name:"Derek3",
			Height: 72,
			Weight: 145,
		},
	}
	posts := []Post{
		Post{
			Id: "1",
			Content: "This is the content of the first post",
			UserId: "1",
		},
		Post{
			Id: "2",
			Content: "This is the content of the second post",
			UserId: "2",
		},
		Post{
			Id: "3",
			Content: "What are you looking at me for!!!",
			UserId: "2",
		},
	}

	userservice := MemoryUserService{users:users}
	postservice := MemoryPostService{posts:posts}

	userType := graphql.NewObject(graphql.ObjectConfig{
		Name:"User",
		Fields: graphql.Fields{
			"id" :&graphql.Field{
				Type:graphql.ID,
			},
			"name":&graphql.Field{
				Type:graphql.String,
			},
			"height":&graphql.Field{
				Type:graphql.Int,
			},
			"weight":&graphql.Field{
				Type:graphql.Int,
			},
		},
	})

	postType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:graphql.ID,
			},
			"content": &graphql.Field{
				Type:graphql.String,
			},
			"userid": &graphql.Field{
				Type:graphql.ID,
			},
		},
	})

	postType.AddFieldConfig("user", &graphql.Field{
		Type:userType,
		Resolve: func (params graphql.ResolveParams) (interface{}, error){
			idQuery := params.Source.(*Post).Id
			return userservice.GetUserById(idQuery)
		},
	})


	userType.AddFieldConfig("posts", &graphql.Field{
		Type:graphql.NewList(postType),
		Resolve: func (params graphql.ResolveParams) (interface{}, error){
			userId := params.Source.(*User).Id
			return postservice.GetPostsByUser(userId)
		},
	})


	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "UserQuery",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:userType,
				Args:graphql.FieldConfigArgument{
					"id" : &graphql.ArgumentConfig{
						Type: graphql.ID,
					},
				},
				Resolve: func (params graphql.ResolveParams) (interface{}, error){
					idQuery := params.Args["id"].(string)
					return userservice.GetUserById(idQuery)
				},
			},
			"post": &graphql.Field{
				Type:postType,
				Args:graphql.FieldConfigArgument{
					"id" : &graphql.ArgumentConfig{
						Type:graphql.ID,
					},
				},
				Resolve: func (params graphql.ResolveParams) (interface{}, error){
					idQuery := params.Args["id"].(string)
					return postservice.GetPostById(idQuery)
				},

			},
		},
	})

	schema,_ := graphql.NewSchema(graphql.SchemaConfig{
		Query:queryType,
	})

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})


	mux := http.NewServeMux()
	mux.Handle("/graphql", h)

	http.ListenAndServe(":8090", mux)

}

