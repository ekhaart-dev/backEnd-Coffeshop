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