-- name: GetAllDepartments :many
SELECT
    *
FROM
    public.department;

-- name: CreateDepartment :many
INSERT INTO
    public.department (name)
SELECT
    unnest(@names::VARCHAR(255)[]) AS name
RETURNING
    *;