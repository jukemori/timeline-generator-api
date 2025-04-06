package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jukemori/timeline-generator/internal/database"
	"github.com/jukemori/timeline-generator/internal/models"
)

// TaskRepository handles database operations for timeline tasks
type TaskRepository struct {
	db *sql.DB
}

// NewTaskRepository creates a new TaskRepository
func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		db: database.DB,
	}
}

// Create creates a new timeline task
func (r *TaskRepository) Create(timelineID, title, description, duration string, startDate, endDate time.Time, priority int) (*models.TimelineTask, error) {
	task := &models.TimelineTask{
		ID:          uuid.New().String(),
		TimelineID:  timelineID,
		Title:       title,
		Description: description,
		StartDate:   startDate,
		EndDate:     endDate,
		Duration:    duration,
		Priority:    priority,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	query := `INSERT INTO timeline_tasks 
	(id, timeline_id, title, description, start_date, end_date, duration, priority, completed, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	
	_, err := r.db.Exec(
		query, 
		task.ID, 
		task.TimelineID, 
		task.Title, 
		task.Description, 
		task.StartDate, 
		task.EndDate, 
		task.Duration, 
		task.Priority, 
		task.Completed, 
		task.CreatedAt, 
		task.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}

	return task, nil
}

// GetByID gets a task by ID
func (r *TaskRepository) GetByID(id string) (*models.TimelineTask, error) {
	query := `SELECT 
	id, timeline_id, title, description, start_date, end_date, duration, priority, completed, created_at, updated_at 
	FROM timeline_tasks WHERE id = ?`
	
	row := r.db.QueryRow(query, id)

	task := &models.TimelineTask{}
	err := row.Scan(
		&task.ID, 
		&task.TimelineID, 
		&task.Title, 
		&task.Description, 
		&task.StartDate, 
		&task.EndDate, 
		&task.Duration, 
		&task.Priority, 
		&task.Completed, 
		&task.CreatedAt, 
		&task.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}

	return task, nil
}

// GetByTimelineID gets all tasks for a timeline
func (r *TaskRepository) GetByTimelineID(timelineID string) ([]models.TimelineTask, error) {
	query := `SELECT 
	id, timeline_id, title, description, start_date, end_date, duration, priority, completed, created_at, updated_at 
	FROM timeline_tasks WHERE timeline_id = ? ORDER BY start_date ASC, priority DESC`
	
	rows, err := r.db.Query(query, timelineID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []models.TimelineTask{}
	for rows.Next() {
		task := models.TimelineTask{}
		err := rows.Scan(
			&task.ID, 
			&task.TimelineID, 
			&task.Title, 
			&task.Description, 
			&task.StartDate, 
			&task.EndDate, 
			&task.Duration, 
			&task.Priority, 
			&task.Completed, 
			&task.CreatedAt, 
			&task.UpdatedAt,
		)
		
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// UpdateCompletionStatus updates a task's completion status
func (r *TaskRepository) UpdateCompletionStatus(id string, completed bool) error {
	query := "UPDATE timeline_tasks SET completed = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.Exec(query, completed, time.Now(), id)
	return err
}