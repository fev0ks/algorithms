-- id | name       | client_name | total
-- ----+------------+-------------+-------
--  10 | tool       | Pineapple   | 29000
--  40 | electronic | Pineapple   | 29000
--  30 | outdoor    | Booble      |   865
--  20 | office     | Fakebook    |    36
-- (4 rows)

create table clients (
                         id bigint primary key,
                         name varchar not null unique
);

create table products (
                          id bigint primary key,
                          name varchar not null unique
);

create table product_prices (
                                id bigint primary key,
                                product_id bigint,
                                price bigint,
                                created_at timestamp not null,
                                constraint fk_product_prices_product_id foreign key (product_id) references products (id)
);

create unique index on product_prices (product_id, created_at);

create table product_tags (
                              id bigint primary key,
                              name varchar not null unique
);

create table products_product_tags (
                                       product_id bigint,
                                       product_tag_id bigint,
                                       constraint fk_products_product_tags_product_id foreign key (product_id) references products (id),
                                       constraint fk_products_product_tags_product_tag_id foreign key (product_tag_id) references product_tags (id)
);

create unique index on products_product_tags (product_id, product_tag_id);

create table orders (
                        id bigint primary key,
                        client_id bigint,
                        product_id bigint,
                        quantity int,
                        created_at timestamp not null,
                        constraint fk_orders_client_id foreign key (client_id) references clients (id),
                        constraint fk_orders_product_id foreign key (product_id) references products (id)
);

insert into clients
values (1, 'Booble'),
       (2, 'Macrohard'),
       (3, 'Pineapple'),
       (4, 'Fakebook');

insert into products
values (1, 'Paper'),
       (2, 'Laptop'),
       (3, 'Shovel');

insert into product_prices
values (1, 1, 5, '2023-04-01 12:30:45'),
       (2, 3, 20, '2023-04-01 17:45:00'),
       (3, 2, 900, '2023-04-02 00:00:00'),
       (4, 1, 6, '2023-04-02 12:30:45'),
       (5, 3, 15, '2023-04-02 09:00:00'),
       (6, 2, 1250, '2023-04-03 00:00:00'),
       (7, 3, 25, '2023-04-01 19:45:00'),
       (8, 2, 1000, '2023-04-01 00:00:00');

insert into product_tags
values (10, 'tool'),
       (20, 'office'),
       (30, 'outdoor'),
       (40, 'electronic'),
       (50, 'none');

insert into products_product_tags
values (1, 20),
       (2, 10),
       (2, 40),
       (3, 10),
       (3, 30);

insert into orders
values (1, 3, 2, 10, '2023-04-01 09:10:00'),
       (2, 3, 2, 10, '2023-04-01 23:59:59'),
       (3, 3, 2, 10, '2023-04-02 00:00:00'),
       (4, 3, 1, 1, '2023-03-02 00:00:00'),
       (5, 3, 1, 1, '2023-04-02 00:00:00'),
       (6, 1, 3, 2, '2023-04-01 17:45:00'),
       (7, 1, 3, 33, '2023-04-01 19:45:01'),
       (8, 4, 1, 6, '2024-04-01 19:45:01'),
       (9, 4, 2, 7, '2024-04-01 19:45:01'),
       (10, 4, 3, 10, '2024-04-01 19:45:01');
