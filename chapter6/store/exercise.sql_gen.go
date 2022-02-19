// Code generated by sqlc. DO NOT EDIT.
// source: exercise.sql

package store

import (
	"context"
)

const createDefaultSetForExercise = `-- name: CreateDefaultSetForExercise :one
INSERT INTO gowebapp.sets (
  	Workout_ID,
    Exercise_Name, 
    Weight      
) VALUES (
    $1,
    $2,
    $3
) RETURNING set_id, workout_id, exercise_name, weight, set1, set2, set3
`

type CreateDefaultSetForExerciseParams struct {
	WorkoutID    int64  `db:"workout_id"`
	ExerciseName string `db:"exercise_name"`
	Weight       int32  `db:"weight"`
}

func (q *Queries) CreateDefaultSetForExercise(ctx context.Context, arg CreateDefaultSetForExerciseParams) (GowebappSet, error) {
	row := q.db.QueryRowContext(ctx, createDefaultSetForExercise, arg.WorkoutID, arg.ExerciseName, arg.Weight)
	var i GowebappSet
	err := row.Scan(
		&i.SetID,
		&i.WorkoutID,
		&i.ExerciseName,
		&i.Weight,
		&i.Set1,
		&i.Set2,
		&i.Set3,
	)
	return i, err
}

const createSetForExercise = `-- name: CreateSetForExercise :one
INSERT INTO gowebapp.sets (
  	Workout_ID,
    Exercise_Name, 
    Weight,
    Set1,
    Set2,
    Set3
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING set_id, workout_id, exercise_name, weight, set1, set2, set3
`

type CreateSetForExerciseParams struct {
	WorkoutID    int64  `db:"workout_id"`
	ExerciseName string `db:"exercise_name"`
	Weight       int32  `db:"weight"`
	Set1         int64  `db:"set1"`
	Set2         int64  `db:"set2"`
	Set3         int64  `db:"set3"`
}

func (q *Queries) CreateSetForExercise(ctx context.Context, arg CreateSetForExerciseParams) (GowebappSet, error) {
	row := q.db.QueryRowContext(ctx, createSetForExercise,
		arg.WorkoutID,
		arg.ExerciseName,
		arg.Weight,
		arg.Set1,
		arg.Set2,
		arg.Set3,
	)
	var i GowebappSet
	err := row.Scan(
		&i.SetID,
		&i.WorkoutID,
		&i.ExerciseName,
		&i.Weight,
		&i.Set1,
		&i.Set2,
		&i.Set3,
	)
	return i, err
}

const createUserDefaultExercise = `-- name: CreateUserDefaultExercise :exec
INSERT INTO gowebapp.exercises (
    User_ID,
    Exercise_Name
) VALUES (
    1,
    'Bench Press'
),(
    1,
    'Barbell row'
)
`

func (q *Queries) CreateUserDefaultExercise(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createUserDefaultExercise)
	return err
}

const createUserExercise = `-- name: CreateUserExercise :one
INSERT INTO gowebapp.exercises (
    User_ID,
    Exercise_Name
) VALUES (
    $1,
    $2
) ON CONFLICT (Exercise_Name) DO NOTHING RETURNING (
    User_ID, Exercise_Name
)
`

type CreateUserExerciseParams struct {
	UserID       int64  `db:"user_id"`
	ExerciseName string `db:"exercise_name"`
}

func (q *Queries) CreateUserExercise(ctx context.Context, arg CreateUserExerciseParams) (interface{}, error) {
	row := q.db.QueryRowContext(ctx, createUserExercise, arg.UserID, arg.ExerciseName)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const createUserWorkout = `-- name: CreateUserWorkout :one
INSERT INTO gowebapp.workouts (
    User_ID,
    Start_Date
) VALUES (
    $1,
    NOW()
) RETURNING workout_id, user_id, start_date
`

func (q *Queries) CreateUserWorkout(ctx context.Context, userID int64) (GowebappWorkout, error) {
	row := q.db.QueryRowContext(ctx, createUserWorkout, userID)
	var i GowebappWorkout
	err := row.Scan(&i.WorkoutID, &i.UserID, &i.StartDate)
	return i, err
}

const deleteUserExercise = `-- name: DeleteUserExercise :exec
DELETE FROM gowebapp.exercises
WHERE User_ID = $1 AND Exercise_Name = $2
`

type DeleteUserExerciseParams struct {
	UserID       int64  `db:"user_id"`
	ExerciseName string `db:"exercise_name"`
}

func (q *Queries) DeleteUserExercise(ctx context.Context, arg DeleteUserExerciseParams) error {
	_, err := q.db.ExecContext(ctx, deleteUserExercise, arg.UserID, arg.ExerciseName)
	return err
}

const listUserExercises = `-- name: ListUserExercises :many
SELECT Exercise_Name
FROM gowebapp.exercises
WHERE User_ID = $1
`

func (q *Queries) ListUserExercises(ctx context.Context, userID int64) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listUserExercises, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var exercise_name string
		if err := rows.Scan(&exercise_name); err != nil {
			return nil, err
		}
		items = append(items, exercise_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSet = `-- name: UpdateSet :one
UPDATE gowebapp.sets SET
    Weight = $2,
    Set1 = $3,
    Set2 = $4,
    Set3 = $5
WHERE Set_ID = $1 RETURNING set_id, workout_id, exercise_name, weight, set1, set2, set3
`

type UpdateSetParams struct {
	SetID  int64 `db:"set_id"`
	Weight int32 `db:"weight"`
	Set1   int64 `db:"set1"`
	Set2   int64 `db:"set2"`
	Set3   int64 `db:"set3"`
}

func (q *Queries) UpdateSet(ctx context.Context, arg UpdateSetParams) (GowebappSet, error) {
	row := q.db.QueryRowContext(ctx, updateSet,
		arg.SetID,
		arg.Weight,
		arg.Set1,
		arg.Set2,
		arg.Set3,
	)
	var i GowebappSet
	err := row.Scan(
		&i.SetID,
		&i.WorkoutID,
		&i.ExerciseName,
		&i.Weight,
		&i.Set1,
		&i.Set2,
		&i.Set3,
	)
	return i, err
}
