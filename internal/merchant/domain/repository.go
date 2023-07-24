package domain

import (
	"context"
	"fmt"

	"github.com/fiber-html/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type MerchantRepository struct {
	DB     *pgxpool.Pool
	Logger config.Logger
}

type MerchantPersister interface {
	GetMerchants(req MerchantListRequest) ([]MerchantData, error)
	GetMerchantByID(id int) (*MerchantData, error)
}

func NewRepository(db *pgxpool.Pool, logger config.Logger) MerchantPersister {
	return &MerchantRepository{
		DB:     db,
		Logger: logger,
	}
}

func (r *MerchantRepository) GetMerchants(req MerchantListRequest) ([]MerchantData, error) {
	query := `SELECT id, name, store_name, joined_at, status, phone_no
						FROM merchants
						WHERE 1=1`
	whereQuery, values := buildWhereMerchantPageQuery(req)
	r.Logger.Info(">> DEBUG ", zap.Any("where", query+whereQuery))
	r.Logger.Info(">> DEBUG ", zap.Any("values", values))

	rows, err := r.DB.Query(context.Background(), query+whereQuery, values...)
	if err != nil {
		r.Logger.Error("failed to query repo", zap.Error(err))
		return []MerchantData{}, err
	}
	defer rows.Close()

	var res []MerchantData
	for rows.Next() {
		var merchant MerchantData
		err = rows.Scan(
			&merchant.ID,
			&merchant.Name,
			&merchant.StoreName,
			&merchant.JoinedAt,
			&merchant.Status,
			&merchant.PhoneNo,
		)
		if err != nil {
			r.Logger.Error("failed to scan merchant", zap.Error(err))
			return res, err
		}

		res = append(res, merchant)
	}

	r.Logger.Info("DEBUG", zap.Any("merchants", res))
	return res, nil
}

func buildWhereMerchantPageQuery(req MerchantListRequest) (string, []interface{}) {
	var where string
	var values []interface{}
	counter := 0

	if req.ID != 0 {
		counter++
		where += fmt.Sprintf(" AND id = $%d", counter)
		values = append(values, req.ID)
	}
	if req.Name != "" {
		counter++
		where += fmt.Sprintf(" AND name = $%d", counter)
		values = append(values, req.Name)
	}
	if req.StoreName != "" {
		counter++
		where += fmt.Sprintf(" AND store_name = $%d", counter)
		values = append(values, req.StoreName)
	}
	if req.Status != "" {
		counter++
		where += fmt.Sprintf(" AND status = $%d", counter)
		values = append(values, req.Status)
	}
	if req.Limit == 0 {
		req.Limit = 10
	}

	where += fmt.Sprintf(" LIMIT $%d OFFSET $%d", counter+1, counter+2)
	values = append(values, req.Limit, req.Offset)

	return where, values
}

func (r *MerchantRepository) GetMerchantByID(id int) (*MerchantData, error) {
	var merchant MerchantData
	query := "select id, name, store_name, phone_no, joined_at, status from merchants where id = $1"
	row := r.DB.QueryRow(context.Background(), query, id)
	err := row.Scan(
		&merchant.ID,
		&merchant.Name,
		&merchant.StoreName,
		&merchant.PhoneNo,
		&merchant.JoinedAt,
		&merchant.Status,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	} else if err != nil {
		r.Logger.Error("GetMerchantByID - failed to scan", zap.Error(err))
		return nil, err
	}

	return &merchant, nil
}
