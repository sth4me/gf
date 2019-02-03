// +build windows

package greuseport

import (
    "syscall"
    "github.com/gogf/gf/third/golang.org/x/sys/windows"
)

// See net.RawConn.Control
func Control(network, address string, c syscall.RawConn) (err error) {
	return c.Control(func(fd uintptr) {
        if err = windows.SetsockoptInt(windows.Handle(fd), windows.SOL_SOCKET, windows.SO_REUSEADDR, 1); err != nil {
            return
        }
	})
}
