// +build linux darwin dragonfly freebsd netbsd openbsd

package greuse

import (
	"golang.org/x/sys/unix"
	"syscall"
)

func init() {
	Enabled = true
}

// See net.RawConn.Control
func Control(network, address string, c syscall.RawConn) (err error) {
	e := c.Control(func(fd uintptr) {
		// SO_REUSEADDR
		if err = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1); err != nil {
			panic(err)
		}
		// SO_REUSEPORT
		if err = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1); err != nil {
			panic(err)
		}
	})
	if e != nil {
		return e
	}
	return
}
