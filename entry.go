package main

import (
	"net/http"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/dbinnersley/graphql-sample/service"
	"github.com/dbinnersley/graphql-sample/model"
)

func main(){


	//Can create any type of service for accessing users really, but here we use mysql
	userservice := service.CreateUserService("mysql", "root@tcp(mysql:3306)/graphql_sample")

	//This is the service that uses mongodb for
	postservice := service.CreatePostService("mongodb", "mongodb://mongo:27017/graphql_sample")

	//This is the service for accessing the comments. Currently from memory
	commentservice := service.CreateCommentService("memory", "")

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

	commentType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:graphql.ID,
			},
			"content": &graphql.Field{
				Type:graphql.String,
			},
			"postid": &graphql.Field{
				Type:graphql.ID,
			},
			"authorid": &graphql.Field{
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
					idQuery, ok := params.Args["id"].(string)
					if ok == true {
						return userservice.GetUserById(idQuery)
					}
					return nil, nil
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
					idQuery, ok := params.Args["id"].(string)
					if ok == true {
						return postservice.GetPostById(idQuery)
					}
					return nil, nil
				},

			},
			"comment": &graphql.Field{
				Type:commentType,
				Args:graphql.FieldConfigArgument{
					"id" : &graphql.ArgumentConfig{
						Type:graphql.ID,
					},
				},
				Resolve: func (params graphql.ResolveParams) (interface{}, error){
					idQuery, ok := params.Args["id"].(string)
					if ok == true {
						return commentservice.GetCommentById(idQuery)
					}
					return nil, nil
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

