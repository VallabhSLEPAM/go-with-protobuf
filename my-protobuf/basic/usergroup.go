package basic

import (
	"fmt"

	"github.com/go-with-grpc/my-protobuf/protogen/basic"
)

func GetUserGroup() {

	batMan := basic.User{
		Id:       1,
		Username: "Batman",
		IsActive: true,
		Password: []byte("I am Batman!"),
		Address: &basic.Address{
			City:    "Gotham City",
			Country: "USA",
		},
	}

	nightWing := basic.User{
		Id:       2,
		Username: "Nightwing",
		IsActive: true,
		Password: []byte("I am Nightwing!"),
		Address: &basic.Address{
			City:    "Gotham City",
			Country: "USA",
		},
	}

	robin := basic.User{
		Id:       3,
		Username: "Robin",
		IsActive: true,
		Password: []byte("I am Robin!"),
		Address: &basic.Address{
			City:    "Unknown",
			Country: "USA",
		},
	}

	userGroup := basic.UserGroup{
		GroupId:     1,
		GroupName:   "Batman",
		Users:       []*basic.User{&batMan, &robin, &nightWing},
		Description: "Batman Family",
	}

	fmt.Println(&userGroup)
}
