CREATE TABLE golang.products (
    id_product uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    id_category uuid,
    banner_product varchar,
    product_title varchar NOT NULL,
    price int NOT NULL,
    favorite boolean NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone,
    CONSTRAINT product_fk FOREIGN KEY (id_category) REFERENCES golang.categorys(id_category)
);