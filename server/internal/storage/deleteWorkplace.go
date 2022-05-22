package storage

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h DbHandler) DeleteWorkplace(w http.ResponseWriter, r *http.Request) {

	ip := r.URL.Query().Get("ip")

	if result := h.DB.Debug().Exec(fmt.Sprintf("DELETE FROM public.workplace_model_dbs WHERE data ->> 'ip' = '%s';", ip)); result.Error != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Defined workplace was not found"})

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"data": "Defined workplace was successfuly deleted"})

}
