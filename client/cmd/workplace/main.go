package workplace

import (
	"flag"

	"github.com/AndrewVasilyev/Go_IT_Expertise/client/pkg"
)

func main() {

	var wrkPlcInf WorkplaceInfo.WorkplaceInfo
	var empty WorkplaceInfo.WorkplaceInfo
	var hostname string
	var err error

	hostnameKey := flag.String("hostname", "", "describes the name of computer")
	ipAddrKey := flag.String("ip", "", "describes network address of current computer")
	currUsernameKey := flag.String("username", "", "describes current username of this computer")

	wrkPlcInf.Hostname = *hostnameKey
	wrkPlcInf.NetworkAddr = *ipAddrKey
	wrkPlcInf.CurrUsername = *currUsernameKey

	if wrkPlcInf == empty {

		hostname, err = pkg.GetHostName()

		if err != nil {
			panic(err)
		}
		wrkPlcInf.Hostname = hostname
	}

}
