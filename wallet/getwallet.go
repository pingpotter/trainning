package wallet

import (
	"net/http"
	"mux"
)

//petch

func getwallet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w_id := vars["walletId"]
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write([]byte("{\"result\":\"OK\"}"+w_id))
}