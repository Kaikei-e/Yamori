DROP SCHEMA IF EXISTS yamori;
CREATE DATABASE yamori;

USE yamori;

drop table if exists users;

create table yamori.users
(
    id                int AUTO_INCREMENT NOT NULL comment 'user id. auto incremented',
    created_at        DATETIME(6)        NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at        DATETIME(6)        NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    deleted_at        DATETIME(6) comment 'soft delete',
    email             varchar(255)       NOT NULL UNIQUE comment 'email address',
    pass_hash         varchar(255)       NOT NULL comment 'password',
    account_status    tinyint            NOT NULL DEFAULT 0 comment 'account status',
    user_name         varchar(255)       NOT NULL comment 'user name',
    user_organization int                NOT NULL comment 'user organization',
    user_role         int                NOT NULL DEFAULT 0 comment 'user role',
    INDEX idx_email (id, email),
    PRIMARY KEY (id)
) ENGINE = InnoDB comment 'users for managing guest appointment';