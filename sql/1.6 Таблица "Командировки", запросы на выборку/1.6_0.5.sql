SELECT name, city, date_first, date_last 
FROM trip
WHERE DATEDIFF(date_last, date_first) IN (
        SELECT MIN(DATEDIFF(date_last, date_first))
        FROM trip
    );