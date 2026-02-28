package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/OlegRozh/Final-progect-todo/structs"
	_ "modernc.org/sqlite"
)

type SQLiteStorage struct {
	DB *sql.DB
}

const schema = `
CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT '',
    title VARCHAR(255) NOT NULL DEFAULT '',
    comment VARCHAR(255) NOT NULL DEFAULT '',
    repeat VARCHAR(128) DEFAULT ''
);
`

func New(dbPath string) (*SQLiteStorage, error) {
	_, err := os.Stat(dbPath)
	var install bool
	if os.IsNotExist(err) {
		install = true
	}
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	if install {
		_, err = db.Exec(schema)
		if err != nil {
			log.Fatal(err)
		}
	}
	storage := &SQLiteStorage{DB: db}
	return storage, nil
}

func (s *SQLiteStorage) GetTask(id string) (*structs.Task, error) {
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`
	var task structs.Task
	err := s.DB.QueryRow(query, id).Scan(
		&task.Id,
		&task.Date,
		&task.Title,
		&task.Comment,
		&task.Repeat,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("task with id %s not found", id)
		}
		return nil, err
	}
	return &task, nil
}

func (s *SQLiteStorage) GetListTasks(limit int) ([]structs.Task, error) {
	query := `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT ?`
	rows, err := s.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []structs.Task
	for rows.Next() {
		var task structs.Task
		err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *SQLiteStorage) CreateTask(date, title, comment, repeat string) (int64, error) {
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`
	result, err := s.DB.Exec(query, date, title, comment, repeat)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (s *SQLiteStorage) UpdateTask(task *structs.Task) error {
	query := `UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`
	result, err := s.DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.Id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("task doesn't exist")
	}
	return nil
}

func (s *SQLiteStorage) DeleteTask(id string) error {
	query := `DELETE FROM scheduler WHERE id = ?`
	result, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("task doesn't exist")
	}
	return nil
}

func (s *SQLiteStorage) UpdateTaskDate(id, newDate string) error {
	query := `UPDATE scheduler SET date = ? WHERE id = ?`
	result, err := s.DB.Exec(query, newDate, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("task doesn't exist")
	}
	return nil
}
