package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jukemori/timeline-generator/internal/database"
	"github.com/jukemori/timeline-generator/internal/models"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting seed process")

	if err := run(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	logrus.Info("Seed completed successfully")
	os.Exit(0)
}

func run() error {
	ctx := context.Background()
	
	// Initialize database connection
	database.InitDB()
	db := database.DB
	
	// Begin transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	// Delete existing data first
	if err := deleteAll(ctx, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete existing data: %v", err)
	}

	// Create seed data
	if err := createAll(ctx, tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create seed data: %v", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func deleteAll(ctx context.Context, tx *sql.Tx) error {
	// Delete in order of dependencies
	_, err := tx.ExecContext(ctx, "DELETE FROM timeline_tasks")
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM timelines")
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM goals")
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM users")
	if err != nil {
		return err
	}

	return nil
}

func createAll(ctx context.Context, tx *sql.Tx) error {
	// Sample IDs
	userID := "1"
	logrus.Infof("Generated userID: %s", userID)
	
	// Create a user first
	_, err := tx.ExecContext(ctx,
		"INSERT INTO users (id, email, created_at, updated_at) VALUES (?, ?, ?, ?)",
		userID, "test@example.com", time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}
	
	// Create a goal with the user ID
	goalID := uuid.New().String()
	_, err = tx.ExecContext(ctx,
		"INSERT INTO goals (id, user_id, title, description, current_level, target_level, start_date, target_date, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		goalID, userID, "Master Go Programming", "Become proficient in Go", "Beginner", "Advanced", 
		time.Now(), time.Now().AddDate(0, 3, 0), time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("failed to insert goal: %v", err)
	}

	// Create sample timelines
	timeline1 := models.Timeline{
		ID:          uuid.New().String(),
		GoalID:      goalID,
		Title:       "Learning Go Programming",
		Description: "A complete pathway to master Go programming language",
		StartDate:   time.Now(),
		EndDate:     time.Now().AddDate(0, 3, 0), // 3 months from now
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Insert timeline
	_, err = tx.ExecContext(ctx, 
		"INSERT INTO timelines (id, goal_id, title, description, start_date, end_date, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		timeline1.ID, timeline1.GoalID, timeline1.Title, timeline1.Description, 
		timeline1.StartDate, timeline1.EndDate, timeline1.CreatedAt, timeline1.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert timeline: %v", err)
	}

	// Create sample tasks
	tasks := []models.TimelineTask{
		{
			ID:          uuid.New().String(),
			TimelineID:  timeline1.ID,
			Title:       "Setup Go Environment",
			Description: "Install Go and set up the development environment",
			StartDate:   timeline1.StartDate,
			EndDate:     timeline1.StartDate.AddDate(0, 0, 7),
			Duration:    "7 days",
			Priority:    1,
			Completed:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New().String(),
			TimelineID:  timeline1.ID,
			Title:       "Learn Go Basics",
			Description: "Learn basic syntax, variables, and control structures",
			StartDate:   timeline1.StartDate.AddDate(0, 0, 8),
			EndDate:     timeline1.StartDate.AddDate(0, 0, 21),
			Duration:    "14 days",
			Priority:    1,
			Completed:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		// Add more tasks as needed
	}

	// Insert tasks
	for _, task := range tasks {
		_, err := tx.ExecContext(ctx, 
			"INSERT INTO timeline_tasks (id, timeline_id, title, description, start_date, end_date, duration, priority, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			task.ID, task.TimelineID, task.Title, task.Description, 
			task.StartDate, task.EndDate, task.Duration, task.Priority, task.Completed, task.CreatedAt, task.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to insert task: %v", err)
		}
	}

	return nil
}