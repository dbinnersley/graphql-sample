package main

import (
	"net/http"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/dbinnersley/graphql-sample/service"
	"github.com/dbinnersley/graphql-sample/model"


	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func main(){

	db,err := sql.Open("mysql", "root@tcp(mysql:3306)/graphql_sample")
	if err != nil{
		panic(err)
	}
	userservice := service.MysqlUserService{DB:db}

	//userservice := MemoryUserService{users:users}
	postservice := service.MemoryPostService{Posts:service.Posts}

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
			idQuery := params.Source.(*model.Post).UserId
			return userservice.GetUserById(idQuery)
		},
	})


	userType.AddFieldConfig("posts", &graphql.Field{
		Type:graphql.NewList(postType),
		Resolve: func (params graphql.ResolveParams) (interface{}, error){
			userId := params.Source.(*model.User).Id
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

