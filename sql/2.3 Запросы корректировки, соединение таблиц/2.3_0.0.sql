UPDATE book b
    INNER JOIN author a ON a.author_id = b.author_id
    INNER JOIN supply s ON b.title = s.title
        AND s.author = a.name_author
SET
    b.amount = b.amount + s.amount,
    s.amount = 0,
    b.price = (b.price * b.amount + s.price  * s.amount)/(b.amount + s.amount)
WHERE b.price <> s.price;

select * from book;
select * from supply;