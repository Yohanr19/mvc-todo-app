package main

import (
	"github.com/yohanr19/mvc-todo-app/pkg/controlers"
	"log"
	"net/http"
)

func main() {
	TaskControler := controlers.TaskControler{}
	err := TaskControler.Init()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/", TaskControler.GetAll)
	mux.HandleFunc("/insert", TaskControler.Insert)
	mux.HandleFunc("/state", TaskControler.SetState)
	mux.HandleFunc("/delete", TaskControler.Delete)
	log.Fatal(http.ListenAndServe("localhost:3001", mux))
}
