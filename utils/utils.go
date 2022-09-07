package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func PaginationDetails(limit int, page int, total int, pervious_page int, next_page int) map[string]interface{} {
	return map[string]interface{}{"page_number": page, "total_page": total, "limit": limit, "pervious_page": pervious_page, "next_page": next_page, "status": true, "message": "success"}
}
