package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jukemori/timeline-generator/internal/database"
	"github.com/jukemori/timeline-generator/internal/models"
)

// UserRepository handles database operations for users
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.DB,
	}
}

// Create creates a new user
func (r *UserRepository) Create(email string) (*models.User, error) {
	user := &models.User{
		ID:        uuid.New().String(),
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := "INSERT INTO users (id, email, created_at, updated_at) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, user.ID, user.Email, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetByID gets a user by ID
func (r *UserRepository) GetByID(id string) (*models.User, error) {
	query := "SELECT id, email, created_at, updated_at FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetByEmail gets a user by email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := "SELECT id, email, created_at, updated_at FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}