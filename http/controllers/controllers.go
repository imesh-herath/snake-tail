package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"snake-tail/config"
	"snake-tail/domain/entities"

	// "log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// Snake struct represents the snake model
type Snake struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Color          string             `json:"color,omitempty" bson:"color,omitempty"`
	Description    string             `json:"description,omitempty" bson:"description,omitempty"`
	FirstAid       string             `json:"firstAid,omitempty" bson:"firstAid,omitempty"`
	HeadShape      string             `json:"headShape,omitempty" bson:"headShape,omitempty"`
	Image          string             `json:"image,omitempty" bson:"image,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	OtherName      string             `json:"otherName,omitempty" bson:"otherName,omitempty"`
	Pattern        string             `json:"pattern,omitempty" bson:"pattern,omitempty"`
	ScientificName string             `json:"scientificName,omitempty" bson:"scientificName,omitempty"`
	VenomLevel     string             `json:"venomLevel,omitempty" bson:"venomLevel,omitempty"`
}

var client *mongo.Client
var collection *mongo.Collection

// func main() {
// 	// Initialize MongoDB client
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	client, err := mongo.Connect(nil, clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(nil)

// 	// Initialize router
// 	router := mux.NewRouter()

// 	log.Fatal(http.ListenAndServe(":8080", router))
// }

func GetSnakes(w http.ResponseWriter, r *http.Request) {
	// Set up your Firestore API URL and API key
	firestoreURL := config.App.FirebaseSnake.Url
	apiKey := config.App.FirebaseSnake.ApiKey

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", firestoreURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add the API key as a query parameter
	q := req.URL.Query()
	q.Add("key", apiKey)
	req.URL.RawQuery = q.Encode()

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	apiResponse := &entities.SnakeResponse{}
	_ = json.Unmarshal([]byte(body), apiResponse)

	// Process the response (body) as needed
	fmt.Println("Response:", apiResponse)
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}

func GetSnakeByID(w http.ResponseWriter, r *http.Request) {
	// Get the snake ID from the request URL
	snakeID := mux.Vars(r)["id"]

	// Set up your Firestore API URL and API key
	firestoreURL := fmt.Sprintf("%s/%s", config.App.FirebaseSnake.Url, snakeID)
	apiKey := config.App.FirebaseSnake.ApiKey

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", firestoreURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add the API key as a query parameter
	q := req.URL.Query()
	q.Add("key", apiKey)
	req.URL.RawQuery = q.Encode()

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Check if the document exists
	if resp.StatusCode == http.StatusNotFound {
		http.NotFound(w, r)
		return
	}

	apiResponse := &entities.SnakeResponse{}
	_ = json.Unmarshal(body, apiResponse)

	// Process the response (body) as needed
	fmt.Println("Response:", apiResponse)
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}

// func GetSnakeByID(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id, err := primitive.ObjectIDFromHex(params["id"])
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(fmt.Sprintf("Invalid ID: %v", id)))
// 		return
// 	}

// 	var snake Snake
// 	err = collection.FindOne(nil, bson.M{"_id": id}).Decode(&snake)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("Unable to fetch data from database"))
// 		return
// 	}

// 	json.NewEncoder(w).Encode(snake)
// }

func CreatePatient(w http.ResponseWriter, r *http.Request) {
	// Set up your Firestore API URL and API key
	firestoreURL := config.App.FirebasePatient.Url
	apiKey := config.App.FirebasePatient.ApiKey

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Unmarshal the JSON request body into PatientRequest struct
	var patientRequest entities.PatientRequest
	err = json.Unmarshal(body, &patientRequest)
	if err != nil {
		fmt.Println("Error unmarshaling request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Convert PatientRequest to JSON
	requestBody, err := json.Marshal(patientRequest)
	if err != nil {
		fmt.Println("Error marshaling request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create a POST request
	req, err := http.NewRequest("POST", firestoreURL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Add the API key as a query parameter
	q := req.URL.Query()
	q.Add("key", apiKey)
	req.URL.RawQuery = q.Encode()

	// Set content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Process the response (responseBody) as needed
	fmt.Println("Response:", string(responseBody))

	// Return the response to the client
	w.Header().Add("Content-Type", "application/json")
	w.Write(responseBody)
}

// func CreateSnake(w http.ResponseWriter, r *http.Request) {
// 	var snake Snake
// 	err := json.NewDecoder(r.Body).Decode(&snake)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("Invalid request payload"))
// 		return
// 	}

// 	snake.ID = primitive.NewObjectID()

// 	_, err = collection.InsertOne(nil, snake)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("Unable to create snake"))
// 		return
// 	}

// 	response := struct {
// 		Message string             `json:"message"`
// 		Data    primitive.ObjectID `json:"data"`
// 	}{
// 		Message: "New Snake Created!",
// 		Data:    snake.ID,
// 	}

// 	json.NewEncoder(w).Encode(response)
// }

func GetSnakeFromSpec(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Pattern   string `json:"pattern"`
		HeadShape string `json:"headShape"`
		Color     string `json:"color"`
	}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil || params.Pattern == "" || params.HeadShape == "" || params.Color == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Params cannot be null"))
		return
	}

	filter := bson.M{
		"pattern":   params.Pattern,
		"headShape": params.HeadShape,
		"color":     bson.M{"$regex": params.Color},
	}

	var snake Snake
	err = collection.FindOne(nil, filter).Decode(&snake)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Couldn't find any mating result"))
		return
	}

	json.NewEncoder(w).Encode(snake)
}
