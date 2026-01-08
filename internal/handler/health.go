package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
)

// HealthHandler checks API and DB status
func HealthHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.PingContext(context.Background()); err != nil {
			http.Error(w, `{"status":"db error"}`, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	}
}