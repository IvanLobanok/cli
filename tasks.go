package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type TaskManager struct {
	name  string
	tasks []Task
}

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *Task) SetID(a int) {
	t.Id = a
}
func (t *Task) SetDescription(a string) {
	t.Description = a
}
func (t *Task) SetStatusInProgress() {
	t.Status = false
}
func (t *Task) SetStatusDone() {
	t.Status = true
}
func (t *Task) SetCreatedAt(a time.Time) {
	t.CreatedAt = a
}
func (t *Task) SetUpdatedAt(a time.Time) {
	t.UpdatedAt = a
}

func (tm *TaskManager) GetNextID() int {
	maxID := 0
	for _, t := range tm.tasks {
		if t.Id > maxID {
			maxID = t.Id
		}
	}
	return (maxID + 1)

}

func (t *Task) StatusString() string {
	if t.Status {
		return "done"
	}
	return "in progress"
}

func (tm *TaskManager) Add(description string) {
	var t Task
	t.SetID(tm.GetNextID())
	t.SetDescription(description)
	t.SetStatusInProgress()
	t.SetCreatedAt(time.Now())
	t.SetUpdatedAt(time.Now())
	tm.tasks = append(tm.tasks, t)
}

func (tm *TaskManager) UpdateStatus(id int) error {
	for i := range tm.tasks {
		if tm.tasks[i].Id == id {
			tm.tasks[i].Status = !tm.tasks[i].Status
			tm.tasks[i].SetUpdatedAt(time.Now())
			return nil
		}
	}
	return errors.New("Task not found")
}

func (tm *TaskManager) Delete(id int) error {
	for i := range tm.tasks {
		if tm.tasks[i].Id == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			fmt.Println("DELETED")
			return nil
		}
	}
	return errors.New("Task not found")
}

func (tm *TaskManager) ViewAll() {
	for i := range tm.tasks {
		fmt.Println("===========================")
		fmt.Println("Task number", tm.tasks[i].Id)
		fmt.Println("Description:", tm.tasks[i].Description)
		fmt.Println("Status:", tm.tasks[i].StatusString())
		fmt.Println("Created at", tm.tasks[i].CreatedAt.Format(time.RFC822))
		fmt.Println("Updated at", tm.tasks[i].UpdatedAt.Format(time.RFC822))
		fmt.Println("===========================")
	}
}

func (tm *TaskManager) ViewAllDone() {
	for i := range tm.tasks {
		if tm.tasks[i].Status == true {
			fmt.Println("===========================")
			fmt.Println("Task number", tm.tasks[i].Id)
			fmt.Println("Description:", tm.tasks[i].Description)
			fmt.Println("Status:", tm.tasks[i].StatusString())
			fmt.Println("Created at", tm.tasks[i].CreatedAt.Format(time.RFC822))
			fmt.Println("Updated at", tm.tasks[i].UpdatedAt.Format(time.RFC822))
			fmt.Println("===========================")
		}
	}
}

func (tm *TaskManager) ViewAllInProgress() {
	for i := range tm.tasks {
		if tm.tasks[i].Status == false {
			fmt.Println("===========================")
			fmt.Println("Task number", tm.tasks[i].Id)
			fmt.Println("Description:", tm.tasks[i].Description)
			fmt.Println("Status:", tm.tasks[i].StatusString())
			fmt.Println("Created at", tm.tasks[i].CreatedAt.Format(time.RFC822))
			fmt.Println("Updated at", tm.tasks[i].UpdatedAt.Format(time.RFC822))
			fmt.Println("===========================")
		}
	}
}

func (tm *TaskManager) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(tm.tasks, "", "   ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (tm *TaskManager) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &tm.tasks)
}
