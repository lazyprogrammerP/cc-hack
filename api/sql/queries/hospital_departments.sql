-- name: CreateHospitalDepartments :many
INSERT INTO
    public.hospital_departments (department_id, hospital_id)
SELECT
    unnest(@department_ids::INT[]) AS department_id,
    unnest(@hospital_ids::INT[]) AS hospital_id
RETURNING
    *;