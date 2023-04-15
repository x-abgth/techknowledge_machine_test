CREATE TABLE customers (customer_id INT PRIMARY KEY, first_name VARCHAR(50), last_name VARCHAR(50), email VARCHAR(100));

CREATE TABLE orders ( order_id INT PRIMARY KEY, customer_id INT, order_date DATE, order_total DECIMAL(10, 2), FOREIGN KEY (customer_id) REFERENCES customers(customer_id));

CREATE TABLE order_items ( item_id INT PRIMARY KEY, order_id INT, product_id INT, quantity INT, price DECIMAL(10, 2), FOREIGN KEY (order_id) REFERENCES orders(order_id));

CREATE TABLE products ( product_id INT PRIMARY KEY, product_name VARCHAR(100), product_category VARCHAR(50), price DECIMAL(10, 2));
