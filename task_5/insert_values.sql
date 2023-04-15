-- INSERTION QUERIES FOR customers TABLE
INSERT INTO customers(
    customer_id, first_name, last_name, email) 
    VALUES(100, 'Abhijith', 'A', 'abhi@abc.com');

INSERT INTO customers(
    customer_id, first_name, last_name, email) 
    VALUES(101, 'Aanand', 'Jaya', 'aanand@abc.com');

INSERT INTO customers(
    customer_id, first_name, last_name, email) 
    VALUES(102, 'Aswathi', 'Vijay', 'aswathiachu@abc.com');

INSERT INTO customers(
    customer_id, first_name, last_name, email) 
    VALUES(103, 'Anil', 'C J', 'anil@abc.com');

INSERT INTO customers(
    customer_id, first_name, last_name, email) 
    VALUES(104, 'Bibin', 'P', 'bibin@abc.com');

INSERT INTO customers(
    customer_id, first_name, last_name, email) 
    VALUES(105, 'Hareendran', 'K', 'hari@abc.com');


-- INSERTION QUERIES FOR orders TABLE
INSERT INTO orders(
    order_id, customer_id, order_date, order_total) 
    VALUES(1, 103, DATE '2023-04-12', 3000.00);

INSERT INTO orders(
    order_id, customer_id, order_date, order_total) 
    VALUES(2, 101, DATE '2023-04-10', 500.00);

INSERT INTO orders(
    order_id, customer_id, order_date, order_total) 
    VALUES(3, 105, DATE '2023-04-14', 25000.00);

INSERT INTO orders(
    order_id, customer_id, order_date, order_total) 
    VALUES(4, 100, DATE '2023-04-12', 9000.00);


-- INSERTION QUERIES FOR order-items TABLE
INSERT INTO order_items(
    item_id, order_id, product_id, quantity, price) 
    VALUES(1, 4, 14, 1, 8000.00);

INSERT INTO order_items(
    item_id, order_id, product_id, quantity, price) 
    VALUES(2, 4, 15, 1, 1000.00);

INSERT INTO order_items(
    item_id, order_id, product_id, quantity, price) 
    VALUES(3, 3, 12, 1, 20000.00);

INSERT INTO order_items(
    item_id, order_id, product_id, quantity, price) 
    VALUES(4, 3, 13, 2, 2500.00);

INSERT INTO order_items(
    item_id, order_id, product_id, quantity, price) 
    VALUES(5, 2, 11, 1, 500.00);

INSERT INTO order_items(
    item_id, order_id, product_id, quantity, price) 
    VALUES(6, 1, 11, 1, 500.00);

INSERT INTO order_items(
    item_id, order_id, product_id, quantity, price) 
    VALUES(7, 1, 13, 1, 2500.00);


-- INSERTION QUERIES FOR order-items TABLE
INSERT INTO products(
    product_id, product_name, product_category, price) 
    VALUES(11, 'Black & blue t-shirt', 'Clothing', 500.00);

INSERT INTO products(
    product_id, product_name, product_category, price) 
    VALUES(12, 'XYZ model 5 smartphone', 'Mobiles', 20000.00);

INSERT INTO products(
    product_id, product_name, product_category, price) 
    VALUES(13, 'Sound MX headphone', 'Headphones', 2500.00);

INSERT INTO products(
    product_id, product_name, product_category, price) 
    VALUES(14, 'X mini smartphone', 'Mobiles', 8000.00);

INSERT INTO products(
    product_id, product_name, product_category, price) 
    VALUES(15, 'Toughmen blue jeans', 'Clothing', 1000.00);

INSERT INTO products(
    product_id, product_name, product_category, price) 
    VALUES(16, 'Power forearms builder', 'Fitness', 300.00);