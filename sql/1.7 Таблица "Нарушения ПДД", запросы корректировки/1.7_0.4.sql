UPDATE fine f, (
    SELECT name, number_plate, violation
    FROM fine
    GROUP BY name, number_plate, violation
    HAVING COUNT(2) = 2) AS nv
set sum_fine = sum_fine * 2
WHERE f.name = nv.name AND date_payment IS NULL;

SELECT * FROM fine;