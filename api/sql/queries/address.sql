-- name: CreateAddress :one
INSERT INTO
    public.address (street, landmark, city, pincode, hospital_id)
VALUES
    ($1, $2, $3, $4, $5)
RETURNING
    *;