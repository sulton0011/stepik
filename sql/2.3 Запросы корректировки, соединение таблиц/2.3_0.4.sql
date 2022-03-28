DELETE FROM author
WHERE author_id IN (
        SELECT author_id FROM book group by author_id having  sum(amount)  <20
    );
SELECT * FROM book;
SELECT * FROM author;