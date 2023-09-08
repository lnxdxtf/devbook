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


DROP TABLE IF EXISTS followers;
CREATE TABLE IF NOT EXISTS followers(
    user_id int not null,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    follower_id int not null,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,

    primary key (user_id, follower_id)

) ENGINE=INNODB;

DROP TABLE IF EXISTS posts;
CREATE TABLE IF NOT EXISTS posts(
    id int auto_increment primary key,
    author_id int not null,
    title varchar(255) not null,
    content varchar(500) not null,
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE,

    likes int default 0,
    created_at timestamp default current_timestamp()
) ENGINE=INNODB;