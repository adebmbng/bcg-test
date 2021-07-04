create table promotions
(
    id int,
    item varchar(100) not null,
    promotion_type enum('free-item', 'buy-3-pay-2', 'discount') default 'discount' null,
    minimum_qty int default 0 null,
    promotion_data varchar(100) null,
    is_active tinyint 0,
    created_at timestamp default current_timestamp null,
    created_by varchar(50) null,
    updated_at timestamp null,
    updated_by varchar(50) null
);

create unique index promotions_id_uindex
	on promotions (id);

create index promotions_item_index
	on promotions (item);

alter table promotions
    add constraint promotions_pk
        primary key (id);

alter table promotions modify id int auto_increment;

