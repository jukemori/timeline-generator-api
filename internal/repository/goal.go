package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jukemori/timeline-generator/internal/database"
	"github.com/jukemori/timeline-generator/internal/models"
)

// GoalRepository handles database operations for goals
type GoalRepository struct {
	db *sql.DB
}

// NewGoalRepository creates a new GoalRepository
func NewGoalRepository() *GoalRepository {
	return &GoalRepository{
		db: database.DB,
	}
}

// Create creates a new goal
func (r *GoalRepository) Create(userID, title, description, currentLevel, targetLevel string, startDate, targetDate time.Time) (*models.Goal, error) {
	goal := &models.Goal{
		ID:           uuid.New().String(),
		UserID:       userID,
		Title:        title,
		Description:  description,
		CurrentLevel: currentLevel,
		TargetLevel:  targetLevel,
		StartDate:    startDate,
		TargetDate:   targetDate,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	query := `INSERT INTO goals 
	(id, user_id, title, description, current_level, target_level, start_date, target_date, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	
	_, err := r.db.Exec(
		query, 
		goal.ID, 
		goal.UserID, 
		goal.Title, 
		goal.Description, 
		goal.CurrentLevel, 
		goal.TargetLevel, 
		goal.StartDate, 
		goal.TargetDate, 
		goal.CreatedAt, 
		goal.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}

	return goal, nil
}

// GetByID gets a goal by ID
func (r *GoalRepository) GetByID(id string) (*models.Goal, error) {
	query := `SELECT 
	id, user_id, title, description, current_level, target_level, start_date, target_date, created_at, updated_at 
	FROM goals WHERE id = ?`
	
	row := r.db.QueryRow(query, id)

	goal := &models.Goal{}
	err := row.Scan(
		&goal.ID, 
		&goal.UserID, 
		&goal.Title, 
		&goal.Description, 
		&goal.CurrentLevel, 
		&goal.TargetLevel, 
		&goal.StartDate, 
		&goal.TargetDate, 
		&goal.CreatedAt, 
		&goal.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}

	return goal, nil
}

// GetByUserID gets all goals for a user
func (r *GoalRepository) GetByUserID(userID string) ([]*models.Goal, error) {
	query := `SELECT 
	id, user_id, title, description, current_level, target_level, start_date, target_date, created_at, updated_at 
	FROM goals WHERE user_id = ?`
	
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	goals := []*models.Goal{}
	for rows.Next() {
		goal := &models.Goal{}
		err := rows.Scan(
			&goal.ID, 
			&goal.UserID, 
			&goal.Title, 
			&goal.Description, 
			&goal.CurrentLevel, 
			&goal.TargetLevel, 
			&goal.StartDate, 
			&goal.TargetDate, 
			&goal.CreatedAt, 
			&goal.UpdatedAt,
		)
		
		if err != nil {
			return nil, err
		}
		goals = append(goals, goal)
	}

	return goals, nil
}