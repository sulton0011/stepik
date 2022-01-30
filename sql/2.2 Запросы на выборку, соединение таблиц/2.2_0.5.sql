SELECT name_author
FROM 
    author JOIN book using(author_id)
GROUP BY name_author
HAVING COUNT(DISTINCT genre_id) = 1
ORDER BY name_author;