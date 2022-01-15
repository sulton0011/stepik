DELETE FROM supply
WHERE author IN (
    SELECT author FROM book
    GROUP BY author
    having sum(amount) >10
);

select * from supply;