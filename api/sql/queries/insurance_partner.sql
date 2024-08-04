-- name: GetAllInsurancePartners :many
SELECT
    *
FROM
    public.insurance_partner;

-- name: CreateInsurancePartner :many
INSERT INTO
    public.insurance_partner (name)
SELECT
    unnest(@names::INT[]) AS name
RETURNING
    *;