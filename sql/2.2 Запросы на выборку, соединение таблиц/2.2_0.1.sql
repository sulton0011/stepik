SELECT g.name_genre
FROM genre g LEFT JOIN book b
ON g.genre_id = b.genre_id
WHERE b.genre_id IS Null;