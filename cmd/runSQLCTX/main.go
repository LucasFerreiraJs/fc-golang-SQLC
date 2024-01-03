package main

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lucasferreirajs/17-SQLC/internal/db"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

type CourseParams struct {
	ID string
	Name string
	Description sql.NullString
}

type CategoryParams struct {
	ID string
	Name string
	Description sql.NullString
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn: dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {

	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err !=nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)

	if err != nil {

		errRb := tx.Rollback()
		if errRb != nil {
			return fmt.Errorf("Error on Rollback: %v, original error: %w", errRb, err)
		}
		return err
	}

	return tx.Commit()
}


func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx,func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams {
			ID: argsCategory.ID,
			Name: argsCategory.Name,
			Description: argsCategory.Description,
		})

		if err !=nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams {
			ID: argsCourse.ID,
			Name: argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID: argsCategory.ID,
		})

		if err !=nil {
			return err
		}

		return nil
	})

	if err !=nil {
		return  err
	}

	return nil
}

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	queries := db.New(dbConn)


}
