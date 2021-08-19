package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
)

var TodoList map[string]Todo

func AddTodo(response http.ResponseWriter, request *http.Request) {
	request_data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal("Error in getting request Data!")
	}
	todo := &Todo{}

	err = proto.Unmarshal(request_data, todo)
	if err != nil {
		log.Fatal("Error While decoding!")
	}
	fmt.Print(todo)

}

func ListTodo(response http.ResponseWriter, request *http.Request) {

}

func DeleteTodo(response http.ResponseWriter, request *http.Request) {

}

func main() {
	http.HandleFunc("/list", ListTodo)
	http.HandleFunc("/delete", DeleteTodo)
	http.HandleFunc("/add", AddTodo)
	fmt.Print("Server Listening!")
	log.Fatal(http.ListenAndServe(":80", nil))
}
