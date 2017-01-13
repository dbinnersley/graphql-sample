package main

import (
	"net/http"
	"github.com/dbinnersley/graphql-sample/wiring"
	"github.com/dbinnersley/graphql-sample/service"
)



func main(){

	//Can create any type of service for accessing users really, but here we use mysql
	userservice := service.CreateUserService("mysql", "root@tcp(mysql:3306)/graphql_sample")

	//This is the service that uses mongodb for
	postservice := service.CreatePostService("mongodb", "mongodb://mongo:27017/graphql_sample")

	//This is the service for accessing the comments. Currently from memory
	commentservice := service.CreateCommentService("cassandra", "cassandra")

	wiring := &wiring.Wiring{Userservice : userservice,
				Commentservice : commentservice,
				Postservice : postservice}

	handler := wiring.CreateHandler()

	mux := http.NewServeMux()
	mux.Handle("/graphql", handler)

	http.ListenAndServe(":8090", mux)

}

