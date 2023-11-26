package check

import (
	"fmt"
	"net"
	"time"
)

func Check(domain string, port string) string {
	address := domain + ":" + port
	timeout := time.Duration(5 * time.Second)
	connection, err := net.DialTimeout("tcp", address, timeout)

	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable , \n Error : %v", domain, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable , \n From : %v \n To : %v", domain, connection.LocalAddr(), connection.RemoteAddr())
	}
	return status
}
