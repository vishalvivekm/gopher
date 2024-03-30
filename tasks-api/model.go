package main

import (
	"errors"
)

func CreateInitialTasks() ([]Task, int) {
	tasks := []Task{
		{ID: 1, Name: "Create project proposal", Description: "Write a proposal for the new project", DueDate: "2022-02-01"},
		{ID: 2, Name: "Design website layout", Description: "Create a layout for the company website", DueDate: "2022-03-01"},
		{ID: 3, Name: "Implement payment system", Description: "Integrate a payment system into the website", DueDate: "2022-04-01"},
		{ID: 4, Name: "Conduct user testing", Description: "Gather feedback from user testing to improve the website", DueDate: "2022-05-01"},
		{ID: 5, Name: "Launch website", Description: "Make the website live for the public", DueDate: "2022-06-01"},
		{ID: 6, Name: "Evaluate website performance", Description: "Collect data and analyze website performance", DueDate: "2022-07-01"},
	}
	return tasks, 6
}

func getTasks() ([]Task, error) {

	return tasks, nil
}

func (t *Task) createTask() error {
	currentID = currentID + 1
	t.ID = currentID
	tasks = append(tasks, *t)
	return nil
}

func (t *Task) getTask() error {
	for index, val := range tasks {
		if val.ID == t.ID {
			*t = tasks[index]
			return nil
		}
	}
	//return fmt.Errorf("no taks with id, %v", t.ID)
	return errors.New("task not found")
}

func (t *Task) updateTask() error {
	for index, val := range tasks {
		if val.ID == t.ID {
			tasks[index] = *t
			return nil
		}
	}
	return errors.New("task not found")
}
func (t *Task) deleteTask() error {
	for index, val := range tasks {
		if val.ID == t.ID {
			tasks[index] = Task{}
			return nil
		}
	}
	//return fmt.Errorf("no task with id, %v", t.ID)
	return errors.New("task not found")
}
