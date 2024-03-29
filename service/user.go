package service

import (
	"github.com/dbinnersley/graphql-sample/model"
	 "database/sql"
 	_ "github.com/go-sql-driver/mysql"
)


//Interface to the actual user service
//This defines the actions that able to happen to get user objects
type UserService interface{
	GetUserById(string) (interface{}, error)
}

//Constructor for the user service
func CreateUserService(servicetype string, connstring string) UserService{
	var userservice UserService
	switch servicetype{
	case "memory":
		userservice = &MemoryUserService{Users:users}
		break
	case "mysql":
		db, err := sql.Open("mysql", connstring)
		if err != nil {
			panic(err)
		}
		userservice = &SqlUserService{DB:db}
		break
	default:
		panic("Invalid servicetype specified")
	}
	return userservice

}

////////////////////////////////////////
// Memory User Service Values
////////////////////////////////////////

//Default constant for a memory value. Not exported, so it is private to the services
var users = []model.User{
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

//Service with an in memory backend
type MemoryUserService struct{
	Users []model.User
}


//Get the User by Id using the memory service
func (m *MemoryUserService) GetUserById(userid string) (interface{}, error){
	for _,user := range m.Users{
		if user.Id == userid{
			return &user, nil
		}
	}
	return nil, nil
}


////////////////////////////////////////
// Mysql User Service Values
////////////////////////////////////////


//Service with a SQL backend
type SqlUserService struct{
	DB *sql.DB
}

//This uses the sql to make the query for a user by id. Returns a user model if found, else nil
func (m *SqlUserService) GetUserById(userid string) (interface{}, error){

	prep,err := m.DB.Prepare("Select * from user where id = ?")

	if err != nil{
		panic(err)
	}

	result, err := prep.Query(userid)

	defer result.Close()

	if err != nil{
		panic(err)
	}

	var user *model.User = &model.User{}

	hasnext := result.Next()
	if hasnext == true {
		err = result.Scan(&user.Id, &user.Name, &user.Height, &user.Weight)
		if err != nil {
			panic(err)
		}
	}else{
		return nil, nil
	}

	return user, nil

}


