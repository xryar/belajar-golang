package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := new(CustomMux)

	mux.HandleFunc("/student", ActionStudent)

	mux.RegisterMiddleware(MiddlewareAuth)
	mux.RegisterMiddleware(MiddlewareAllowOnlyGet)

	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = mux

	fmt.Println("server started at localhost:9000")
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, selectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(res)
}
