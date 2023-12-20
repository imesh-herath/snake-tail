package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"snake-tail/config"
	"snake-tail/domain/entities"

	"net/http"

	"github.com/gorilla/mux"
)

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

func GetSnakeFromSpec(w http.ResponseWriter, r *http.Request) {

}
