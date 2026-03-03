package api

import (
	"net/http"

	"github.com/OlegRozh/Final-progect-todo/internal/storage"
	"github.com/OlegRozh/Final-progect-todo/structs"
)

const tasksLimit = 50

type TasksResp struct {
	Tasks []structs.Task `json:"tasks"`
}

func GetListTasks(store storage.TaskStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
			return
		}
		tasks, err := store.GetListTasks(tasksLimit)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		if tasks == nil {
			tasks = []*structs.Task{}
		}
		response := struct {
			Tasks []*structs.Task `json:"tasks"`
		}{Tasks: tasks}
		writeJSON(w, http.StatusOK, response)
	}
}
