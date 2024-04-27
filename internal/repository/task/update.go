package task

import (
	"fmt"

	"github.com/GalichAnton/go_final_project/internal/models/task"
	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) Update(task *task.Task) error {
	query := sq.
		Update(tableName).
		Set("date", task.Date).
		Set("title", task.Title).
		Set("comment", task.Comment).
		Set("repeat", task.Repeat).
		Where(sq.Eq{"id": task.ID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	res, err := r.db.DB().Exec(sql, args...)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("задача не найдена")
	}

	return nil
}
