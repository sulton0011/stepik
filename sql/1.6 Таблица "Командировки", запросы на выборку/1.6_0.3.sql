SELECT city, COUNT(city) AS Количество
FROM trip
GROUP BY city
ORDER BY 2 DESC
LIMIT 2;