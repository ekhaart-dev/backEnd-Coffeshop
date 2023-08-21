CREATE TABLE golang.categorys (
    id_category uuid NULL DEFAULT gen_random_uuid(),
    category_title varchar NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NULL,
    CONSTRAINT category_pk PRIMARY KEY (id_category)
);