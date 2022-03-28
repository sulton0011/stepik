INSERT INTO book (book.title, book.author_id, book.price, book.amount)
SELECT supply.title, author.author_id, supply.price, supply.amount 
FROM supply
    LEFT JOIN book ON book_id = supply_id 
    INNER JOIN author ON supply.author = author.name_author
WHERE supply.amount <> 0;

SELECT * FROM book;