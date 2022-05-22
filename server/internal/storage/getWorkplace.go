package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AndrewVasilyev/Go_IT_Expertise/server/internal/models"
)

func (h DbHandler) GetWorkplace(w http.ResponseWriter, r *http.Request) {

	ip := r.URL.Query().Get("ip")

	var workplace models.WorkplaceModelDB

	result := h.DB.Debug().Exec(fmt.Sprintf("SELECT data FROM public.workplace_model_dbs WHERE data ->> 'ip' = '%s';", ip))

	if result.Error != nil {

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(w).Encode(map[string]string{"error": "No such record"})

		return
	}

	rows, err := h.DB.Model(&models.WorkplaceModelDB{}).Exec(fmt.Sprintf("SELECT data FROM public.workplace_model_dbs WHERE data ->> 'ip' = '%s';", ip)).Rows()

	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		h.DB.ScanRows(rows, &workplace)
	}

	// err := result.Row().Scan(&workplace)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(workplace)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&workplace)

}
