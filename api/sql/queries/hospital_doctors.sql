-- name: CreateHospitalDoctors :many
INSERT INTO
    public.hospital_doctors (doctor_id, hospital_id)
SELECT
    unnest(@doctor_ids::INT[]) AS doctor_id,
    unnest(@hospital_ids::INT[]) AS hospital_id
RETURNING
    *;