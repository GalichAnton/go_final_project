package task

import (
	"time"

	"github.com/GalichAnton/go_final_project/internal/models/task"
	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) GetTasks(searchStr string) ([]task.Task, error) {
	builderGet := sq.
		Select("*").
		From(tableName)

	searchDate, err := time.Parse("02.01.2006", searchStr)
	if err == nil {
		builderGet = builderGet.Where(sq.Eq{"date": searchDate.Format("20060102")})
	} else {
		builderGet = builderGet.Where(
			sq.Or{
				sq.Like{"title": "%" + searchStr + "%"},
				sq.Like{"comment": "%" + searchStr + "%"},
			},
		)
	}

	builderGet = builderGet.OrderBy("date ASC").
		Limit(10)

	query, args, err := builderGet.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.DB().Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []task.Task{}
	for rows.Next() {
		var t task.Task
		err = rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
