// Code generated by sqlc. DO NOT EDIT.
// source: testcase.sql

package db

import (
	"context"
)

const createTestCase = `-- name: CreateTestCase :one
INSERT INTO testcases (
  username,
  classname,
  filename,
  linenumber,
  testcasename,
  duration
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, classname, filename, linenumber, testcasename, duration, username
`

type CreateTestCaseParams struct {
	Username     string  `json:"username"`
	Classname    string  `json:"classname"`
	Filename     string  `json:"filename"`
	Linenumber   int32   `json:"linenumber"`
	Testcasename string  `json:"testcasename"`
	Duration     float64 `json:"duration"`
}

func (q *Queries) CreateTestCase(ctx context.Context, arg CreateTestCaseParams) (Testcase, error) {
	row := q.db.QueryRowContext(ctx, createTestCase,
		arg.Username,
		arg.Classname,
		arg.Filename,
		arg.Linenumber,
		arg.Testcasename,
		arg.Duration,
	)
	var i Testcase
	err := row.Scan(
		&i.ID,
		&i.Classname,
		&i.Filename,
		&i.Linenumber,
		&i.Testcasename,
		&i.Duration,
		&i.Username,
	)
	return i, err
}

const deleteTestCase = `-- name: DeleteTestCase :exec
DELETE FROM testcases
WHERE id = $1
`

func (q *Queries) DeleteTestCase(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTestCase, id)
	return err
}

const getTestCase = `-- name: GetTestCase :one
SELECT id, classname, filename, linenumber, testcasename, duration, username FROM testcases
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTestCase(ctx context.Context, id int64) (Testcase, error) {
	row := q.db.QueryRowContext(ctx, getTestCase, id)
	var i Testcase
	err := row.Scan(
		&i.ID,
		&i.Classname,
		&i.Filename,
		&i.Linenumber,
		&i.Testcasename,
		&i.Duration,
		&i.Username,
	)
	return i, err
}

const listTestCases = `-- name: ListTestCases :many
SELECT id, classname, filename, linenumber, testcasename, duration, username FROM testcases
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListTestCasesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTestCases(ctx context.Context, arg ListTestCasesParams) ([]Testcase, error) {
	rows, err := q.db.QueryContext(ctx, listTestCases, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Testcase{}
	for rows.Next() {
		var i Testcase
		if err := rows.Scan(
			&i.ID,
			&i.Classname,
			&i.Filename,
			&i.Linenumber,
			&i.Testcasename,
			&i.Duration,
			&i.Username,
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
