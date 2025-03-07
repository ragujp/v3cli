package v3

import (
	"time"

	"gorm.io/gorm"
)

type SpeedtestSession struct {
	UUID            string         `gorm:"type:uuid;primaryKey" json:"uuid"`
	CreatedAt       time.Time      `json:"CreatedAt"`              // gorm.Model default
	UpdatedAt       time.Time      `json:"UpdatedAt"`              // gorm.Model default
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"DeletedAt"` // gorm.Model default
	DeviceId        string         `json:"deviceId"`
	UserAgent       string         `json:"userAgent"`
	IPv4Addr        string         `json:"ipv4Addr"`
	IPv6Addr        string         `json:"ipv6Addr"`
	InfoIPv4Org     *string        `json:"infoIPv4Org"`
	InfoIPv4AS      *string        `json:"infoIPv4AS"`
	InfoIPv4Country *string        `json:"infoIPv4Country"`
	InfoIPv4City    *string        `json:"infoIPv4City"`
	InfoIPv4Route   *string        `json:"infoIPv4Route"`
	InfoIPv6Org     *string        `json:"infoIPv6Org"`
	InfoIPv6AS      *string        `json:"infoIPv6AS"`
	InfoIPv6Country *string        `json:"infoIPv6Country"`
	InfoIPv6City    *string        `json:"infoIPv6City"`
	InfoIPv6Route   *string        `json:"infoIPv6Route"`
	SpeedIPv4Id     *string        `json:"speedIPv4Id"`
	SpeedIPv6Id     *string        `json:"speedIPv6Id"`
	Finished        bool           `json:"finished"`
	PreferIPv6      bool           `json:"preferIPv6"`
	OrgId           *string        `json:"orgId"`
	FreeTag         *string        `json:"freeTag"`
}

// セッションを終了させるためのリクエストボディ
type FinishSessionRequest struct {
	UUID        string  `json:"uuid"`
	DeviceId    string  `json:"deviceId"`
	SpeedIPv4Id *string `json:"speedIPv4Id"`
	SpeedIPv6Id *string `json:"speedIPv6Id"`
}
