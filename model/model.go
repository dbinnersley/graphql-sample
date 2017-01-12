package model

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