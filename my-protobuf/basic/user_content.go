package basic

import (
	"log"
	"os"

	"github.com/go-with-grpc/my-protobuf/protogen/basic"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteUserContent(fileName string) {

	userContent := basic.UserContent{
		UserContentId: 1,
		Slug:          "I dont know what it means",
		// Title:         "This is title",
		HtmlContent: "<i>There is some HTML content here</i>",
		// AuthorId:      1,
		Category:    "some category",
		SubCategory: "some sub-category",
	}

	userContentBytes, err := proto.Marshal(&userContent)
	if err != nil {
		log.Fatal("Error while marshalling user content: ", err)
	}
	err = os.WriteFile(fileName, userContentBytes, 0644)
	if err != nil {
		log.Fatal("Error while writing data to file: ", err)
	}
	log.Printf("Writing to file with name %v done\n", fileName)
}

func ReadUserContentFromFile(fileName string) {

	userContentInBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading data from the file: ", err)
	}

	var userContent basic.UserContent //proto.Message
	err = proto.Unmarshal(userContentInBytes, &userContent)
	if err != nil {
		log.Fatal("Error unmarshalling data read from the file: ", err)
	}

	jsonBytes, err := protojson.Marshal(&userContent)
	if err != nil {
		log.Fatal("Error marshalling user content to json format: ", err)
	}
	log.Printf("Data read from file %v :%v \n, %v", fileName, &userContent, string(jsonBytes))
}
