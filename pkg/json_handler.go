package pkg

import (
	"beginner_task_tracker/entity"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
)

var filePath string

func init() {
	filePath = filepath.Join("./", "tasks.json")
}

func LoadTasksFromJson() *entity.Tasks {
	jf, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	defer func(jf *os.File) {
		err := jf.Close()
		if err != nil {
			panic(err)
		}
	}(jf)

	var tasks entity.Tasks

	if stat, _ := jf.Stat(); stat.Size() <= 0 {
		return &tasks
	}

	err = json.NewDecoder(jf).Decode(&tasks)
	if err != nil {
		panic(err)
	}

	return &tasks
}

func SaveTasksToJson(t *entity.Tasks) error {
	jf, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, fs.ModePerm)
	if err != nil {
		return err
	}

	defer func(jf *os.File) {
		err := jf.Close()
		if err != nil {
			panic(err)
		}
	}(jf)

	bytesJson, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		return err
	}

	_, err = jf.Write(bytesJson)
	if err != nil {
		return err
	}

	return nil
}
