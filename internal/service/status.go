package service

import (
	"time"

	"github.com/OlegRozh/Final-progect-todo/structs"
)

func StatusDone(task *structs.Task, now time.Time) (bool, error) {
	if task.Repeat == "" {
		return true, nil
	}
	NextDate, err := NextDate(now, task.Date, task.Repeat)
	if err != nil {
		return false, err
	}
	task.Date = NextDate
	return false, nil
}
