package controlers

import (
	"github.com/yohanr19/mvc-todo-app/models/psqlStore"
	"net/http"
	"encoding/json"
	"log"
	"strconv"
)

type TaskControler struct {
	 Store psqlStore.TaskStore
}
type ResponseTask struct{
	Id string `json:"id"`
	IsActive bool `json:"is_active"`
	Text string	`json:"text"`
}
func (tc *TaskControler)Init() error{
	error := tc.Store.InitDB()
	return error
}

func (tc *TaskControler)Insert(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err!=nil {
		http.Error(w, "Internal Server Error",http.StatusInternalServerError)
		log.Printf("found while decoding request body: %v",err)
		return
	}
	text := data.(string)
	tc.Store.InsertTask(text)
}
func (tc *TaskControler)GetAll(w http.ResponseWriter, r *http.Request) {
	tasks := tc.Store.GetTasks()
	var response []ResponseTask
	for _, v := range tasks {
		response = append(response, ResponseTask{
			Id: strconv.Itoa(int(v.ID)),
			IsActive: v.IsActive,
			Text:  v.Text,
		})
	} 
	err := json.NewEncoder(w).Encode(response)
	if err!=nil {
		http.Error(w, "Internal Server Error",http.StatusInternalServerError)
		log.Printf("found while encoding json to response writer %v",err)
		return
	}
}
func (tc *TaskControler)SetState(w http.ResponseWriter, r *http.Request) {
		var data struct{
			Id string `json:"id"`
			IsActive bool `json:"is_active"`
		}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err!=nil {
			http.Error(w, "Internal Server Error",http.StatusInternalServerError)
			log.Printf("found while decoding json from  body on SetState %v",err)
			return
		}
		tc.Store.SetIsActive(data.Id,data.IsActive)
}