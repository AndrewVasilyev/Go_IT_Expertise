package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AndrewVasilyev/Go_IT_Expertise/client/internal/models"
	WorkplaceInfo "github.com/AndrewVasilyev/Go_IT_Expertise/client/internal/params"
	"github.com/AndrewVasilyev/Go_IT_Expertise/client/pkg"
)

func main() {

	var serverAddr = "http://localhost:8080"

	var wrkPlcInf WorkplaceInfo.WorkplaceInfo
	var empty WorkplaceInfo.WorkplaceInfo

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

		client := &http.Client{}

		request, err := http.NewRequest(http.MethodGet, serverAddr+"/workplace", bytes.NewBufferString(""))

		query := request.URL.Query()
		query.Add("ip", wrkPlcInf.NetworkAddr)

		request.URL.RawQuery = query.Encode()

		if err != nil {
			log.Fatal(err)
		}

		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		resp, err := client.Do(request)

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		var workplace models.WorkplaceModelDB

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(respBody, &workplace)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode == 200 {
			log.Printf("Hostname: %s, IP Address: %s, Username: %s", workplace.Data.Hostname, workplace.Data.IPAddr, workplace.Data.Username)
		} else {
			log.Printf("Desired action was not implemented. %s", resp.Body)
		}
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

		request, err := http.NewRequest(http.MethodPut, serverAddr+"/workplace", bytes.NewBuffer(reqBody))

		query := request.URL.Query()
		query.Add("ip", wrkPlcInf.NetworkAddr)

		request.URL.RawQuery = query.Encode()

		if err != nil {
			log.Fatal(err)
		}

		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		resp, err := client.Do(request)

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			log.Printf("Desired action done. Status code: %d", resp.StatusCode)
		} else {
			log.Printf("Desired action was not implemented. %s", resp.Body)
		}
		break

	case "d":
		if wrkPlcInf.NetworkAddr == "" {
			log.Println("Can't delete user. IP address not specified.")
		}

		client := &http.Client{}

		request, err := http.NewRequest(http.MethodDelete, serverAddr+"/workplace", bytes.NewBufferString(""))

		query := request.URL.Query()
		query.Add("ip", wrkPlcInf.NetworkAddr)

		request.URL.RawQuery = query.Encode()

		if err != nil {
			log.Fatal(err)
		}

		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		resp, err := client.Do(request)

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			log.Printf("Desired action done. Status code: %d", resp.StatusCode)
		} else {
			log.Printf("Desired action was not implemented. %s", resp.Body)
		}

		break
	default:
	}

}
