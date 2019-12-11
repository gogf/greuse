// +build windows

package greuse

import (
	"golang.org/x/sys/windows"
	"syscall"
)

// See net.RawConn.Control
func Control(network, address string, c syscall.RawConn) (err error) {
	e := c.Control(func(fd uintptr) {
		if err = windows.SetsockoptInt(windows.Handle(fd), windows.SOL_SOCKET, windows.SO_REUSEADDR, 1); err != nil {
			return
		}
	})
	if e != nil {
		return e
	}
	return
}
