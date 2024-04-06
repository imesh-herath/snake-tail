package router

import (
	// "fmt"
	// "net/http"
	// "snake-tail/config"
	"snake-tail/http/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	// cfg := config.App
	router := mux.NewRouter()

	// Set up routes
	router.HandleFunc("/snakes/img", controllers.HandleImagePrediction).Methods("POST")

	router.HandleFunc("/snakes", controllers.GetSnakes).Methods("GET")
	router.HandleFunc("/snakes/{id}", controllers.GetSnakeByID).Methods("GET")
	router.HandleFunc("/getHeadShapes", controllers.GetUniqueHeadShapes).Methods("GET")
	router.HandleFunc("/getSkinColor", controllers.GetUniqueSkinColor).Methods("GET")
	router.HandleFunc("/getSkinPattern", controllers.GetUniqueSkinPattern).Methods("GET")

	router.HandleFunc("/patient", controllers.CreatePatient).Methods("POST")
	router.HandleFunc("/snakes/spec", controllers.GetSnakeFromSpec).Methods("POST")

	// Start the server
	// fmt.Println("Server listening on port 8080...")
	// http.ListenAndServe(":" + cfg.Server.Port, router)
	// Create a new mux.Router
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),                            // Allow requests from any origin
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), // specify allowed methods
		handlers.AllowedHeaders([]string{"Content-Type"}),                 // specify allowed headers
	)

	// Apply CORS middleware to the router
	routerWithCORS := corsMiddleware(router)

	// Create a new mux.Router
	finalRouter := mux.NewRouter()

	// Register routes from routerWithCORS to finalRouter
	finalRouter.PathPrefix("/").Handler(routerWithCORS)

	return finalRouter
}
