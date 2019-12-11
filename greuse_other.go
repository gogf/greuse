// +build !windows,!linux,!darwin,!dragonfly,!freebsd,!netbsd,!openbsd

package greuse

import (
	"syscall"
)

// See net.RawConn.Control
func Control(network, address string, c syscall.RawConn) (err error) {
	return nil
}
