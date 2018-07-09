package restkafka

import (
	"net/http"
)

//RestToKafka ...
func RestToKafka(w http.ResponseWriter, r *http.Response) {
	w.Header().Set("Content-Type", "application/json")

	var attempt model.Attempted

}
