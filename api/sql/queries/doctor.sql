-- name: GetDoctorByID :one
-- SELECT
--     d.id,
--     d.name,
--     d.background,
--     d.education_details,
--     d.experience,
--     d.fees,
--     JSON_BUILD_OBJECT('id', dpt.id, 'name', dpt.name) AS department
-- FROM
--     public.doctor d
--     LEFT JOIN public.department dpt ON d.department_id = dpt.id
-- WHERE
--     id = $1;
-- name: GetAllDoctors :many
SELECT
    d.id,
    d.name,
    d.background,
    d.education_details,
    d.experience,
    d.fees,
    JSON_BUILD_OBJECT('id', dpt.id, 'name', dpt.name) AS department
FROM
    public.doctor d
    LEFT JOIN public.department dpt ON d.department_id = dpt.id;

-- name: CreateDoctor :one
INSERT INTO
    public.doctor (name, background, education_details, experience, fees, department_id)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING
    *;