CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE golang.products (
    id_product uuid NULL DEFAULT uuid_generate_v4(),
    id_category uuid NULL,
    banner_product VARCHAR NULL,
    product_title VARCHAR NOT NULL,
    price INT NOT NULL,
    favorite BOOLEAN NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NULL,
    CONSTRAINT product_pk PRIMARY KEY (id_product),
    CONSTRAINT product_fk FOREIGN KEY (id_category) REFERENCES golang.categorys(id_category) ON DELETE CASCADE
);