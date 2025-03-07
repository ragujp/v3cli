package v3

import (
	"gorm.io/gorm"
)

type ConnectionType struct {
	gorm.Model                                // id, created_at, updated_at, deleted_at
	Name          string                      `json:"name"`
	AddressFamily ConnectionTypeAddressFamily `json:"addressFamily"`
	Mss           int                         `json:"mss"`
}

type ConnectionTypeResponse struct {
	Name          string                      `json:"name"`
	AddressFamily ConnectionTypeAddressFamily `json:"addressFamily"`
	Mss           int                         `json:"mss"`
}

type ConnectionTypeAddressFamily string

const (
	ConnectionTypeAddressFamilyIPv4 ConnectionTypeAddressFamily = "ipv4"
	ConnectionTypeAddressFamilyIPv6 ConnectionTypeAddressFamily = "ipv6"
)

type FletsRegion string

const (
	FletsRegionEast FletsRegion = "east"
	FletsRegionWest FletsRegion = "west"
	FletsRegionBoth FletsRegion = "both"
)

type AccessTypeSession struct {
	gorm.Model                             // id, created_at, updated_at, deleted_at
	SpeedtestSessionUUID string            `gorm:"type:uuid" json:"speedTestSessionUUID"`
	SpeedtestSession     *SpeedtestSession `gorm:"foreignKey:SpeedtestSessionUUID;references:UUID" json:"speedtestSession,omitempty"`
	IPv4Mss              *int              `json:"ipv4Mss,omitempty"`
	IPv6Mss              *int              `json:"ipv6Mss,omitempty"`
	IPv4ConnectionTypeID *uint             `gorm:"index" json:"ipv4ConnectionTypeID,omitempty"`
	IPv6ConnectionTypeID *uint             `gorm:"index" json:"ipv6ConnectionTypeID,omitempty"`
	IPv4ConnectionType   *ConnectionType   `gorm:"foreignKey:IPv4ConnectionTypeID" json:"ipv4ConnectionType,omitempty"`
	IPv6ConnectionType   *ConnectionType   `gorm:"foreignKey:IPv6ConnectionTypeID" json:"ipv6ConnectionType,omitempty"`
	Flets                *FletsRegion      `json:"flets,omitempty"`
}
