package app

import (
	"net/http"
	"time"

	"github.com/DeKal/bookstore_items-api/pkg/controllers"
	"github.com/gorilla/mux"
)

// StartApplication start application with server and router
func StartApplication() {
	router := mux.NewRouter()
	itemsController := controllers.NewItemsController()
	mapUrls(router, itemsController)

	// Testing api with ping controller
	pingController := controllers.NewPingController()
	router.HandleFunc("/ping", pingController.Ping).Methods(http.MethodGet)

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:9003",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
