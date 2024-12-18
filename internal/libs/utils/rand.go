package utils

import (
	"fmt"
	"net"
)

func GetRandOpenPort() string {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	address := listener.Addr().(*net.TCPAddr)
	return fmt.Sprint(address.Port)
}
