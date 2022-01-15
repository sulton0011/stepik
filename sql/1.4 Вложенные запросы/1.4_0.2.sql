SELECT author, title, amount
FROM book
WHERE amount IN (SELECT amount
        FROM book
        GROUP BY amount 
        HAVING count(amount) = 1
    );