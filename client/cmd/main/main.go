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

		respBody, err := http.Post(serverAddr+"/workplace", "application/json", bytes.NewBuffer(reqBody))

		if err != nil {
			log.Fatal(err)
		}

		defer respBody.Body.Close()

		log.Printf("Adding new workplace done. Status code: %d", respBody.StatusCode)
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

		err = json.Unmarshal(body, &wrkPlcMdl)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Hostname: %s, IP Address: %s, Username: %s", wrkPlcMdl.Hostname, wrkPlcMdl.IPAddr, wrkPlcMdl.Username)
		break

	case "u":
		if wrkPlcInf.Hostname == "" && wrkPlcInf.NetworkAddr == "" && wrkPlcInf.CurrUsername == "" {
			log.Println("Can't update user. Not enough data provided.")
		}

		reqBody, err := json.Marshal(map[string]string{
			"hostname": wrkPlcInf.Hostname,
			"ip":       wrkPlcInf.NetworkAddr,
			"username": wrkPlcInf.CurrUsername,
		})

		if err != nil {
			log.Fatal(err)
		}

		client := &http.Client{}

		request, err := http.NewRequest(http.MethodPut, serverAddr+fmt.Sprintf("/%s", wrkPlcInf.NetworkAddr), bytes.NewBuffer(reqBody))

		if err != nil {
			log.Fatal(err)
		}

		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		resp, err := client.Do(request)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Desired action done. Status code: %d", resp.StatusCode)
		break

	case "d":
		if wrkPlcInf.NetworkAddr == "" {
			log.Println("Can't delete user. IP address not specified.")
		}

		client := &http.Client{}

		request, err := http.NewRequest(http.MethodDelete, serverAddr+fmt.Sprintf("/%s", wrkPlcInf.NetworkAddr), nil)

		request.URL.Query().Add("ip", wrkPlcInf.NetworkAddr)

		if err != nil {
			log.Fatal(err)
		}

		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		resp, err := client.Do(request)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Desired action done. Status code: %d", resp.StatusCode)
		break
	default:
	}

}
