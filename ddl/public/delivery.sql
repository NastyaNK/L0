create table delivery
(
    phone   varchar(30) not null
        primary key,
    name    varchar(30) not null,
    zip     integer     not null,
    city    varchar(30) not null,
    address varchar(30) not null,
    region  varchar(30) not null,
    email   varchar(30) not null
);

alter table delivery
    owner to anastasia;

