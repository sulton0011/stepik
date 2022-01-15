SELECT title,
    author,
    amount,
    ROUND(price - price / 100 * 30, 2) AS new_price
FROM book;