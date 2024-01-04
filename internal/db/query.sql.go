// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createCategory = `-- name: CreateCategory :exec
insert into  categories (id, name, description)
values (?,?,?)
`

type CreateCategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, createCategory, arg.ID, arg.Name, arg.Description)
	return err
}

const createCourse = `-- name: CreateCourse :exec
insert into courses (id, name, description, category_id, price)
values (?,?,?,?,?)
`

type CreateCourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	CategoryID  string
	Price       float64
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) error {
	_, err := q.db.ExecContext(ctx, createCourse,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.CategoryID,
		arg.Price,
	)
	return err
}

const deleteCategory = `-- name: DeleteCategory :exec
delete from categories
where id = ?
`

func (q *Queries) DeleteCategory(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const getCategory = `-- name: GetCategory :one
select id, name, description from categories
where id = ?
`

func (q *Queries) GetCategory(ctx context.Context, id string) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const listCategories = `-- name: ListCategories :many
select id, name, description from categories
`

func (q *Queries) ListCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCourses = `-- name: ListCourses :many
select co.id, co.category_id, co.name, co.description, co.price, ca.name as category_name
from courses co join categories ca on co.category_id = ca.id
`

type ListCoursesRow struct {
	ID           string
	CategoryID   string
	Name         string
	Description  sql.NullString
	Price        float64
	CategoryName string
}

func (q *Queries) ListCourses(ctx context.Context) ([]ListCoursesRow, error) {
	rows, err := q.db.QueryContext(ctx, listCourses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListCoursesRow
	for rows.Next() {
		var i ListCoursesRow
		if err := rows.Scan(
			&i.ID,
			&i.CategoryID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.CategoryName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :exec
update categories
set name = ?, description = ?
where id = ?
`

type UpdateCategoryParams struct {
	Name        string
	Description sql.NullString
	ID          string
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCategory, arg.Name, arg.Description, arg.ID)
	return err
}
