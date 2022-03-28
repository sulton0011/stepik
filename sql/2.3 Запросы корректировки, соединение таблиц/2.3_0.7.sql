insert into author (name_author)
select author from supply 
where 
    author not in (select name_author from author);
select * from author;
insert into book (title,author_id,genre_id,price,amount)
select 
    title,
    author_id,
    (select genre_id from genre where name_genre="Роман"),
    price,
    amount
from supply inner join author on author=name_author where 
    title not in (select title from book) and author_id not in (select author_id from book);
select * from book;