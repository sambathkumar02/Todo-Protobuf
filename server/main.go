package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/golang/protobuf/proto"
)

var TodoList []TodoFull

func GenerateID() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)

}

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
	id := GenerateID()
	full_object := TodoFull{Id: id, Todo: todo}
	TodoList = append(TodoList, full_object)

}

func ListTodo(response http.ResponseWriter, request *http.Request) {

	for i := range TodoList {
		data, _ := proto.Marshal(TodoList[i])
		response.Write(data)

	}

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
