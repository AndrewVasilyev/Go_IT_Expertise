package pkg

import (
	"fmt"
	"net"
	"os"
)

func GetHostName() (hostname string, err error) {

	hostname, err = os.Hostname()

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return hostname, nil
}

func GetIPAddress() (ipAddr string, errIP error) {

	iface, err := net.Interfaces()

	if err != nil {
		fmt.Println(err)
		errIP = err
		return "", errIP
	}

	for _, i := range iface {
		addrs, err := i.Addrs()

		if err != nil {
			fmt.Println(err)
			errIP = err
			return "", errIP
		}

		for _, addr := range addrs {

			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			ipAddr = ip.String()
		}

	}

	return ipAddr, nil
}
