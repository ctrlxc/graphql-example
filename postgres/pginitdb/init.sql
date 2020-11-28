CREATE USER graphql WITH PASSWORD 'graphql';
CREATE DATABASE graphql OWNER=graphql TEMPLATE=template0 ENCODING=UTF8 LC_COLLATE='C';
GRANT ALL ON DATABASE graphql TO graphql;

\c graphql graphql

CREATE SCHEMA graphql;

CREATE TABLE shop (
    id bigserial primary key,
    shop_name text
);

INSERT INTO shop (shop_name) VALUES
    ('shop-name-1'),
    ('shop-name-2');

CREATE TABLE book (
    id bigserial primary key,
    book_title text
);

INSERT INTO book (book_title) VALUES
    ('book-title-1'),
    ('book-title-2'),
    ('book-title-3');

CREATE TABLE stock (
    id bigserial primary key,
    shop_id bigint references shop(id),
    book_id bigint references book(id)
);

INSERT INTO stock (shop_id, book_id) VALUES
    (1, 1),
    (1, 2),
    (2, 2),
    (2, 3);
