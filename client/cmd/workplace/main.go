package workplace

import (
	"flag"
)

func main() {

	var wrkPlcInf workplaceInfo

	hostnameKey := flag.String("hostname", "", "describes the name of computer")
	ipAddrKey := flag.String("ip", "", "describes network address of current computer")
	currUsernameKey := flag.String("username", "", "describes current username of this computer")

}
