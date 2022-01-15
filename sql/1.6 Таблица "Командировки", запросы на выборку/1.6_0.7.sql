SELECT 
    MONTHNAME(date_first) AS Месяц,
    COUNT(MONTH(date_first)) AS Количество
FROM trip
GROUP BY Месяц
ORDER BY 2 DESC, 1;