package entity_test

import (
	"beginner_task_tracker/entity"
	"testing"
	"time"
)

var Tasks entity.Tasks

func init() {
	Tasks = entity.Tasks{Tasks: []entity.Task{
		{
			Id:        1,
			Desc:      "Go to gym",
			Status:    "todo",
			CreatedAt: time.Now().Truncate(time.Hour * 5),
			UpdatedAt: time.Now(),
		}, {
			Id:        2,
			Desc:      "Go to gym again",
			Status:    "todo",
			CreatedAt: time.Now().Truncate(time.Hour * 4),
			UpdatedAt: time.Now(),
		}, {
			Id:        3,
			Desc:      "Go to gym again and again",
			Status:    "todo",
			CreatedAt: time.Now().Truncate(time.Hour * 3),
			UpdatedAt: time.Now(),
		},
	}}
}

func TestMain(t *testing.M) {
	t.Run()
}

func TestAddingTask(t *testing.T) {
	t.Run("Add task with valid desc", func(t *testing.T) {
		Tasks.AddTask("Go to supermarket to buy milks")
	})

	t.Run("Add task with existed desc", func(t *testing.T) {
		Tasks.AddTask("Go to gym again and again")
	})
}

func TestEditTask(t *testing.T) {
	t.Run("Edit task with valid id and valid desc", func(t *testing.T) {
		Tasks.EditTask(1, "Go to barber")
	})

	t.Run("Edit task with valid id and existed desc", func(t *testing.T) {
		Tasks.EditTask(1, "Go to supermarket to buy milks")
	})

	t.Run("Edit task with invalid id and valid desc", func(t *testing.T) {
		Tasks.EditTask(18, "Go to some store")
	})

	t.Run("Edit task with invalid id and existed desc", func(t *testing.T) {
		Tasks.EditTask(98, "Go to gym again and again")
	})
}

func TestRemoveTask(t *testing.T) {
	t.Run("Remove task with valid id", func(t *testing.T) {
		Tasks.RemoveTask(3)
	})

	t.Run("Remove task with invalid id", func(t *testing.T) {
		Tasks.RemoveTask(100)
	})
}

func TestMarkTaskAs(t *testing.T) {
	t.Run("Mark task with valid id and valid status", func(t *testing.T) {
		Tasks.MarkTaskAs(1, "done")
	})

	t.Run("Mark task with invalid id and valid status", func(t *testing.T) {
		Tasks.MarkTaskAs(95, "done")
	})

	t.Run("Mark task with valid id and invalid status", func(t *testing.T) {
		Tasks.MarkTaskAs(1, "not-yet")
	})

	t.Run("Mark task with invalid id and invalid status", func(t *testing.T) {
		Tasks.MarkTaskAs(100, "not-yet")
	})
}
