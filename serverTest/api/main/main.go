package main

import (
	"LeetCodeTasksGO/serverTest/api/v1/routing"
	"net/http"
	"strings"
)

func main() {
	go http.Handle("/", http.FileServer(http.Dir("./front")))
	go http.HandleFunc("/"+routing.Prefix, func(writer http.ResponseWriter, request *http.Request) {
		ctrl := strings.Split(request.URL.Path, routing.Prefix)
		routing.Delegate(ctrl[1], writer, request)
	})
	_ = http.ListenAndServe(":6054", nil)
}
