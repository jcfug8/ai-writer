package replies

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Write -
func Write(w http.ResponseWriter, httpStatus int, reply interface{}) {
	br, err := json.Marshal(reply)
	if err != nil {
		log.Panicf("unable to marshal reply: %s", err)
	}
	log.Infof("about to write status: %d - body: %s ", httpStatus, br)
	w.WriteHeader(httpStatus)
	w.Write(br)
}
