package router

import (
	// "fmt"
	// "net/http"
	// "snake-tail/config"
	"snake-tail/http/controllers"
	"github.com/gorilla/mux"
)

func Init() *mux.Router{
	// cfg := config.App
	router := mux.NewRouter()

	// Set up routes
	router.HandleFunc("/snakes", controllers.GetSnakes).Methods("GET")
	router.HandleFunc("/snakes/{id}", controllers.GetSnakeByID).Methods("GET")
	// router.HandleFunc("/patient", controllers.CreatePatient).Methods("POST")
	router.HandleFunc("/snakes/spec", controllers.GetSnakeFromSpec).Methods("POST")

	// Start the server
	// fmt.Println("Server listening on port 8080...")
	// http.ListenAndServe(":" + cfg.Server.Port, router)

	return router
}