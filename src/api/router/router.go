package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"simpleBank/src/api/handler"
)

func SetupPayaRouter(h *handler.PayaHandler) *mux.Router {
	router := mux.NewRouter()
	var api = router.PathPrefix("/api").Subrouter()

	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	//apiV1 := api.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/sheba", h.CreatePayaRequest).Methods(http.MethodPost)
	api.HandleFunc("/sheba", h.ListPayaRequests).Methods(http.MethodGet)
	api.HandleFunc("/sheba/{id}", h.UpdatePayaRequest).Methods(http.MethodPut, http.MethodPost)

	return router
}
