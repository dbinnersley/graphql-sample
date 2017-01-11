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

var testvals = []User{
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

var userType = graphql.NewObject(graphql.ObjectConfig{
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

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserQuery",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type:userType,
			Args:graphql.FieldConfigArgument{
				"id" : &graphql.ArgumentConfig{
					Type: graphql.ID,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error){
				idQuery := params.Args["id"].(string)
				for _,user := range testvals{
					if user.Id == idQuery{
						return &user, nil
					}
				}
				return nil, nil
			},
		},
	},
})

var schema,_ = graphql.NewSchema(graphql.SchemaConfig{
	Query:queryType,
})


func main(){


	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})


	mux := http.NewServeMux()
	mux.Handle("/graphql", h)

	http.ListenAndServe(":8090", mux)

}

