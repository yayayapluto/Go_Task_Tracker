package pkg_test

import (
	"beginner_task_tracker/entity"
	"beginner_task_tracker/helper"
	"beginner_task_tracker/pkg"
	"testing"
)

func TestLoadJsonFile(t *testing.T) {
	tasks := pkg.LoadTasksFromJson()
	helper.PrintAnyAsJson(tasks)
}

func TestSaveTasksToJson(t *testing.T) {
	tasks := entity.Tasks{Tasks: make([]entity.Task, 5)}
	err := pkg.SaveTasksToJson(&tasks)
	if err != nil {
		panic(err)
	}

	helper.PrintAnyAsJson(pkg.LoadTasksFromJson())
}
