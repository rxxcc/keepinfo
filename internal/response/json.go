package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		_, _ = fmt.Fprintf(w, "%s", err.Error())
	}
}
