package cmd

import (
	"beginner_task_tracker/entity"
	"beginner_task_tracker/pkg"
	"fmt"
	"os"
	"strconv"
)

var tasks entity.Tasks

func init() {
	tasks = *pkg.LoadTasksFromJson()
}

func FireUp() {

	userArgs := os.Args[1:]

	switch mainCmd := userArgs[0]; mainCmd {
	case "help":
		fmt.Println("add [desc]\n" +
			"update [id] [desc]\n" +
			"delete [id]\n" +
			"mark-in-progress [id]\n" +
			"mark-done [id]\n" +
			"list\n" +
			"list done\n" +
			"list todo\n" +
			"list in-progress")

	case "add":
		defer func(t *entity.Tasks) {
			err := pkg.SaveTasksToJson(t)
			if err != nil {
				panic(err)
			}
		}(&tasks)

		tasks.AddTask(userArgs[1])
		break

	case "update":
		defer func(t *entity.Tasks) {
			err := pkg.SaveTasksToJson(t)
			if err != nil {
				panic(err)
			}
		}(&tasks)

		taskId, err := strconv.Atoi(userArgs[1])
		if err != nil {
			panic(err)
		}
		tasks.EditTask(taskId, userArgs[2])
		break

	case "delete":
		defer func(t *entity.Tasks) {
			err := pkg.SaveTasksToJson(t)
			if err != nil {
				panic(err)
			}
		}(&tasks)

		taskId, err := strconv.Atoi(userArgs[1])
		if err != nil {
			panic(err)
		}
		tasks.RemoveTask(taskId)
		break

	case "mark-in-progress":
		defer func(t *entity.Tasks) {
			err := pkg.SaveTasksToJson(t)
			if err != nil {
				panic(err)
			}
		}(&tasks)

		taskId, err := strconv.Atoi(userArgs[1])
		if err != nil {
			panic(err)
		}
		tasks.MarkTaskAs(taskId, "in-progress")
		break

	case "mark-done":
		defer func(t *entity.Tasks) {
			err := pkg.SaveTasksToJson(t)
			if err != nil {
				panic(err)
			}
		}(&tasks)

		taskId, err := strconv.Atoi(userArgs[1])
		if err != nil {
			panic(err)
		}
		tasks.MarkTaskAs(taskId, "done")
		break

	case "list":
		if len(userArgs) == 2 {
			status := userArgs[1]
			tasks.GetAllTasks(&status)
			break
		}

		tasks.GetAllTasks(nil)
		break
	default:
		fmt.Println("Unknown command, type help to see commands list")
		break
	}
}
