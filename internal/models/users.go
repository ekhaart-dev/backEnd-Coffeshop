package models

import (
	"time"
)

var schemaUsers = `
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE golang.users (
    id_user uuid NULL DEFAULT uuid_generate_v4(),
    username varchar NOT NULL,
    "password" VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    "role" VARCHAR NOT NULL,
	first_name varchar NOT NULL,
    last_name varchar NOT NULL,
    "address" VARCHAR NOT NULL,
    birthday DATE NOT NULL,
    gender VARCHAR NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NULL,
    CONSTRAINT user_pk PRIMARY KEY (id_user)
);
`

type Users struct {
	Id_user    string     `db:"id_user" json:"user_id,omitempty" form:"user_id" valid:"-"`
	Username   string     `db:"username" json:"username" form:"username" valid:"type(string)"`
	Password   string     `db:"password" json:"password,omitempty" valid:"stringlength(6|10)~Password minimal 6"`
	Email      string     `db:"email" form:"email" json:"email" valid:"type(string)"`
	Role       string     `db:"role" json:"role,omitempty" valid:"-"`
	First_name string     `db:"first_name" form:"first_name" json:"first_name" valid:"-"`
	Last_name  string     `db:"last_name" form:"last_name" json:"last_name" valid:"-"`
	Address    string     `db:"address" form:"address" json:"address" valid:"-"`
	Birthday   string     `db:"birthday" form:"birthday" json:"birthday" valid:"-"`
	Gender     string     `db:"gender" form:"gender" json:"gender" valid:"-"`
	CreatedAt  *time.Time `db:"created_at" json:"created_at" valid:"-"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}
