DROP SCHEMA IF EXISTS yamori;
CREATE DATABASE yamori;

create table yamori.users
(
    id             int AUTO_INCREMENT NOT NULL comment 'user id. auto incremented',
    created_at     DATETIME(6)        NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at     DATETIME(6)        NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    deleted_at     DATETIME(6)        NOT NULL comment 'soft delete',
    email          varchar(255)       NOT NULL comment 'email address',
    password       varchar(255)       NOT NULL comment 'password',
    clearance      tinyint            NOT NULL DEFAULT 0 comment 'is admin user',
    account_status tinyint            NOT NULL DEFAULT 0 comment 'account status',
    user_name      varchar(255)       NOT NULL comment 'user name',
    INDEX         idx_email (id, email),
    PRIMARY KEY   (id)
) ENGINE = InnoDB comment 'users for managing guest appointment';
