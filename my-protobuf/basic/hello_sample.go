package basic

import (
	"fmt"

	"github.com/go-with-grpc/my-protobuf/protogen/basic"
)

func SayHello() {

	h := basic.Hello{
		Name: "Vallabh",
	}

	user := basic.User{
		Id:       1,
		IsActive: false,
		Username: "Vallabh Lakade",
		Password: []byte("I am the required password"),
		Emails:   []string{"vallabh.lakade@gmail.com"},
		Gender:   basic.Gender_GENDER_MALE,
	}

	//  ProtoToJSONConversion()
	fmt.Println("Hello", h.Name, &user)

}
