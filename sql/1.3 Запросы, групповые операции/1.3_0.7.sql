SELECT author,
SUM(price * amount) AS Стоимость
FROM book
WHERE title not IN ('Идиот', 'Белая_гвардия')
GROUP BY author
HAVING SUM(price * amount) > 5000
ORDER BY 2 DESC;