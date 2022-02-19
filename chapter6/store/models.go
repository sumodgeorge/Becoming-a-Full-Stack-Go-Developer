// Code generated by sqlc. DO NOT EDIT.

package store

import (
	"encoding/json"
	"time"
)

type GowebappExercise struct {
	UserID       int64  `db:"user_id"`
	ExerciseName string `db:"exercise_name"`
}

type GowebappImage struct {
	ImageID     int64  `db:"image_id"`
	UserID      int64  `db:"user_id"`
	ContentType string `db:"content_type"`
	ImageData   []byte `db:"image_data"`
}

type GowebappSet struct {
	SetID        int64  `db:"set_id"`
	WorkoutID    int64  `db:"workout_id"`
	ExerciseName string `db:"exercise_name"`
	Weight       int32  `db:"weight"`
	Set1         int64  `db:"set1"`
	Set2         int64  `db:"set2"`
	Set3         int64  `db:"set3"`
}

type GowebappUser struct {
	UserID       int64           `db:"user_id"`
	UserName     string          `db:"user_name"`
	PasswordHash string          `db:"password_hash"`
	Name         string          `db:"name"`
	Config       json.RawMessage `db:"config"`
	CreatedAt    time.Time       `db:"created_at"`
	IsEnabled    bool            `db:"is_enabled"`
}

type GowebappWorkout struct {
	WorkoutID int64     `db:"workout_id"`
	UserID    int64     `db:"user_id"`
	StartDate time.Time `db:"start_date"`
}
