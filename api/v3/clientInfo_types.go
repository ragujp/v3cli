package v3

import (
	"net"

	"github.com/ipinfo/go/v2/ipinfo"
)

type ClientInfo struct {
	IP     net.IP      `json:"ip"`
	Port   int         `json:"port"`
	UnixTime   int64   `json:"unixtime"`
	IsIPv4 bool        `json:"isIPv4"`
	IPInfo ipinfo.Core `json:"ipInfo"`
}
