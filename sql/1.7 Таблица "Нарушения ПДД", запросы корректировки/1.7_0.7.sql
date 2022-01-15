DELETE FROM fine
WHERE date_violation < date('2020-02-01');

SELECT * FROM fine;