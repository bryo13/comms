package pkg

import (
	"net"
	"time"
)

func CheckInternet() bool {
	timeout := 2 * time.Second
	_, err := net.DialTimeout("tcp", "www.google.com:80", timeout)
	return err == nil
}

// Write or ammend friendlist
// get IP
// listen on the connection
// write on the connection
