package task

import (
	"github.com/GalichAnton/go_final_project/internal/models/task"
	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) GetById(id string) (*task.Task, error) {
	query := sq.
		Select("*").
		From(tableName).
		Where(sq.Eq{"id": id}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.DB().QueryRow(sql, args...)

	var t task.Task
	err = row.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
