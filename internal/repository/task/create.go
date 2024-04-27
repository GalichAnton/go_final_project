package task

import (
	"github.com/GalichAnton/go_final_project/internal/models/task"
	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) Create(info *task.Info) (int64, error) {
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(colDate, colTitle, colComment, colRepeat).
		Values(info.Date, info.Title, info.Comment, info.Repeat).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	stmt, err := r.db.DB().Prepare(query)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
