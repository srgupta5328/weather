package router

import "net/http"

//HealthCheck ... Health Check Endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//Callback ... Callback url
func Callback(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Callback URI"))
}
