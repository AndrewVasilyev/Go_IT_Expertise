package storage

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h DbHandler) DeleteWorkplace(w http.ResponseWriter, r *http.Request) {

	//params := mux.Vars(r)

	//ip, _ := strconv.Atoi(params["ip"])

	ip := r.URL.Query().Get("ip")

	//var workplace models.WorkplaceModelDB

	if result := h.DB.Exec(fmt.Sprintf("DELETE data FROM workplace_model_dbs WHERE data -> 'ip'='%s'", ip)); result.Error != nil {

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Defined workplace was not found"})

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"data": "Defined workplace was successfuly deleted"})

}
