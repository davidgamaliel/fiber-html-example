package domain

import (
	"time"

	"github.com/fiber-html/internal/helper/enum"
)

type MerchantData struct {
	ID        uint64              `db:"id"`
	Name      string              `db:"name"`
	StoreName string              `db:"store_name"`
	Status    enum.MerchantStatus `db:"status"`
	PhoneNo   string              `db:"phone_no"`
	JoinedAt  time.Time           `db:"joined_at"`
	CreatedAt time.Time           `db:"created_at"`
	UpdatedAt time.Time           `db:"Updated_at"`
}

type MerchantDataResponse struct {
	ID        uint64    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	StoreName string    `json:"store_name" db:"store_name"`
	Status    string    `json:"status" db:"status"`
	PhoneNo   string    `json:"phone_no" db:"phone_no"`
	JoinedAt  time.Time `json:"joined_at" db:"joined_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"Updated_at" db:"Updated_at"`
}

type MerchantListRequest struct {
	ID        uint64 `json:"id" form:"id" db:"id"`
	Name      string `json:"name" form:"name" db:"name"`
	StoreName string `json:"store_name" form:"store_name" db:"store_name"`
	Status    string `json:"status" form:"status" db:"status"`
	Limit     int    `json:"limit" form:"limit"`
	Offset    int    `json:"offset" form:"offset"`
}
