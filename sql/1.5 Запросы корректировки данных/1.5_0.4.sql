UPDATE book
SET price = price * .90
WHERE amount BETWEEN 5 AND 10;

SELECT * FROM book;