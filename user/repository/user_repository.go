package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/erhemdiputra/go-crud/user"

	"github.com/erhemdiputra/go-crud/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(conn *sql.DB) user.IUserRepository {
	return &UserRepository{conn}
}

func (u *UserRepository) GetList(ctx context.Context) ([]models.User, error) {
	query := `SELECT id, name, age FROM users`

	rows, err := u.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	result := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		result = append(result, user)
	}

	return result, nil
}

func (u *UserRepository) Add(ctx context.Context, user models.User) (int64, error) {
	query := `INSERT INTO users (name, age) VALUES (?, ?)`
	stmt, err := u.DB.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, user.Name, user.Age)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (u *UserRepository) GetByID(ctx context.Context, id int64) (models.User, error) {
	query := `SELECT id, name, age FROM users WHERE id = ?`

	row := u.DB.QueryRowContext(ctx, query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		log.Println(err)
		return models.User{}, err
	}

	return user, nil
}

func (u *UserRepository) Update(ctx context.Context, user models.User) (int64, error) {
	query := `UPDATE users SET name = ?, age = ? WHERE id = ?`

	stmt, err := u.DB.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, user.Name, user.Age, user.ID)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (u *UserRepository) Delete(ctx context.Context, id int64) (int64, error) {
	query := `DELETE FROM users WHERE id = ?`

	stmt, err := u.DB.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
