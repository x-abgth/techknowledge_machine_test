SELECT c.first_name, c.last_name, c.email, COUNT(o.order_id) AS total_orders, SUM(o.order_total) AS total_order_value
FROM customers c
LEFT JOIN orders o ON c.customer_id = o.customer_id
GROUP BY c.customer_id
ORDER BY total_order_value DESC;