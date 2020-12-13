create database if not exists testuser default charset utf8mb4;

create table if not exists users(
    id bigint primary key auto_increment,
    name varchar(32) not null default '',
    addr text
) engine=innodb default charset utf8mb4

create table if not exists users(
    id bigint primary key auto_increment,

)