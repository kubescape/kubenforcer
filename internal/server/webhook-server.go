package server

import (
	"fmt"
	"net/http"

	"github.com/armosec/admission-controler/internal/server/handlers"
)

func NewServer(port string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/armo-admission-controler/validating", handlers.AdmissionControlerHandler)
	mux.HandleFunc("/armo-admission-controler/mutating", handlers.AdmissionControlerHandler)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
}
