package service

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/OlegRozh/Final-progect-todo/structs"
)

const dateLayout = "20060102"

func afterNow(date, now time.Time) bool {
	dateUTC := date.UTC().Truncate(24 * time.Hour)
	nowUTC := now.UTC().Truncate(24 * time.Hour)
	return dateUTC.After(nowUTC)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("repeat can't be empty")
	}
	date, err := time.Parse(dateLayout, dstart)
	if err != nil {
		return "", err
	}
	parts := strings.Split(repeat, " ")
	if len(parts) > 2 {
		return "", errors.New("invalid repeat format")
	}
	rule := parts[0]
	switch rule {
	case "d":
		if len(parts) < 2 {
			return "", errors.New("for d rule need number of days")
		}
		interval, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", errors.New("days interval must be an integer")
		}
		if interval <= 0 || interval > 400 {
			return "", errors.New("days must be between 1 and 400")
		}
		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				break
			}
		}
	case "y":
		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
	default:
		return "", errors.New("invalid rule")
	}
	return date.UTC().Format(dateLayout), nil
}

func CheckDate(task *structs.Task) error {
	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format(dateLayout)
		return nil
	}
	t, err := time.Parse(dateLayout, task.Date)
	if err != nil {
		return err
	}
	var next string
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}
	if afterNow(now, t) {
		if task.Repeat == "" {
			task.Date = now.Format(dateLayout)
		} else {
			task.Date = next
		}
	}
	return nil
}
