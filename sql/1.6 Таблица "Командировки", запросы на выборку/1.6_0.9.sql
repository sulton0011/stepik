SELECT 
    name,
    SUM(datediff(date_last+1, date_first) * per_diem) AS  Сумма
FROM trip
GROUP BY name
ORDER BY name
LIMIT 2;