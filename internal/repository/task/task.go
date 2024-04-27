package task

import (
	"github.com/GalichAnton/go_final_project/internal/clients/db"
)

const (
	tableName  = "scheduler"
	colTitle   = "title"
	colDate    = "date"
	colComment = "comment"
	colRepeat  = "repeat"
)

type Repository struct {
	db db.Client
}

func NewTaskRepository(db db.Client) *Repository {
	return &Repository{db: db}
}
