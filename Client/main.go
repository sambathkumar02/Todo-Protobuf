package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func main() {

	todo_data := &Todo{
		Name:  "Learn Python",
		Type:  "Learning",
		Owner: "Sambath Kumar",
		Date:  "16/09/2020",
	}
	todo_buf, err := proto.Marshal(todo_data)

	//response, err := http.Get("https://www.gooogle.com")

	post_response, err := http.Post("http://0.0.0.0:80/add", "application/x-protobuf", bytes.NewBuffer(todo_buf))
	fmt.Print(post_response, err)
}
