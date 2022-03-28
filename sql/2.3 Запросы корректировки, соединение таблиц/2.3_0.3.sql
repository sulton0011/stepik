UPDATE book
SET genre_id = 
    (
        SELECT genre_id
        FROM genre
        WHERE name_genre = "Поэзия"
    )
WHERE book_id in (10);

UPDATE book
SET genre_id = 
    (
        SELECT genre_id
        FROM genre
        WHERE name_genre = "Приключения"
    )
WHERE book_id in (11);

SELECT * FROM book;