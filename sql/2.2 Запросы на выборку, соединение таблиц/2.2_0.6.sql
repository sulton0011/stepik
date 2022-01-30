select title, name_author, name_genre, price, amount
from book 
join author using(author_id)
join genre using(genre_id)
where genre.genre_id in(    
             SELECT genre_id
             FROM book GROUP BY genre_id 
             having SUM(amount) = (
                 SELECT SUM(amount)
                 FROM book GROUP BY genre_id 
                 order by 1 desc 
                 limit 1 
             )
    )
order by 1;