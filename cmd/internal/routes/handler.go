package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger" // Import the missing package
)

type Handler struct {
	// Add any necessary fields here
}

func (h Handler) InitRouter() *mux.Router {
	// Add your code here

	r := mux.NewRouter()
	r.HandleFunc("/health", h.HealthHandler).Methods("GET")
	r.PathPrefix("/docs").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // Replace with your OpenAPI JSON endpoint
	))
	return r
}

func (h Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
