package tablets

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Golang-labs-ip/Golang-lab3/server/tools"
)

// HTTPHandlerFunc is Balancers HTTP handler.
type HTTPHandlerFunc http.HandlerFunc

// HTTPHandler creates a new instance of Balancers HTTP handler.
func HTTPHandler(store *Store) HTTPHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListTableta(store, rw)
		} else if r.Method == "POST" {
			handleDeviceUpdate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleListTableta(store *Store, rw http.ResponseWriter) {
	res, err := store.ListOfTablets()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJSONInternalError(rw)
		return
	}
	tools.WriteJSONOk(rw, res)
}

func handleDeviceUpdate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var dev UpdateDev
	if err := json.NewDecoder(r.Body).Decode(&dev); err != nil {
		log.Printf("Error decoding machine input: %s", err)
		tools.WriteJSONBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.UpdateDevice(dev.ID, dev.Battery, dev.CurrentVideo, dev.DeviceTime)
	if err == nil {
		tools.WriteJSONOk(rw, &dev)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJSONInternalError(rw)
	}
}
