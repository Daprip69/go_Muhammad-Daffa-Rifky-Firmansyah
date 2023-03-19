CREATE TABLE user (
    id int(11) not null primary key auto_increment,
    name varchar(255),
    email varchar(255),
    address varchar(255),
    status smallint(6) DEFAULT NULL,
    created_at timestamp default now(),
    updated_at timestamp default now()    
);
CREATE TABLE payment_methods (
	id int(11) not null auto_increment primary key,
	name varchar(255),
    status smallint(6) DEFAULT NULL,
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
    status smallint(6) DEFAULT NULL,
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
    status smallint(6) DEFAULT NULL,
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
INSERT INTO operators(name) VALUES
('TOKO Ayam'),
('TOKO Babi'),
('TOKO Cicak'),
('TOKO Dragon'),
('TOKO Elang');
INSERT INTO product_types(name) VALUES
('Monitor'),
('CPU'),
('Keyboard');
INSERT INTO products(product_type_id, operators_id, name, status) VALUES
(1, 3, 'kb1', 'keyboard rgb', 2),
(1, 3, 'mt1', 'monitor rgb', 1);
INSERT INTO products(product_type_id, operators_id, name, status) VALUES
(2, 1, 'kb2', 'keyboard no rgb', 4),
(2, 1, 'mt2', 'monitor no rgb', 3),
(2, 1, 'cp1', 'cpu rgb', 1);
INSERT INTO products(product_type_id, operators_id, name, status) VALUES
(3, 4, 'kb3', 'keyboard pro', 2),
(3, 4, 'mt3', 'monitor pro', 5),
(3, 4, 'cp2', 'cpu pro', 1);
INSERT INTO product_descriptions(description) VALUES
('keyboard dengan fitur rgb'),
('keyboard ttanpa fitur rgb'),
('keyboard kualitas pro'),
('monitor dengan fitur rgb'),
('monitor tanpa fitur rgb'),
('monitor kualitas pro'),
('cpu fitur rgb'),
('cpu tanpa fitur rgb');
INSERT INTO payment_methods(name, status) VALUES
('BNI', 1),
('BRI', 2),
('BRI', 3);
INSERT INTO users(name, status, email, address) VALUES
('kakashi', 1, 'k@g.com', 'jl. k'),
('sakura', 2, 's@g.com', 'jl. e'),
('orochimaru', 3, 'o@g.com', 'jl. k'),
('yanto', 4, '-y@gmail.com', 'jl. y'),
('antonio', 5, 'a@g.com', 'jl. a');
INSERT INTO transactions(users_id, methods_id, status, quantity, total_price) VALUES
(1, 1, 'ngutang', 2, 20000),
(2, 2, 'ngutang', 3, 30000),
(3, 3, 'cash', 4, 40000);
INSERT INTO product(product_type_id, operator_id, name, status) VALUES
(2, 1, 'kb2', 'keyboard no rgb', 4),
(2, 1, 'mt2', 'monitor no rgb', 3),
(2, 1, 'cp1', 'cpu rgb', 1);
SELECT name FROM users WHERE email = 'k@gmail.com';
SELECT * FROM products WHERE id = 3;
SELECT * FROM users WHERE create_at BETWEEN DATE_SUB(NOW(), INTERVAL 7 DAY) AND NOW() AND name LIKE '%a%';
SELECT COUNT(*) FROM users WHERE address = 'jl. k';
SELECT * FROM users ORDER BY name ASC;
SELECT * FROM products LIMIT 5;
UPDATE products SET name = 'products dummy' WHERE id = 1;
UPDATE transaction_details SET quantity = 3 WHERE products_id = 1;
DELETE FROM products WHERE id = 1;
DELETE FROM products WHERE products_type_id = 1;

"soal join";

SELECT * FROM transactions WHERE users_id = 1 OR user_id = 2;
SELECT SUM(total_price) FROM transaksi WHERE user_id = 1;
SELECT product.*, product_types.name FROM product JOIN product_types ON product.product_type_id = product_types.id;
SELECT transaksi.*, products.name, users.name 
FROM transactions
INNER JOIN product ON transaction.products_id = products.id
INNER JOIN user ON transaction.users_id = users.id;

