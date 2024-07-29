SELECT * FROM sample;
SELECT * FROM sample;

CREATE TABLE users
(
    id VARCHAR(100) not NULL,
    password VARCHAR(100) not NULL,
    name VARCHAR(100) not NULL,
    created_at TIMESTAMP not NULL default current_timestamp,
    updated_at TIMESTAMP not NULL default current_timestamp on update current_timestamp,
    primary key (id)
) engine = InnoDB;

SELECT * FROM users;

ALTER TABLE users
rename column name to first_name;


ALTER TABLE users
add column middle_name VARCHAR(100) NULL after first_name;

ALTER TABLE users
add column last_name VARCHAR(100) NULL after middle_name;

delete FROM users where id = '';


CREATE TABLE users_logs
(
    id int auto_increment,
    user_id VARCHAR(100) not NULL,
    action VARCHAR(100) not NULL,
    created_at TIMESTAMP not NULL default current_timestamp,
    updated_at TIMESTAMP not NULL default current_timestamp on update current_timestamp,
    primary key (id)
) engine = InnoDB;

SELECT * FROM users_logs;

DESCRIBE users_logs;

delete FROM users_logs;

ALTER TABLE users_logs
modify created_at bigint not null;

ALTER TABLE users_logs
modify updated_at bigint not null;



CREATE TABLE todos
(
    id BIGINT not null auto_increment,
    user_id VARCHAR(100) not NULL,
    title VARCHAR(100) not NULL,
    description text null,
    created_at TIMESTAMP not NULL default current_timestamp,
    updated_at TIMESTAMP not NULL default current_timestamp on update current_timestamp,
    deleted_at TIMESTAMP NULL ,
    primary key (id)
) engine = InnoDB;

SELECT * FROM todos;

DESCRIBE todos;


CREATE TABLE wallets
(
    id VARCHAR(100) not null,
    user_id VARCHAR(100) not NULL,
    balance bigint not NULL,
    created_at TIMESTAMP not NULL default current_timestamp,
    updated_at TIMESTAMP not NULL default current_timestamp on update current_timestamp,
    primary key (id),
    foreign key (user_id) references users (id)
) engine = InnoDB;

SELECT * FROM wallets;

DESCRIBE wallets;

CREATE TABLE addresses
(
    id INT not null auto_increment,
    user_id VARCHAR(100) not NULL,
    address VARCHAR(100) not NULL,
    created_at TIMESTAMP not NULL default current_timestamp,
    updated_at TIMESTAMP not NULL default current_timestamp on update current_timestamp,
    primary key (id),
    foreign key (user_id) references users (id)
) engine = InnoDB;

DESCRIBE addresses;
SELECT * FROM addresses;


CREATE TABLE products
(
    id VARCHAR(100) not null,
    name VARCHAR(100) not NULL,
    price BIGINT not NULL,
    created_at TIMESTAMP not NULL default current_timestamp,
    updated_at TIMESTAMP not NULL default current_timestamp on update current_timestamp,
    primary key (id)
) engine = InnoDB;

DESCRIBE products;
SELECT * FROM products;
delete FROM products;



CREATE TABLE user_like_product
(
    user_id VARCHAR(100) not null,
    product_id VARCHAR(100) not NULL,
    primary key (user_id, product_id),
    foreign key (user_id) references users (id),
    foreign key (product_id) references products (id)
) engine = InnoDB;

DESCRIBE user_like_product;
delete FROM user_like_product;

SELECT * FROM users

SELECT count(id) from users;

DESCRIBE guest_books;