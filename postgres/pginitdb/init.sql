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

INSERT INTO shop (shop_name, created_at, updated_at) VALUES
    ('shop-name-1', '2020/01/01 00:00:00'::timestamp without time zone, '2020/01/05 00:00:00'::timestamp without time zone),
    ('shop-name-2', '2020/01/02 00:00:00'::timestamp without time zone, '2020/01/04 00:00:00'::timestamp without time zone),
    ('shop-name-3', '2020/01/03 00:00:00'::timestamp without time zone, '2020/01/03 00:00:00'::timestamp without time zone),
    ('shop-name-4', '2020/01/04 00:00:00'::timestamp without time zone, '2020/01/02 00:00:00'::timestamp without time zone),
    ('shop-name-5', '2020/01/05 00:00:00'::timestamp without time zone, '2020/01/01 00:00:00'::timestamp without time zone)
;

CREATE TABLE book (
    id bigserial primary key,
    book_title text,
    created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL
);

INSERT INTO book (book_title, created_at, updated_at) VALUES
    ('book-title-1', '2020/01/01 00:00:00'::timestamp without time zone, '2020/01/03 00:00:00'::timestamp without time zone),
    ('book-title-2', '2020/01/02 00:00:00'::timestamp without time zone, '2020/01/02 00:00:00'::timestamp without time zone),
    ('book-title-3', '2020/01/03 00:00:00'::timestamp without time zone, '2020/01/01 00:00:00'::timestamp without time zone)
;

CREATE TABLE stock (
    id bigserial primary key,
    shop_id bigint references shop(id),
    book_id bigint references book(id),
    created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL
);

INSERT INTO stock (shop_id, book_id, created_at, updated_at) VALUES
    (1, 1, now(), now()),
    (1, 2, now(), now()),
    (2, 2, now(), now()),
    (2, 3, now(), now())
;
