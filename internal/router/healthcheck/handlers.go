package healthcheck

import "net/http"

func getStatus(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
