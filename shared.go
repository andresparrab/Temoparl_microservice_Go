package app

import "time"

const NumberTaskQueue = "UPDATE_NUMBER_TASK_QUEUE"

type Post struct {
	ID          int64     `db:"id"`
	TestNumbers int64     `json:"test_numbers" db:"test_numbers" `
	CreatedAt   time.Time `db:"created_at"`
}
