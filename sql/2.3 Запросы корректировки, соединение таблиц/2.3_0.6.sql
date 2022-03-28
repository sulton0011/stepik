DELETE FROM author
USING 
    author 
    INNER JOIN book ON author.author_id = book.author_id
    INNER JOIN genre ON genre.genre_id = book.genre_id
WHERE name_genre = "Поэзия";

SELECT * FROM book;
SELECT * FROM author;