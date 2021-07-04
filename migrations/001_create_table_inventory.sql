create table inventories
(
    Id int comment 'Primary key',
    sku varchar(100) not null,
    name varchar(255) null,
    price float not null,
    qty int default 0 null,
    is_active tinyint 0,
    created_at timestamp default current_timestamp null,
    created_by varchar(50) null,
    updated_at timestamp null,
    updated_by varchar(50) null
);

create unique index inventories_Id_uindex
	on inventories (Id);

create unique index inventories_sku_uindex
	on inventories (sku);

alter table inventories
    add constraint inventories_pk
        primary key (Id);

