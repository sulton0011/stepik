SELECT name_genre, title, name_author
FROM genre g 
    INNER JOIN book b using(genre_id)
    INNER JOIN author a using(author_id)
WHERE name_genre = 'Роман'
ORDER BY title;