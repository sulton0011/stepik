UPDATE fine f, payment p
SET f.date_payment = p.date_payment,
    f.sum_fine = IF(DATEDIFF(f.date_payment , f.date_violation)  <= 20, f.sum_fine/2, f.sum_fine)
WHERE 
    (f.name, f.number_plate, f.violation) = (p.name, p.number_plate, p.violation)
    AND f.date_payment IS NULL;

SELECT * FROM fine;