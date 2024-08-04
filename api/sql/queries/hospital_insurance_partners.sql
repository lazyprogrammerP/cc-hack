-- name: CreateHospitalInsurancePartners :many
INSERT INTO
    public.hospital_insurance_partners (insurance_partner_id, hospital_id)
SELECT
    unnest(@insurance_partner_ids::INT[]) AS insurance_partner_id,
    unnest(@hospital_ids::INT[]) AS hospital_id
RETURNING
    *;