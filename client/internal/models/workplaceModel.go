package models

type WorkplaceModel struct {
	Hostname string `json:"hostname"`
	IPAddr   string `json:"ip"`
	Username string `json:"username"`
}
