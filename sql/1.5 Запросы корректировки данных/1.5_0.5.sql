UPDATE book
SET price = IF(buy <= 0, price * .90, price), buy = IF (buy > 15, buy - 3, buy);

SELECT * FROM book;