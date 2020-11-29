CREATE USER graphql WITH PASSWORD 'graphql';
CREATE DATABASE graphql OWNER=graphql TEMPLATE=template0 ENCODING=UTF8 LC_COLLATE='C';
GRANT ALL ON DATABASE graphql TO graphql;

\c graphql graphql

CREATE SCHEMA graphql;

CREATE TABLE shop (
    id bigserial primary key,
    shop_name text,
    created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL
);

INSERT INTO shop (shop_name) VALUES
    ('shop-name-1', now(), now()),
    ('shop-name-2', now(), now());

CREATE TABLE book (
    id bigserial primary key,
    book_title text,
    created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL
);

INSERT INTO book (book_title) VALUES
    ('book-title-1', now(), now()),
    ('book-title-2', now(), now()),
    ('book-title-3', now(), now());

CREATE TABLE stock (
    id bigserial primary key,
    shop_id bigint references shop(id),
    book_id bigint references book(id),
    created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL
);

INSERT INTO stock (shop_id, book_id) VALUES
    (1, 1, now(), now()),
    (1, 2, now(), now()),
    (2, 2, now(), now()),
    (2, 3, now(), now());
