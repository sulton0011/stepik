select 
    title, 
    author, 
    price, 
    amount, 
    round((SELECT avg(amount) FROM book) - amount, 2) AS Средний_заказ
from book
where author in (
    select author
    from book
    group by author
    having count(author) < 3
    )