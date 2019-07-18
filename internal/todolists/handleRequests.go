package todolists

import (
	"encoding/json"
	"golang-todo-list/internal/util"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// HandleRequests handles the GET requests
func HandleRequests(w http.ResponseWriter, r *http.Request) {
	var (
		todolists []*TodoList
		err       error
	)

	vars := mux.Vars(r)
	listID, err := strconv.Atoi(vars["listID"])

	util.Check(err)

	switch r.Method {
	case "GET":
		if listID > 0 {
			todolists, err = GetTodoList(int(listID))
		} else {
			todolists, err = GetTodoLists()
		}
	default:
		todolists, err = GetTodoLists()
	}

	util.Check(err)

	json.NewEncoder(w).Encode(todolists)
}
