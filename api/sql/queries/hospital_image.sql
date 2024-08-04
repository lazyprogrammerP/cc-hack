-- name: CreateHospitalImageWithHospitalID :many
INSERT INTO
    public.hospital_image (src, hospital_id)
SELECT
    unnest(@srcs::VARCHAR(255)[]) AS src,
    unnest(@hospital_ids::INT[]) AS hospital_id
RETURNING
    *;