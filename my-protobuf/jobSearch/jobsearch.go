package jobsearch

import (
	"log"

	"github.com/go-with-grpc/my-protobuf/protogen/basic"
	"github.com/go-with-grpc/my-protobuf/protogen/dummy"
)

func SearchApplication() {

	jobSearchApplication := &dummy.Application{
		ApplicationId: 1,
		ApplicantName: "Vallabh",
		Phone:         uint32(918788469),
		Email:         "someemail.gmail.com",
	}

	softwareApplication := &basic.Application{
		Name:      "VSCode",
		Version:   "15.3",
		Platforms: []string{"Windiws", "Ubuntu", "Mac"},
	}

	log.Println("job application: ", jobSearchApplication)
	log.Println("software application: ", softwareApplication)

}
