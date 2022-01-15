SELECT title, author, price, amount, price * amount AS total_cost,
IF(amount <= 2, 'скидка 10%', IF(amount <= 3, 'скидка 20%', 'скидка 30%')) AS sale_size,
ROUND(IF(amount <= 2, price * amount * 0.1, IF(amount <= 3, price * amount * 0.2, price * amount * 0.3)), 2) AS sale,
price * amount - ROUND(IF(amount <= 2, price * amount * 0.1, IF(amount <= 3, price * amount * 0.2, price * amount * 0.3)), 2) AS total_cost_with_sale
FROM book
WHERE title LIKE '% %' AND ROUND(price) % 10 != 0
ORDER BY sale_size DESC