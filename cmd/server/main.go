package main

import (
	"bufio"
	"log"
	"os"

	"github.com/yohanr19/mvc-todo-app/models/psqlStore"
)

func main(){

	//Before creating the server, we will use this app to test the packages
	scanner := bufio.NewScanner(os.Stdin)
	taskStore := &psqlStore.TaskStore{}
	err := taskStore.InitDB()
	if err!=nil{
		log.Fatal(err)
	}
	for scanner.Scan() {
		text := scanner.Text()
		taskStore.InsertTask(text)
	}
	
}