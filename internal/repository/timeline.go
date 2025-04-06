package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jukemori/timeline-generator/internal/database"
	"github.com/jukemori/timeline-generator/internal/models"
)

// TimelineRepository handles database operations for timelines
type TimelineRepository struct {
	db *sql.DB
}

// NewTimelineRepository creates a new TimelineRepository
func NewTimelineRepository() *TimelineRepository {
	return &TimelineRepository{
		db: database.DB,
	}
}

// Create creates a new timeline
func (r *TimelineRepository) Create(goalID, title, description string, startDate, endDate time.Time) (*models.Timeline, error) {
	timeline := &models.Timeline{
		ID:          uuid.New().String(),
		GoalID:      goalID,
		Title:       title,
		Description: description,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	query := `INSERT INTO timelines 
	(id, goal_id, title, description, start_date, end_date, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	
	_, err := r.db.Exec(
		query, 
		timeline.ID, 
		timeline.GoalID, 
		timeline.Title, 
		timeline.Description, 
		timeline.StartDate, 
		timeline.EndDate, 
		timeline.CreatedAt, 
		timeline.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}

	return timeline, nil
}

// GetByID gets a timeline by ID with its tasks
func (r *TimelineRepository) GetByID(id string) (*models.Timeline, error) {
	query := `SELECT 
	id, goal_id, title, description, start_date, end_date, created_at, updated_at 
	FROM timelines WHERE id = ?`
	
	row := r.db.QueryRow(query, id)

	timeline := &models.Timeline{}
	err := row.Scan(
		&timeline.ID, 
		&timeline.GoalID, 
		&timeline.Title, 
		&timeline.Description, 
		&timeline.StartDate, 
		&timeline.EndDate, 
		&timeline.CreatedAt, 
		&timeline.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}

	// Get tasks for this timeline
	taskRepo := NewTaskRepository()
	tasks, err := taskRepo.GetByTimelineID(timeline.ID)
	if err != nil {
		return nil, err
	}
	
	timeline.Tasks = tasks
	return timeline, nil
}

// GetByGoalID gets all timelines for a goal
func (r *TimelineRepository) GetByGoalID(goalID string) ([]*models.Timeline, error) {
	query := `SELECT 
	id, goal_id, title, description, start_date, end_date, created_at, updated_at 
	FROM timelines WHERE goal_id = ?`
	
	rows, err := r.db.Query(query, goalID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	timelines := []*models.Timeline{}
	for rows.Next() {
		timeline := &models.Timeline{}
		err := rows.Scan(
			&timeline.ID, 
			&timeline.GoalID, 
			&timeline.Title, 
			&timeline.Description, 
			&timeline.StartDate, 
			&timeline.EndDate, 
			&timeline.CreatedAt, 
			&timeline.UpdatedAt,
		)
		
		if err != nil {
			return nil, err
		}
		timelines = append(timelines, timeline)
	}

	return timelines, nil
}