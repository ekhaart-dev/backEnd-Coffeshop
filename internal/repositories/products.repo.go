package repositories

import (
	"backEnd_Coffeshop/internal/models"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type RepoProduct struct {
	*sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

func (r *RepoProduct) GetBy(page int, limit int, search string, orderby string) ([]models.Products, error) {
	var list_products_data []models.Products
	if search == "" {
		search = ""
	} else {
		search = fmt.Sprintf(` AND LOWER(product_title) like LOWER('%s')`, "%"+search+"%")
	}
	if orderby == "" {
		orderby = ""
	} else {
		orderby = fmt.Sprintf(` ORDER BY %s`, orderby)
	}
	query := `
		SELECT p.id_product, p.banner_product, p.product_title, p.price, p.favorite, c.category_title, p.created_at, p.updated_at
		FROM golang.products p
		INNER JOIN golang.categorys c ON p.id_category = c.id_category
		WHERE TRUE ` + search + orderby + ` LIMIT $1 OFFSET $2`
	rows, err := r.Queryx(query, limit, page)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var product models.Products
		err := rows.Scan(
			&product.Id_product,
			&product.Banner_product,
			&product.Product_title,
			&product.Price,
			&product.Favorite,
			&product.Id_category,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		list_products_data = append(list_products_data, product)
	}
	rows.Close()
	return list_products_data, nil
}

func (r *RepoProduct) Get_Count_by_Id(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM golang.products WHERE id_product=$1", id)
	return count_data
}

func (r *RepoProduct) Get_Count_Data(search string) int {
	if search == "" {
		search = ""
	} else {
		search = fmt.Sprintf(` AND LOWER(product_title) like LOWER('%s')`, "%"+search+"%")
	}
	var id int
	r.Get(&id, `SELECT count(*) FROM golang.products WHERE TRUE `+search)
	return id
}

// func (r *RepoProduct) CreateProductData(data *models.Products) (string, error) {
// 	q := `INSERT INTO golang.products (id_category, banner_product, product_title, price, favorite)
// 	VALUES(:id_category, :banner_product, :product_title, :price, :favorite)`

// 	_, err := r.NamedExec(q, data)
// 	if err != nil {
// 		return "", err
// 	}

// 	return "add product data successful", nil
// }

func (r *RepoProduct) CreateProductData(data *models.Products) (string, error) {
	q := `INSERT INTO golang.products(
		id_category,
		banner_product,
		product_title,
		price,
		favorite
	) VALUES (
		:id_category,
		:banner_product,
		:product_title,
		:price,
		:favorite
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	return "1 data Products created", nil

}

func (r *RepoProduct) UpdateProduct(data *models.Products) (string, error) {

	tx := r.MustBegin()
	_, err := tx.NamedExec(
		`UPDATE golang.products SET
			id_category = :id_category,
			banner_product = :banner_product,
			product_title = :product_title,
			price = :price,
			favorite = :favorite,
			updated_at = :updated_at
		WHERE id_product = :id_product`, data)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	tx.Commit()

	return "1 data Product updated", nil
}

func (r *RepoProduct) DeleteProduct(idProduct string) (string, error) {
	q := `
		DELETE FROM golang.products WHERE id_product = $1
	`

	_, err := r.Exec(q, idProduct)
	if err != nil {
		return "", err
	}

	return "1 data Product deleted", nil
}
