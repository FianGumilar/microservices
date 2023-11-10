package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/FianGumilar/microservices/user-service/interfaces"
	"github.com/FianGumilar/microservices/user-service/models/dto"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRespository(db *sql.DB) interfaces.UserRepository {
	return &UserRepository{db: db}
}

// CountAllUsers implements interfaces.UserRepository.
func (r *UserRepository) CountAllUsers() (int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var count int32

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	query := `SELECT COUNT(*) FROM users`
	err = r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, nil
	}
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}

		return 0, err
	}

	return count, nil
}

// DeleteUser implements interfaces.UserRepository.
func (r *UserRepository) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `DELETE FROM user WHERE id = $1`
	_, err = r.db.QueryContext(ctx, query, id)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

// FindAllUsers implements interfaces.UserRepository.
func (r *UserRepository) FindAllUsers(pagination *dto.Pagination) ([]dto.FindAllUsersDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var users []dto.FindAllUsersDTO

	query := `SELECT * FROM users ORDER BY id DESC LIMIT $1 OFFSET $2`
	rows, err := r.db.QueryContext(ctx, query, pagination.Limit, pagination.Offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user dto.FindAllUsersDTO
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// FindUserByEmail implements interfaces.UserRepository.
func (r *UserRepository) FindUserByEmail(email string) (*dto.FindUserDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var user dto.FindUserDTO

	query := `SELECT * FROM users WHERE email = $1`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindUserByID implements interfaces.UserRepository.
func (r *UserRepository) FindUserByID(id string) (*dto.FindUserDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var user dto.FindUserDTO

	query := `SELECT * FROM users WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// InsertUser implements interfaces.UserRepository.
func (r *UserRepository) InsertUser(user *dto.InsertUserDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = r.db.ExecContext(ctx, query, &user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return tx.Commit()
}

// UpdateUser implements interfaces.UserRepository.
func (r *UserRepository) UpdateUser(id string, user *dto.UpdateUserDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE users SET name = $1, email = $2, updated_at = $3 WHERE id = $4`
	_, err = tx.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.UpdatedAt)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
