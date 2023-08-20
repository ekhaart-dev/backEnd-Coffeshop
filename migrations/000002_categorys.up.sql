CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE golang.categorys (
    id_category uuid NULL DEFAULT uuid_generate_v4(),
    category_title varchar NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NULL,
    CONSTRAINT category_pk PRIMARY KEY (id_category)
);