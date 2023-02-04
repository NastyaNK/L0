create table items
(
    chrt_id      integer     not null
        primary key,
    track_number varchar(30) not null,
    price        integer     not null,
    rid          varchar(30) not null,
    name         varchar(30) not null,
    sale         integer     not null,
    size         integer     not null,
    total_price  integer     not null,
    nm_id        integer     not null,
    brand        varchar(30) not null,
    status       integer     not null,
    "order"      varchar(30) not null
        constraint items_orders_order_uid_fk
            references orders
);

alter table items
    owner to anastasia;

