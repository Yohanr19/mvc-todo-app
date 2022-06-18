package controlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/yohanr19/mvc-todo-app/models/psqlStore"
)

type TaskControler struct {
	Store psqlStore.TaskStore
}
type ResponseTask struct {
	Id       string `json:"id"`
	IsActive bool   `json:"is_active"`
	Text     string `json:"text"`
}

func (tc *TaskControler) Init() error {
	error := tc.Store.InitDB()
	return error
}

func (tc *TaskControler) Insert(w http.ResponseWriter, r *http.Request) {
	if !assertMethod(w, r, http.MethodPost) {
		return
	}
	var data struct {
		Text string `json:"text"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil && err != io.EOF {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("found while decoding request body: %v", err)
		return
	}
	tc.Store.InsertTask(data.Text)
}
func (tc *TaskControler) GetAll(w http.ResponseWriter, r *http.Request) {
	if !assertMethod(w, r, http.MethodGet) {
		return
	}
	tasks := tc.Store.GetTasks()
	response := make([]ResponseTask, 0)
	for _, v := range tasks {
		response = append(response, ResponseTask{
			Id:       strconv.Itoa(int(v.ID)),
			IsActive: v.IsActive,
			Text:     v.Text,
		})
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil && err != io.EOF {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("found while encoding json to response writer %v", err)
		return
	}
}
func (tc *TaskControler) SetState(w http.ResponseWriter, r *http.Request) {
	if !assertMethod(w, r, http.MethodPut) {
		return
	}
	var data struct {
		Id       string `json:"id"`
		IsActive bool   `json:"is_active"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil && err != io.EOF {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("found while decoding json from  body on SetState %v", err)
		return
	}
	tc.Store.SetIsActive(data.Id, data.IsActive)
}

// if the request body has the string "completed" instead of an ID, then DeleteCompleted() will be run instead of Delete()
func (tc *TaskControler) Delete(w http.ResponseWriter, r *http.Request) {
	if !assertMethod(w, r, http.MethodDelete) {
		return
	}
	var data struct {
		Id string `json:"id"`
	}
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &data)
	if err != nil && err != io.EOF {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("found while decoding request body: %v", err)
		return
	}
	if data.Id == "completed" {
		tc.Store.DeleteCompleted()
	} else {
		tc.Store.Delete(data.Id)
	}
}

func assertMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		msg := fmt.Sprintf("Method not allowed, only %s allowed", method)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return false
	} else {
		return true
	}
}
