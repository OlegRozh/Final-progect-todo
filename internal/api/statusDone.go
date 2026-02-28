package api

import (
	"net/http"
	"time"

	"github.com/OlegRozh/Final-progect-todo/internal/service"
	"github.com/OlegRozh/Final-progect-todo/internal/storage"
)

func StatusDoneHandler(store storage.TaskStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
			return
		}
		id := r.URL.Query().Get("id")
		if id == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id is required"})
			return
		}
		task, err := store.GetTask(id)
		if err != nil {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "Task not found"})
			return
		}
		shouldDelete, err := service.StatusDone(task, time.Now())
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		if shouldDelete {
			err = store.DeleteTask(id)
		} else {
			err = store.UpdateTask(task)
		}
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{})
	}
}
