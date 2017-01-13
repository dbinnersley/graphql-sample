package wiring

import (
	"github.com/dbinnersley/graphql-sample/service"
	"github.com/dbinnersley/graphql-sample/model"
	"github.com/graphql-go/handler"
	"github.com/graphql-go/graphql"
)

type Wiring struct{
	Userservice    service.UserService
	Postservice    service.PostService
	Commentservice service.CommentService
	
}

func (w* Wiring) CreateHandler() *handler.Handler{

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
			"userid": &graphql.Field{
				Type:graphql.ID,
			},
		},
	})

	//get user by post
	postType.AddFieldConfig("user", &graphql.Field{
		Type:userType,
		Resolve: func (params graphql.ResolveParams) (interface{}, error){
			idQuery := params.Source.(*model.Post).UserId
			return w.Userservice.GetUserById(idQuery)
		},
	})

	//get comments by post
	postType.AddFieldConfig("comments", &graphql.Field{
		Type:graphql.NewList(commentType),
		Resolve: func (params graphql.ResolveParams) (interface{}, error){
			idQuery := params.Source.(*model.Post).Id
			return w.Commentservice.GetCommentsByPost(idQuery)
		},
	})

	//get posts by user
	userType.AddFieldConfig("posts", &graphql.Field{
		Type:graphql.NewList(postType),
		Resolve: func (params graphql.ResolveParams) (interface{}, error){
			userId := params.Source.(*model.User).Id
			return w.Postservice.GetPostsByUser(userId)
		},
	})

	//get comments by user
	userType.AddFieldConfig("comments", &graphql.Field{
		Type:graphql.NewList(commentType),
		Resolve: func (params graphql.ResolveParams) (interface{}, error){
			userId := params.Source.(*model.User).Id
			return w.Commentservice.GetCommentsByUser(userId)
		},
	})

	//get user by comment
	commentType.AddFieldConfig("user", &graphql.Field{
		Type:userType,
		Resolve: func (params graphql.ResolveParams) (interface{}, error){
			userId := params.Source.(*model.Comment).UserId
			return w.Userservice.GetUserById(userId)
		},
	})

	//get post by comment
	commentType.AddFieldConfig("post", &graphql.Field{
		Type:userType,
		Resolve: func (params graphql.ResolveParams) (interface{}, error){
			postId := params.Source.(*model.Comment).PostId
			return w.Postservice.GetPostById(postId)
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
						return w.Userservice.GetUserById(idQuery)
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
						return w.Postservice.GetPostById(idQuery)
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
						return w.Commentservice.GetCommentById(idQuery)
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

	return h
}
