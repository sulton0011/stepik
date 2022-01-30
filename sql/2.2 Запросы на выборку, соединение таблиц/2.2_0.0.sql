SELECT b.title, g.name_genre, b.price
FROM book b INNER JOIN  genre g 
ON b.genre_id = g.genre_id
WHERE b.amount > 8
ORDER BY price DESC;