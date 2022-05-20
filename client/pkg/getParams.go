package pkg

import (
	"net"
	"os"
	"os/user"
)

func GetHostName() (hostname string) {

	hostname, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	return hostname
}

func GetIPAddress() (ipAddr string) {

	ifaces, err := net.Interfaces()

	if err != nil {
		panic(err)
	}

	for _, i := range ifaces {

		addrs, err := i.Addrs()

		if err != nil {
			panic(err)
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

	return ipAddr
}

func GetCurrUsername() (currUser string) {

	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	currUser = user.Username

	return currUser
}
