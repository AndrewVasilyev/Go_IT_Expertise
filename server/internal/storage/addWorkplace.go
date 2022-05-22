package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AndrewVasilyev/Go_IT_Expertise/server/internal/models"
)

func (h DbHandler) AddWorkplace(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	}
	var workplace models.WorkplaceModel

	json.Unmarshal(body, &workplace)

	if result := h.DB.Create(&workplace); result.Error != nil {

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(w).Encode(map[string]string{"error": "Record Can Not Be Created"})

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"data": "Record Created Successfuly"})

}
