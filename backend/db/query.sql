-- name: CheckExists :one
SELECT id FROM vouchers 
WHERE flight_number = sqlc.arg('flight_number') AND flight_date = sqlc.arg('flight_date') LIMIT 1;

-- name: InsertVoucher :exec
INSERT INTO vouchers (
    crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3, created_at
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
);