package models

import "time"

var schemaCategorys = `
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE golang.categorys (
    id_category uuid NULL DEFAULT uuid_generate_v4(),
    category_title varchar NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NULL,
    CONSTRAINT category_pk PRIMARY KEY (id_category)
);
`

type Category struct {
	Id_category    string     `db:"id_category" form:"id_category" json:"id_category"`
	Category_title string     `db:"category_title" form:"category_title" json:"category_title"`
	CreatedAt      *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updated_at"`
}
