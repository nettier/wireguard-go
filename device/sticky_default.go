//go:build !linux

package device

import (
	"github.com/nettier/wireguard/conn"
	"github.com/nettier/wireguard/rwcancel"
)

func (device *Device) startRouteListener(_ conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
