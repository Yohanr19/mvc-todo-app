package psqlStore

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)
type TaskStore struct {
	DB *gorm.DB
}
func (ts *TaskStore)InitDB()error{
	connStr := `host=localhost user=yohan password=yohan1234 dbname=mydb`
	db , err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("found while connecting to db: %v",err)
	}
	err = db.AutoMigrate(&Task{})
	if err != nil {
		return fmt.Errorf("found while migrating table %v",err)
	}
	ts.DB = db
	return nil
}

func (ts *TaskStore)InsertTask(text string){
	var task Task 
	task.Text = text
	task.IsActive = true
	ts.DB.Create(&task)
}