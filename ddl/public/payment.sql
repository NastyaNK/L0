create table payment
(
    transaction   varchar(30) not null
        primary key,
    request_id    varchar(30),
    currency      varchar(30) not null,
    provider      varchar(30) not null,
    amount        integer     not null,
    payment_dt    integer     not null,
    bank          varchar(30) not null,
    delivery_cost integer     not null,
    goods_total   integer     not null,
    custom_fee    integer     not null
);

alter table payment
    owner to anastasia;

