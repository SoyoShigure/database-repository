package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/soyoshigure/database-repository/option"
)

type RepositoryBase struct {
}

func (repo *RepositoryBase) Select(tx *sql.Tx, ctx context.Context, opt *option.SQLSelectOption) (*sql.Rows, error) {
	sql := "SELECT "
	if opt.Columns != nil {
		for i, col := range *opt.Columns {
			if i != 0 {
				sql += ", "
			}
			sql += col
		}
		sql += fmt.Sprintf(" FROM %s", opt.Table)
	} else {
		sql += fmt.Sprintf("* FROM %s", opt.Table)
	}

	if opt.WherePhrase != nil {
		where, err := opt.WherePhrase.ToSQL()
		if err != nil {
			return nil, err
		}
		sql += fmt.Sprintf(" WHERE %s", where)
	}

	if opt.OrderBy != nil {
		if opt.OrderBy.IsASC {
			sql += fmt.Sprintf(" ORDER BY %s ASC", opt.OrderBy.Column)
		} else {
			sql += fmt.Sprintf(" ORDER BY %s DESC", opt.OrderBy.Column)
		}
	}

	if opt.Limit != 0 {
		sql += fmt.Sprintf(" LIMIT %d", opt.Limit)
	}

	if opt.Offset != nil {
		sql += fmt.Sprintf(" OFFSET %d", opt.Offset)
	}

	return tx.QueryContext(ctx, sql)
}

func (repo *RepositoryBase) SelectRow(tx *sql.Tx, ctx context.Context, opt *option.SQLSelectOption) (*sql.Row, error) {
	sql := "SELECT "
	if opt.Columns != nil {
		for i, col := range *opt.Columns {
			if i != 0 {
				sql += ", "
			}
			sql += col
		}
		sql += fmt.Sprintf(" FROM %s", opt.Table)
	} else {
		sql += fmt.Sprintf("* FROM %s", opt.Table)
	}

	if opt.WherePhrase != nil {
		where, _ := opt.WherePhrase.ToSQL()
		sql += fmt.Sprintf(" WHERE %s", where)
	}

	if opt.OrderBy != nil {
		if opt.OrderBy.IsASC {
			sql += fmt.Sprintf(" ORDER BY %s ASC", opt.OrderBy.Column)
		} else {
			sql += fmt.Sprintf(" ORDER BY %s DESC", opt.OrderBy.Column)
		}
	}

	sql += "LIMIT 1"

	if opt.Offset != nil {
		sql += fmt.Sprintf(" OFFSET %d", opt.Offset)
	}

	return tx.QueryRowContext(ctx, sql), nil
}

func (repo *RepositoryBase) Insert(tx *sql.Tx, ctx context.Context, opt *option.SQLInsertOption) error {

	sql := fmt.Sprintf("INSERT INTO %s (", opt.Table)

	sqlValues := ""

	for i, col := range *opt.Columns {
		if i != 0 {
			sql += fmt.Sprintf(", %s", col)
			sqlValues += ", ?"
		} else {
			sql += fmt.Sprintf("%s", col)
			sqlValues += "?"
		}
	}

	sql += fmt.Sprintf(") VALUES (%s)", sqlValues)

	_, err := tx.ExecContext(ctx, sql, *opt.Values...)

	return err
}

func (repo *RepositoryBase) Update(tx *sql.Tx, ctx context.Context, opt *option.SQLUpdateOption) error {
	sql := fmt.Sprintf("UPDATE %s SET ", opt.Table)

	for i, col := range *opt.Columns {
		if i != 0 {
			sql += fmt.Sprintf(", %s = ?", col)
		} else {
			sql += fmt.Sprintf("%s = ?", col)
		}
	}

	if opt.WherePhrase != nil {
		where, err := opt.WherePhrase.ToSQL()

		if err != nil {
			return err
		}

		sql += fmt.Sprintf(" WHERE %s", where)
	}

	_, err := tx.ExecContext(ctx, sql, *opt.Values...)

	return err
}

func (repo *RepositoryBase) Delete(tx *sql.Tx, ctx context.Context, opt *option.SQLDeleteOption) error {
	sql := fmt.Sprintf("DELETE FROM %s", opt.Table)

	where, err := opt.WherePhrase.ToSQL()

	if err != nil {
		return err
	}

	sql += fmt.Sprintf(" WHERE %s", where)

	_, err = tx.ExecContext(ctx, sql)

	return nil

}
