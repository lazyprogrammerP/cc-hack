-- name: GetHospitalByID :one
-- SELECT
--     h.id AS hospital_id,
--     h.name AS hospital_name,
--     h.lat,
--     h.long,
--     h.cost,
--     h.accreditation,
--     h.wait_time,
--     h.capacity,
--     h.site_url,
--     h.contact_number,
--     h.ratings,
--     ARRAY_AGG(JSON_BUILD_OBJECT('id', hi.id, 'src', hi.src)) AS images,
--     JSON_BUILD_OBJECT('id', a.id, 'street', a.street, 'landmark', a.landmark, 'city', a.city, 'pincode', a.pincode) AS address
-- FROM
--     hospital h
--     LEFT JOIN hospital_image hi ON h.id = hi.hospital_id
--     LEFT JOIN address a ON h.id = a.hospital_id
-- WHERE
--     h.id = $1
-- GROUP BY
--     h.id,
--     a.id;
-- name: CreateHospital :one
INSERT INTO
    public.hospital (name, lat, long, cost, accreditation, wait_time, capacity, site_url, contact_number, ratings)
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING
    *;

-- name: DeleteHospitalByID :exec
DELETE FROM public.hospital
WHERE
    id = $1;