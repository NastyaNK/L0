create table orders
(
    order_uid          varchar(30) not null
        constraint mytable_pkey
            primary key,
    track_number       varchar(30) not null,
    entry              varchar(30) not null,
    delivery           varchar(30) not null
        constraint orders_delivery_phone_fk
            references delivery,
    payment            varchar(30) not null
        constraint orders_payment_transaction_fk
            references payment,
    locale             varchar(30) not null,
    internal_signature varchar(30),
    customer_id        varchar(30) not null,
    delivery_service   varchar(30) not null,
    shardkey           integer     not null,
    sm_id              integer     not null,
    date_created       varchar(30) not null,
    oof_shard          integer     not null
);

alter table orders
    owner to anastasia;

