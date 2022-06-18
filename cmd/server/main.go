package main

import (
	//"bufio"
	"fmt"
	"log"
	//"os"

	"github.com/yohanr19/mvc-todo-app/models/psqlStore"
)

func main(){

	//Before creating the server, we will use this app to test the packages
	//scanner := bufio.NewScanner(os.Stdin)
	taskStore := &psqlStore.TaskStore{}
	err := taskStore.InitDB()
	if err!=nil{
		log.Fatal(err)
	}
	/*
	for scanner.Scan() {
		text := scanner.Text()
		taskStore.InsertTask(text)
	}
	*/
	/*
	for scanner.Scan() {
		id := scanner.Text()
		taskStore.SetIsActive(id,false)
	*/
	tasks := taskStore.GetTasks()
	for _, v := range tasks{
		fmt.Printf("id:%v, text: %s, active: %v \n", v.ID, v.Text, v.IsActive)
	}
}