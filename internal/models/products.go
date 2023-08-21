package models

import (
	"time"
)

// var schemaProduct = `
// CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
// CREATE TABLE golang.products (
//     id_product uuid NULL DEFAULT uuid_generate_v4(),
//     id_category uuid NULL,
//     banner_product VARCHAR NULL,
//     product_title VARCHAR NOT NULL,
//     price INT NOT NULL,
//     favorite BOOLEAN NOT NULL,
//     created_at timestamp without time zone NOT NULL DEFAULT now(),
//     updated_at timestamp without time zone NULL,
//     CONSTRAINT product_pk PRIMARY KEY (id_product),
//     CONSTRAINT product_fk FOREIGN KEY (id_category) REFERENCES golang.categorys(id_category) ON DELETE CASCADE
// );
// `

type Products struct {
	Id_product     string     `db:"id_product" form:"id_product" json:"id_product"`
	Id_category    string     `db:"id_category" form:"id_category" json:"id_category"`
	Banner_product string     `db:"banner_product" json:"banner_product,omitempty" valid:"-"`
	Product_title  string     `db:"product_title" form:"product_title" json:"product_title"`
	Price          int        `db:"price" form:"price" json:"price"`
	Favorite       bool       `db:"favorite" form:"favorite" json:"favorite"`
	CreatedAt      *time.Time `db:"created_at" form:"created_at" json:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at" form:"updated_at" json:"updated_at"`
}

type ProductSearchParams struct {
	Search   string `json:"search"`
	Category string `json:"category"`
	OrderBy  string `json:"order_by"`
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
}

type Meta_Products struct {
	Next       string
	Prev       string
	Last_page  string
	Total_data string
}
