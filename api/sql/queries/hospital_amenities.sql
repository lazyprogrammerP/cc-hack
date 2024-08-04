-- name: CreateHospitalAmenities :many
INSERT INTO
    public.hospital_amenities (amenity_id, hospital_id)
SELECT
    unnest(@amenity_ids::INT[]) AS amenity_id,
    unnest(@hospital_ids::INT[]) AS hospital_id
RETURNING
    *;