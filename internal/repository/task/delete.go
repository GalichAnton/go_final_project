package task

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) Delete(id string) error {
	query := sq.
		Delete(tableName).
		Where(sq.Eq{"id": id})

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
