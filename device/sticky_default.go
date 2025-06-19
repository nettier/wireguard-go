//go:build !linux

package device

import (
	"github.com/nettier/wireguard-go/conn"
	"github.com/nettier/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(_ conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
