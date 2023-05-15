create table product_type (
    product_type_id int primary key,
    product_id int, 
    product_type_name nvarchar(30),
    
    foreign key (product_id) references products(id)
);

create table product_sub_type (
    product_sub_type_id integer primary key,
    product_type_id integer,
    product_sub_type_name nvarchar(30),
    
    foreign key (product_type_id) references product_type(product_type_id)
);

create table product_item (
    product_item_id integer primary key,
    product_sub_type_id integer,
    product_item_name nvarchar(10),
    
    foreign key (product_sub_type_id) references product_sub_type(product_sub_type_id)
);

create table item (
    item_id integer primary key,
    product_item_id integer,
    item_name nvarchar(30),
    item_price decimal(7,2),
    
    foreign key (product_item_id) references product_item(product_item_id)
);