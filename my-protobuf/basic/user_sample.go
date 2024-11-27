package basic

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/go-with-grpc/my-protobuf/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

func ProtoToJSONConversion() {

	address := basic.Address{
		City:    "Nasik",
		Country: "India",
		Street:  "Shri Krishna Nagar",
		Coordinate: &basic.Address_Coordinate{
			Latitude:  45.32,
			Longitude: 54.23,
		},
	}

	comm := randonCommunicationChannel()

	user := basic.User{
		Id:                   1,
		Username:             "Vallabh Lakade",
		IsActive:             true,
		Password:             []byte("Password"),
		Gender:               basic.Gender_GENDER_MALE,
		Emails:               []string{"vallabh.lakade@gmail.com", "vallabh.lakade@epam.com"},
		Address:              &address,
		CommunicationChannel: &comm,
		SkillRating:          map[string]uint32{"maths": 75, "english": 70},
	}

	userInBytes, err := protojson.Marshal(&user)
	if err != nil {
		log.Fatal("Some error while marshalling User using protojson")
	}
	fmt.Println("Marshalled User from protogen", string(userInBytes))
}

func JSONToProto() {

	jsonString := `{"id":2, "username":"Vallabh S Lakade", "is_active":true, "password":"UGFzc3fvcmQ=", "emails":["vallabh.sl@gmail.com", "vallabh.slakade@epam.com"]}`

	user := basic.User{}

	err := protojson.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		log.Fatal("Error unmarshalling jsonString to Proto struct")
	}
	fmt.Println("Unmarshalled user:", &user)
}

func randonCommunicationChannel() anypb.Any {

	rand.Seed(time.Now().UnixNano())

	paperMail := basic.PaperMail{
		PaperMailAddress: "paperMail address",
	}

	socialMediaPlatform := basic.SocialMediaPlatform{
		SocialMediaPlatform: "Instagram",
		SocialMediaUsername: "insta_user_name",
	}

	instantMessagingPlatform := basic.InstantMessagingPlatform{
		InstantMessagingPlatform: "Whatsapp",
		InstantMessagingUsername: "whatsapp_username",
	}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMediaPlatform, proto.MarshalOptions{})
	case 2:
		anypb.MarshalFrom(&a, &instantMessagingPlatform, proto.MarshalOptions{})
	}

	return a

}

func UnMarshalFromKnown() {
	socialMedia := basic.SocialMediaPlatform{
		SocialMediaPlatform: "Instagram",
		SocialMediaUsername: "unknown",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	sm := basic.SocialMediaPlatform{}

	if err := a.UnmarshalTo(&sm); err != nil {
		log.Fatal("Error unmarshalling the data from anypb")
	}
	fmt.Println("Unmarshalled data: ", &sm)
}

func UnmarshalAnyToUnknown() {
	a := randonCommunicationChannel()

	var unmarshalled protoreflect.ProtoMessage

	unmarshalled, err := a.UnmarshalNew()
	if err != nil {
		return
	}

	log.Println("Unmarshalled as: ", unmarshalled.ProtoReflect().Descriptor().Name())

	jsonBytes, _ := protojson.Marshal(unmarshalled)
	log.Println(string(jsonBytes))

}

func BasicOneof() {

	socialMessaging := basic.SocialMediaPlatform{
		SocialMediaPlatform: "BasicOneof-SMP",
		SocialMediaUsername: "BasicOneof-SMU",
	}

	ecomm := basic.User_SocialMedia{
		SocialMedia: &socialMessaging,
	}

	user := &basic.User{
		Id:                    1,
		ElectronicCommChannel: &ecomm,
	}

	fmt.Println(user)

}

func WriteProtoToFile(msg proto.Message, fileName string) {

	msgInBytes, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("error marshalling proto message")
	}

	err = os.WriteFile(fileName, msgInBytes, 0644)
	if err != nil {
		log.Fatal("error writing to file:", err)
	}

}

func ReadProtoFromFile(fileName string, msg proto.Message) {

	bytesToRead, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("error reading from file:", err)
	}

	err = proto.Unmarshal(bytesToRead, msg)
	if err != nil {
		log.Fatal("error unmarshalling data read from file:", err)
	}

	fmt.Println("Data read from file: ", msg)
}

func WriteToAFile(fileName string) {

	address := basic.Address{
		City:    "Nasik",
		Country: "India",
		Street:  "Shri Krishna Nagar",
		Coordinate: &basic.Address_Coordinate{
			Latitude:  45.32,
			Longitude: 54.23,
		},
	}

	comm := randonCommunicationChannel()

	dummyUser := basic.User{
		Id:                   1,
		Username:             "Dummy",
		IsActive:             true,
		Password:             []byte("Password"),
		Gender:               basic.Gender_GENDER_MALE,
		Emails:               []string{"dummy@gmail.com"},
		Address:              &address,
		CommunicationChannel: &comm,
		SkillRating:          map[string]uint32{"maths": 75, "english": 70},
	}

	WriteProtoToFile(&dummyUser, fileName)

}

func WriteJsonToFile(fileName string) {

	address := basic.Address{
		City:    "Nasik",
		Country: "India",
		Street:  "Shri Krishna Nagar",
		Coordinate: &basic.Address_Coordinate{
			Latitude:  45.32,
			Longitude: 54.23,
		},
	}

	comm := randonCommunicationChannel()

	dummyUser := basic.User{
		Id:                   1,
		Username:             "Dummy",
		IsActive:             true,
		Password:             []byte("Password"),
		Gender:               basic.Gender_GENDER_MALE,
		Emails:               []string{"dummy@gmail.com"},
		Address:              &address,
		CommunicationChannel: &comm,
		SkillRating:          map[string]uint32{"maths": 75, "english": 70},
	}

	dummyUserBytes, err := protojson.Marshal(&dummyUser)
	if err != nil {
		log.Fatal("error while marshalling data:", err)
	}

	err = os.WriteFile(fileName, dummyUserBytes, 0644)
	if err != nil {
		log.Fatal("error writing data to file:", err)
	}

}

func ReadJSONFromFile(fileName string) {

	jsonBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("error reading data from file:", err)
	}

	var user basic.User
	err = protojson.Unmarshal(jsonBytes, &user)
	if err != nil {
		log.Fatal("error unmarshalling data read from file:", err)
	}

	fmt.Println("Data read from file: ", user)
}
