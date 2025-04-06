package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Goal represents a user's goal
type Goal struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CurrentLevel string    `json:"current_level"`
	TargetLevel  string    `json:"target_level"`
	StartDate    time.Time `json:"start_date"`
	TargetDate   time.Time `json:"target_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Timeline represents a timeline for a goal
type Timeline struct {
	ID          string    `json:"id"`
	GoalID      string    `json:"goal_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Tasks       []TimelineTask `json:"tasks,omitempty"`
}

// TimelineTask represents a task in a timeline
type TimelineTask struct {
	ID          string    `json:"id"`
	TimelineID  string    `json:"timeline_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Duration    string    `json:"duration"`
	Priority    int       `json:"priority"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TimelineInput represents input for generating a timeline
type TimelineInput struct {
	CurrentLevel string `json:"current_level"`
	Goal         string `json:"goal"`
	Objectives   string `json:"objectives"`
	CurrentDate  string `json:"current_date"`
	TargetDate   string `json:"target_date,omitempty"`
}