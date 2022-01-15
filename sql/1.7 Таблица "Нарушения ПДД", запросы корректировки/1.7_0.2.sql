UPDATE fine f, traffic_violation tv
SET f.sum_fine = tv.sum_fine
WHERE f.violation = tv.violation AND f.sum_fine IS NULL;

select * from fine;