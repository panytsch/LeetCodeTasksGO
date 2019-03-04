package routing

import (
	"LeetCodeTasksGO/serverTest/api/v1/controllers/tasks"
	"net/http"
	"strings"
)

type foo func(w http.ResponseWriter, r *http.Request)

const Prefix = "api/v1/"

var Routes = map[string]foo{
	"tasks": func(w http.ResponseWriter, r *http.Request) {
		tasks.Handle(w, r)
	},
}

func Delegate(path string, w http.ResponseWriter, r *http.Request) {
	ctrl := strings.Split(path, "/")
	route, ok := Routes[ctrl[0]]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	route(w, r)
}
