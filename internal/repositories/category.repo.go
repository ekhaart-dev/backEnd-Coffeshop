package repositories

import (
	"backEnd_Coffeshop/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoCategory struct {
	*sqlx.DB
}

func NewCategory(db *sqlx.DB) *RepoCategory {
	return &RepoCategory{db}
}

func (r *RepoCategory) GetCategory() ([]models.Category, error) {
	var categories []models.Category

	q := `SELECT * FROM golang.categorys ORDER BY created_at DESC`

	err := r.Select(&categories, q)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *RepoCategory) CreateCategory(data *models.Category) (string, error) {
	q := `INSERT INTO golang.categorys(
		category_title
	) VALUES (
		:category_title
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	return "1 data Category created", nil

}

func (r *RepoCategory) DeleteCategory(idCategorys string) (string, error) {
	q := `
		DELETE FROM golang.categorys WHERE id_category = $1
	`

	_, err := r.Exec(q, idCategorys)
	if err != nil {
		return "", err
	}

	return "1 data Category deleted", nil
}
