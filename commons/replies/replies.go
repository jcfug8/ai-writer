package replies

import (
	"encoding/json"
	"log"
	"net/http"
)

// Write -
func Write(w http.ResponseWriter, httpStatus int, reply interface{}) {
	br, err := json.Marshal(reply)
	if err != nil {
		log.Panicf("unable to marshal reply: %s", err)
	}

	w.WriteHeader(httpStatus)
	w.Write(br)
}
