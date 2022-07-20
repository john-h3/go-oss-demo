package utils

import "net"

var ip string

// GetLocalIP try to return the first non-loopback local IP of the host.
// If no non-loopback IP is found, it will return the loopback IP.
func GetLocalIP() string {
	if ip != "" {
		return ip
	}
	interfaces, _ := net.Interfaces()
	ip = "127.0.0.1"
	for _, i := range interfaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			var _ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				_ip = v.IP
			case *net.IPAddr:
				_ip = v.IP
			}
			if _ip != nil {
				_ip = _ip.To4()
			}
			if _ip != nil {
				if !_ip.IsLoopback() {
					ip = _ip.String()
					return ip
				}
			}
		}
	}
	return ip
}
