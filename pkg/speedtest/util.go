//go:build !linux
// +build !linux

package speedtest

import (
	"fmt"
	"net"
)

// How to ignore nusedparams error
func NewDialerInterfaceBound(iface string) (dialer *net.Dialer, err error) {
	_ = iface
	return nil, fmt.Errorf("cannot bound to interface on this platform")
}
