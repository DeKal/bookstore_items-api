package app

import (
	"net/http"

	"github.com/DeKal/bookstore_items-api/pkg/controllers"
	"github.com/gorilla/mux"
)

func mapUrls(router *mux.Router, controller controllers.ItemsControllerInterface) {
	router.HandleFunc("/items", controller.Create).Methods(http.MethodPost)
}
