package repositories

import (
	"backEnd_Coffeshop/config"
	"backEnd_Coffeshop/internal/models"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
)

type RepoUsers struct {
	*sqlx.DB
}

func NewUsers(db *sqlx.DB) *RepoUsers {
	return &RepoUsers{db}
}

func (r *RepoUsers) CreateUsers(data *models.Users) (*config.Result, error) {
	q := `INSERT INTO golang.users(
		username,
		password,
		email,
		role,
		first_name,
		last_name,
		address,
		birthday,
		gender
	)
		VALUES(
			:username,
			:password,
			:email,
			:role,
			:first_name,
			:last_name,
			:address,
			:birthday,
			:gender
		)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &config.Result{Message: "1 data user created"}, nil
}

func (r *RepoUsers) GetAllUser() (*config.Result, error) {
	var data []models.Users
	q := `SELECT * FROM golang.users ORDER BY created_at DESC`

	if err := r.Select(&data, q); err != nil {
		return nil, err
	}

	return &config.Result{Data: data}, nil
}

func (r *RepoUsers) GetAuthData(user string) (*models.Users, error) {
	var result models.Users
	q := `SELECT id_user, username, "role", "password" FROM golang.users WHERE username = ?`

	if err := r.Get(&result, r.Rebind(q), user); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("username not found")
		}

		return nil, err
	}

	return &result, nil
}

func (r *RepoUsers) UpdateUsers(data *models.Users) (string, error) {

	tx := r.MustBegin()
	_, err := tx.NamedExec(
		`UPDATE golang.users SET
			username = :username,
			password = :password,
			email = :email,
			role = :role,
			first_name = :first_name,
			last_name = :last_name,
			address = :address,
			birthday = :birthday,
			gender = :gender
		WHERE id_user = :id_user`, data)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	tx.Commit()

	return "1 data Users updated", nil
}

func (r *RepoUsers) DeleteUser(idUser string) (string, error) {
	q := `
		DELETE FROM golang.users WHERE id_user = $1
	`

	_, err := r.Exec(q, idUser)
	if err != nil {
		return "", err
	}

	return "1 data Users deleted", nil
}
