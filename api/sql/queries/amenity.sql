-- name: GetAllAmenities :many
SELECT
    *
FROM
    public.amenity;

-- name: CreateAmenity :many
INSERT INTO
    public.amenity (name)
SELECT
    unnest(@names::INT[]) AS name
RETURNING
    *;