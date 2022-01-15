SELECT title, author, price, amount FROM book
WHERE (price > 600 OR price < 500) AND price * amount >= 5000;