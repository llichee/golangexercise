create database if not exists users;

create table if not exists user (
    id bigint primary key auto_increment,
    name varchar(32) not null default '',
    gander boolean not null default true ,
    addr text,
    create_at datetime not null,
    update_at datetime not null,
    deleted_at datetime
) engine=innodb default charset utf8mb4;

alter table user add column if not exists remark text;