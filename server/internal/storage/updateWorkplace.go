package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AndrewVasilyev/Go_IT_Expertise/server/internal/models"
)

func (h DbHandler) UpdateWorkplace(w http.ResponseWriter, r *http.Request) {

	ip := r.URL.Query().Get("ip")

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	}

	var workplace models.WorkplaceModel

	err = json.Unmarshal(body, &workplace)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Bad Request"})

		return
	}

	if result := h.DB.Debug().Exec(fmt.Sprintf("UPDATE public.workplace_model_dbs set data = data || jsonb_build_object('hostname', '%s', 'username', '%s') WHERE data ->> 'ip'='%s';", workplace.Hostname, workplace.Username, ip)); result.Error != nil {

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(w).Encode(map[string]string{"error": "Record Can Not Be Updated"})

		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"data": "Record Updated Successfuly"})

}
