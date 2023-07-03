CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(100) not null unique,
    pswrd varchar(255) not null,
    created_at timestamp default current_timestamp()
) ENGINE=INNODB;