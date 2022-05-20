package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	WorkplaceModel "github.com/AndrewVasilyev/Go_IT_Expertise/client/internal/models"
	WorkplaceInfo "github.com/AndrewVasilyev/Go_IT_Expertise/client/internal/params"
	"github.com/AndrewVasilyev/Go_IT_Expertise/client/pkg"
)

func main() {

	var serverAddr = "http://localhost:8080"

	var wrkPlcInf WorkplaceInfo.WorkplaceInfo
	var empty WorkplaceInfo.WorkplaceInfo

	var wrkPlcMdl WorkplaceModel.WorkplaceModel

	var hostnameKey string
	var ipAddressKey string
	var currUsernameKey string
	var actionKey string

	flag.StringVar(&hostnameKey, "hostname", "", "Describes the name of the computer.")
	flag.StringVar(&ipAddressKey, "ip", "", "Describes network address of current computer.")
	flag.StringVar(&currUsernameKey, "username", "", "Describes current username of this computer.")
	flag.StringVar(&actionKey, "action", "", "Describes action with user input. Only READ, UPDATE and DELETE are available for user. Use r for READ, u for UPDATE, d for DELETE.")

	flag.Parse()

	wrkPlcInf.Hostname = hostnameKey
	wrkPlcInf.NetworkAddr = ipAddressKey
	wrkPlcInf.CurrUsername = currUsernameKey
	wrkPlcInf.Action = actionKey

	if wrkPlcInf == empty {

		wrkPlcInf.Hostname = pkg.GetHostName()

		wrkPlcInf.NetworkAddr = pkg.GetIPAddress()

		wrkPlcInf.CurrUsername = pkg.GetCurrUsername()

		reqBody, err := json.Marshal(map[string]string{
			"hostname": wrkPlcInf.Hostname,
			"ip":       wrkPlcInf.NetworkAddr,
			"username": wrkPlcInf.CurrUsername,
		})

		if err != nil {
			log.Fatal(err)
		}

		respBody, err := http.Post(serverAddr, "application/json", bytes.NewBuffer(reqBody))

		if err != nil {
			log.Fatal(err)
		}

		defer respBody.Body.Close()

		fmt.Println(respBody.StatusCode)
	}

	switch wrkPlcInf.Action {

	case "r":
		if wrkPlcInf.NetworkAddr == "" {
			log.Println("Can't read user. IP address not specified.")
		}

		respBody, err := http.Get(serverAddr + fmt.Sprintf("/%s", wrkPlcInf.NetworkAddr))

		if err != nil {
			log.Fatal(err)
		}

		defer respBody.Body.Close()

		body, err := ioutil.ReadAll(respBody.Body)

		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(body, &wrkPlcMdl)

		fmt.Printf("Hostname: %s, IP Address: %s, Username: %s", wrkPlcMdl.Hostname, wrkPlcMdl.IPAddr, wrkPlcMdl.Username)
		break

	case "u":
		if wrkPlcInf.Hostname == "" && wrkPlcInf.NetworkAddr == "" && wrkPlcInf.CurrUsername == "" {
			log.Println("Can't update user. Not enough data provided.")
		}
		break
	case "d":
		if wrkPlcInf.NetworkAddr == "" {
			log.Println("Can't delete user. IP address not specified.")
		}
		break
	default:
	}

	// flag.CommandLine.Output().Write([]byte("\n\r"))
	// flag.CommandLine.Output().Write([]byte("Hostname:"))
	// flag.CommandLine.Output().Write([]byte(wrkPlcInf.Hostname))
	// flag.CommandLine.Output().Write([]byte("\n\r"))
	// flag.CommandLine.Output().Write([]byte("IP Address:"))
	// flag.CommandLine.Output().Write([]byte(wrkPlcInf.NetworkAddr))
	// flag.CommandLine.Output().Write([]byte("\n\r"))
	// flag.CommandLine.Output().Write([]byte("Current user:"))
	// flag.CommandLine.Output().Write([]byte(wrkPlcInf.CurrUsername))
	// flag.CommandLine.Output().Write([]byte("\n\r"))

}
