// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: insurance_partner.sql

package sqlc

import (
	"context"

	"github.com/lib/pq"
)

const createInsurancePartner = `-- name: CreateInsurancePartner :many
INSERT INTO
    public.insurance_partner (name)
SELECT
    unnest($1::INT[]) AS name
RETURNING
    id, name
`

func (q *Queries) CreateInsurancePartner(ctx context.Context, names []int32) ([]InsurancePartner, error) {
	rows, err := q.db.QueryContext(ctx, createInsurancePartner, pq.Array(names))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InsurancePartner
	for rows.Next() {
		var i InsurancePartner
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllInsurancePartners = `-- name: GetAllInsurancePartners :many
SELECT
    id, name
FROM
    public.insurance_partner
`

func (q *Queries) GetAllInsurancePartners(ctx context.Context) ([]InsurancePartner, error) {
	rows, err := q.db.QueryContext(ctx, getAllInsurancePartners)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []InsurancePartner
	for rows.Next() {
		var i InsurancePartner
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}