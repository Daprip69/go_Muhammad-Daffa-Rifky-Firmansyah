CREATE TABLE user (
    id int(11) not null primary key auto_increment,
    name varchar(255),
    email varchar(255),
    address varchar(255),
    created_at timestamp default now(),
    updated_at timestamp default now()    
);
CREATE TABLE payment_methods (
	id int(11) not null auto_increment primary key,
	name varchar(255),
    status BOOLEAN,
    created_at timestamp default now(),
    updated_at timestamp default now()
);
CREATE TABLE operators (
	id int(11) not null auto_increment primary key,
	name varchar(255),
    created_at timestamp default now(),
    updated_at timestamp default now()
);
CREATE TABLE product_types (
	id int(11) not null auto_increment primary key,
    name varchar(255),
    created_at timestamp default now(),
    updated_at timestamp default now()
);
CREATE TABLE products (
	id int(11) not null auto_increment primary key,
    product_type_id int(11),
    operators_id int(11),
    name varchar(255),
    quantity int(11),
    price decimal(10,2),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    foreign key (product_type_id) references product_types(id),
    foreign key (operators_id) references operators(id)
);
CREATE TABLE product_descriptions (
	id int(11) primary key,     
    description text,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    foreign key (id) references products(id)
);
CREATE TABLE transactions (
	id int(11) not null auto_increment primary key,
    users_id int(11),
    payment_id int(11),
    quantity int(11),
    total_price DECIMAL(11,2),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    foreign key (users_id) references users(id),
    foreign key (payment_id) references payment_methods(id)
);
CREATE TABLE transaction_details (
	id int(11) primary key auto_increment,
    product_id int(11),
    quantity int(11),
    transaction_id int(11),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    foreign key (product_id) references products(id),
    foreign key (transaction_id) references transactions(id)
);
CREATE TABLE kurir (
	id int(11) primary key auto_increment,
    name varchar(255),
    created_at timestamp default now(),
    updated_at timestamp default now()
);
ALTER TABLE kurir ADD COLUMN ongkos_dasar int(11);

ALTER TABLE kurir RENAME to shipping;

DROP TABLE shipping;