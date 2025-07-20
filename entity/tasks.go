package entity

import (
	"beginner_task_tracker/helper"
	"fmt"
	"slices"
	"strings"
	"time"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

func (t *Tasks) GetAllTasks(status *string) {
	var taskCollection []Task

	if status != nil {
		*status = strings.ToLower(strings.TrimSpace(*status))
		allowedStatus := []string{"todo", "in-progress", "done"}
		if !slices.Contains(allowedStatus, *status) {
			fmt.Println("[Failed] Invalid status (valid status such as: todo, in-progress, done)")
			return
		}

		for _, task := range t.Tasks {
			if task.Status == *status {
				taskCollection = append(taskCollection, task)
			}
		}

		if len(taskCollection) <= 0 {
			fmt.Printf("[Not Found] Cannot find tasks with status: %v\n", status)
			return
		}

		fmt.Printf("[Done] Successfully find tasks with status: %v\n", status)
		helper.PrintAnyAsJson(taskCollection)
		return
	}

	taskCollection = t.Tasks
	fmt.Println("[Done] Successfully getting all tasks")
	helper.PrintAnyAsJson(taskCollection)
	return
}

func (t *Tasks) AddTask(desc string) {
	desc = strings.TrimSpace(desc)

	for _, task := range t.Tasks {
		if task.Desc == desc {
			fmt.Println("[Failed] Task already exists")
			return
		}
	}

	newTask := Task{
		Id:        len(t.Tasks) + 1,
		Desc:      desc,
		Status:    "todo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	t.Tasks = append(t.Tasks, newTask)

	fmt.Printf("[Done] Sucessfully adding new task (ID: %v)\n", newTask.Id)
	return
}

func (t *Tasks) EditTask(id int, newDesc string) {
	newDesc = strings.TrimSpace(newDesc)

	for i, task := range t.Tasks {

		if task.Id == id {
			for _, t2 := range t.Tasks {
				if t2.Id == task.Id {
					continue
				}

				if t2.Desc == newDesc {
					fmt.Println("[Failed] Task already exists")
					return
				}
			}

			t.Tasks[i].Desc = newDesc
			t.Tasks[i].UpdatedAt = time.Now()
			fmt.Printf("[Done] Successfully edit task (ID: %v) (New Desc: %v)\n", task.Id, newDesc)
			return
		}

	}

	fmt.Printf("[Failed] Cannot find task with id: %v\n", id)
	return
}

func (t *Tasks) RemoveTask(id int) {
	idx := -1
	
	for i, task := range t.Tasks {
		if task.Id == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Printf("[Failed] Cannot find task with id: %v\n", id)
		return
	}

	t.Tasks = append(t.Tasks[:idx], t.Tasks[idx+1:]...)
	fmt.Println("[Done] Successfully removing task")
	return
}

func (t *Tasks) MarkTaskAs(id int, status string) {
	status = strings.ToLower(strings.TrimSpace(status))

	allowedStatus := []string{"todo", "in-progress", "done"}
	if !slices.Contains(allowedStatus, status) {
		fmt.Println("[Failed] Invalid status (valid status such as: todo, in-progress, done)")
		return
	}

	for i, task := range t.Tasks {
		if task.Id == id {
			t.Tasks[i].Status = status
			t.Tasks[i].UpdatedAt = time.Now()
			fmt.Println("[Done] Successfully marking task as", status)
			return
		}
	}

	fmt.Printf("[Failed] Cannot find task with id: %v\n", id)
	return
}
