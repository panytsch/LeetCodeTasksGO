package tasks

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var TaskPool = map[int]Task{
	1: {false, "First", "My First Task"},
}

func Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		uri := strings.Split(r.URL.Path, "/")
		statement := uri[len(uri)-1]
		if statement == "tasks" {
			task, err := getAll()
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.Header().Set("Content-Type", "application/json")
				js, err := json.Marshal(task)
				if err != nil {
					w.WriteHeader(http.StatusNotFound)
					break
				}
				_, _ = w.Write(js)
			}
		} else {
			id, err := strconv.ParseInt(statement, 10, 32)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				break
			}
			task, err := getById(int(id))
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				break
			}
			w.Header().Set("Content-Type", "application/json")
			js, err := json.Marshal(task)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				break
			}
			_, _ = w.Write(js)
		}
	case http.MethodPost:
		task := TaskPostRequest{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&task)
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			log.Print(err.Error())
		} else {
			TaskPool[task.Id] = task.Task
			w.WriteHeader(http.StatusCreated)
		}
	}
}

func getById(id int) (Task, error) {
	task, ok := TaskPool[id]
	if !ok {
		return task, errors.New("task not found")
	}
	return task, nil
}

func getAll() (map[int]Task, error) {
	if len(TaskPool) == 0 {
		return TaskPool, errors.New("u don't have tasks")
	}
	return TaskPool, nil
}
