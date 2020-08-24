package util

import (
	"errors"
	"net"
)

func HostIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok && ipNet.IP.IsLoopback() {
			continue
		}

		return ipNet.IP.String(), nil
	}

	return "", errors.New("HostIP not found")
}

func PrivateIP4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}
