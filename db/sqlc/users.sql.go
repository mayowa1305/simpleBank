// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: users.sql

package db

import (
	"context"
)

const createuser = `-- name: Createuser :one
INSERT INTO "users" (
  "username",
  "hashed_password",
  "full_name",
  "email"
) VALUES (
  $1, $2, $3, $4
)
RETURNING username, hashed_password, full_name, email, password_changed_at, created_at
`

type CreateuserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashedPassword"`
	FullName       string `json:"fullName"`
	Email          string `json:"email"`
}

func (q *Queries) Createuser(ctx context.Context, arg CreateuserParams) (User, error) {
	row := q.queryRow(ctx, q.createuserStmt, createuser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getuser = `-- name: Getuser :one
SELECT username, hashed_password, full_name, email, password_changed_at, created_at FROM "users"
WHERE "username" = $1 LIMIT 1
`

func (q *Queries) Getuser(ctx context.Context, username string) (User, error) {
	row := q.queryRow(ctx, q.getuserStmt, getuser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}