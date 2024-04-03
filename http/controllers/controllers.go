package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"snake-tail/config"
	"snake-tail/domain/entities"
	"strings"

	"net/http"

	"github.com/gorilla/mux"
)

func GetSnakes(w http.ResponseWriter, r *http.Request) {

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

	apiResponse := &entities.SS{}
	err = json.Unmarshal(body, apiResponse)
	fmt.Println(err)

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

// GetSnakeFromSpec retrieves a snake from the snakeMap based on the provided specifications.
func GetSnakeFromSpec(w http.ResponseWriter, r *http.Request) {

	firestoreURL := config.App.FirebaseSnake.Url
	apiKey := config.App.FirebaseSnake.ApiKey

	// Create an HTTP client
	client := &http.Client{}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Unmarshal the JSON request body into a struct representing the specifications
	var spec entities.SnakeSpec
	err = json.Unmarshal(body, &spec)
	if err != nil {
		fmt.Println("Error unmarshaling request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
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

	// Decode the response body into a map
	var firestoreResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&firestoreResponse)
	if err != nil {
		fmt.Println("Error decoding response from Firestore:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Extract snake documents from Firestore response
	snakeDocuments, ok := firestoreResponse["documents"].([]interface{})
	if !ok {
		fmt.Println("Invalid response format from Firestore")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Iterate over the snake documents
	for _, doc := range snakeDocuments {
		snakeDoc, ok := doc.(map[string]interface{})
		if !ok {
			fmt.Println("Invalid document format from Firestore")
			continue
		}

		// Extract document name
		docName, ok := snakeDoc["name"].(string)
		if !ok {
			fmt.Println("Invalid document name format from Firestore")
			continue
		}

		// Extract fields from the document
		fields, ok := snakeDoc["fields"].(map[string]interface{})
		if !ok {
			fmt.Println("Invalid fields format from Firestore")
			continue
		}

		// Check if the document meets the criteria
		if field, ok := fields["head_shape"]; ok {
			if headShape, ok := field.(map[string]interface{})["stringValue"].(string); ok && headShape == spec.Fields.HeadShape &&
				spec.Fields.SkinColor != "" && spec.Fields.SkinColor == fields["skin_color"].(map[string]interface{})["stringValue"].(string) &&
				spec.Fields.SkinPattern != "" && spec.Fields.SkinPattern == fields["skin_pattern"].(map[string]interface{})["stringValue"].(string) {
				// Document matches the criteria for head_shape, skin_color, and skin_pattern
				fmt.Printf("Document '%s' matches the criteria for head_shape, skin_color, and skin_pattern\n", docName)

				// Split the document path by "/"
				parts := strings.Split(docName, "/")

				// Get the last part of the path
				documentName := parts[len(parts)-1]

				// Redirect to the /snakes/{id} endpoint
				http.Redirect(w, r, "/snakes/"+documentName, http.StatusFound)
				break
			}
		}
	}

	// Read the response body
	firestoreBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response from Firestore:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Process the response from Firestore as needed
	// Here you can parse the response to extract the document names or other relevant data
	fmt.Println(string(firestoreBody))

	// Return the response as the HTTP response
	w.Header().Add("Content-Type", "application/json")
	w.Write(firestoreBody)
}
func GetUniqueHeadShapes(w http.ResponseWriter, r *http.Request) {
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


    // Unmarshal the JSON request body into a struct representing the Firestore response
    var firestoreResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&firestoreResponse)
	if err != nil {
		fmt.Println("Error decoding response from Firestore:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

    // Extract snake documents from Firestore response
    snakeDocuments, ok := firestoreResponse["documents"].([]interface{})
    if !ok {
        fmt.Println("Invalid response format from Firestore")
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Call the function to get unique head shapes
    uniqueHeadShapes := GetHeadShapesFromDocuments(snakeDocuments)

    // Marshal the unique head shapes into JSON
    uniqueHeadShapesJSON, err := json.Marshal(uniqueHeadShapes)
    if err != nil {
        fmt.Println("Error marshaling unique head shapes:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Set the response content type header
    w.Header().Set("Content-Type", "application/json")

    // Write the response
    _, err = w.Write(uniqueHeadShapesJSON)
    if err != nil {
        fmt.Println("Error writing response:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}


func GetHeadShapesFromDocuments(snakeDocuments []interface{}) []string {
	// Initialize a map to store unique head_shape entries
	uniqueHeadShapes := make(map[string]bool)

	// Iterate over the snake documents
	for _, doc := range snakeDocuments {
		snakeDoc, ok := doc.(map[string]interface{})
		if !ok {
			fmt.Println("Invalid document format from Firestore")
			continue
		}

		// Extract fields from the document
		fields, ok := snakeDoc["fields"].(map[string]interface{})
		if !ok {
			fmt.Println("Invalid fields format from Firestore")
			continue
		}

		// Extract head_shape field from the document's fields
		headShapeField, ok := fields["head_shape"]
		if !ok {
			fmt.Println("head_shape field not found in document")
			continue
		}

		// Extract head_shape value from the head_shape field
		headShapeValue, ok := headShapeField.(map[string]interface{})["stringValue"].(string)
		if !ok {
			fmt.Println("Invalid head_shape value format from Firestore")
			continue
		}

		// Add the head_shape value to the map if it doesn't already exist
		if _, exists := uniqueHeadShapes[headShapeValue]; !exists {
			uniqueHeadShapes[headShapeValue] = true
		}
	}

	// Convert the map keys to a slice
	var result []string
	for headShape := range uniqueHeadShapes {
		result = append(result, headShape)
	}

	return result
}

func GetUniqueSkinColor(w http.ResponseWriter, r *http.Request) {
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


    // Unmarshal the JSON request body into a struct representing the Firestore response
    var firestoreResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&firestoreResponse)
	if err != nil {
		fmt.Println("Error decoding response from Firestore:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

    // Extract snake documents from Firestore response
    snakeDocuments, ok := firestoreResponse["documents"].([]interface{})
    if !ok {
        fmt.Println("Invalid response format from Firestore")
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Call the function to get unique head shapes
    uniqueSkinColor := GetSnakeColorFromDocuments(snakeDocuments)

    // Marshal the unique head shapes into JSON
    uniqueSkinColorJSON, err := json.Marshal(uniqueSkinColor)
    if err != nil {
        fmt.Println("Error marshaling unique head shapes:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Set the response content type header
    w.Header().Set("Content-Type", "application/json")

    // Write the response
    _, err = w.Write(uniqueSkinColorJSON)
    if err != nil {
        fmt.Println("Error writing response:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}

func GetSnakeColorFromDocuments(snakeDocuments []interface{}) []string {
	// Initialize a map to store unique skin_color entries
	uniqueSkinColor := make(map[string]bool)

	// Iterate over the snake documents
	for _, doc := range snakeDocuments {
		snakeDoc, ok := doc.(map[string]interface{})
		if !ok {
			fmt.Println("Invalid document format from Firestore")
			continue
		}

		// Extract fields from the document
		fields, ok := snakeDoc["fields"].(map[string]interface{})
		if !ok {
			fmt.Println("Invalid fields format from Firestore")
			continue
		}

		// Extract skin_color field from the document's fields
		skinColorField, ok := fields["skin_color"]
		if !ok {
			fmt.Println("skin_color field not found in document")
			continue
		}

		// Extract skin_color value from the skin_color field
		skinColorValue, ok := skinColorField.(map[string]interface{})["stringValue"].(string)
		if !ok {
			fmt.Println("Invalid skin_color value format from Firestore")
			continue
		}

		// Add the skin_color value to the map if it doesn't already exist
		if _, exists := uniqueSkinColor[skinColorValue]; !exists {
			uniqueSkinColor[skinColorValue] = true
		}
	}

	// Convert the map keys to a slice
	var result []string
	for headShape := range uniqueSkinColor {
		result = append(result, headShape)
	}

	return result
}

func GetUniqueSkinPattern(w http.ResponseWriter, r *http.Request) {
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


    // Unmarshal the JSON request body into a struct representing the Firestore response
    var firestoreResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&firestoreResponse)
	if err != nil {
		fmt.Println("Error decoding response from Firestore:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

    // Extract snake documents from Firestore response
    snakeDocuments, ok := firestoreResponse["documents"].([]interface{})
    if !ok {
        fmt.Println("Invalid response format from Firestore")
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Call the function to get unique head shapes
    uniqueSkinPattern := GetSnakePatternFromDocuments(snakeDocuments)

    // Marshal the unique head shapes into JSON
    uniqueSkinPatternJSON, err := json.Marshal(uniqueSkinPattern)
    if err != nil {
        fmt.Println("Error marshaling unique head shapes:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Set the response content type header
    w.Header().Set("Content-Type", "application/json")

    // Write the response
    _, err = w.Write(uniqueSkinPatternJSON)
    if err != nil {
        fmt.Println("Error writing response:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}

func GetSnakePatternFromDocuments(snakeDocuments []interface{}) []string {
	// Initialize a map to store unique skin_pattern entries
	skinpattern := make(map[string]bool)

	// Iterate over the snake documents
	for _, doc := range snakeDocuments {
		snakeDoc, ok := doc.(map[string]interface{})
		if !ok {
			fmt.Println("Invalid document format from Firestore")
			continue
		}

		// Extract fields from the document
		fields, ok := snakeDoc["fields"].(map[string]interface{})
		if !ok {
			fmt.Println("Invalid fields format from Firestore")
			continue
		}

		// Extract skin_pattern field from the document's fields
		skinPatternField, ok := fields["skin_pattern"]
		if !ok {
			fmt.Println("skin_pattern field not found in document")
			continue
		}

		// Extract skin_pattern value from the skin_pattern field
		skinPatternValue, ok := skinPatternField.(map[string]interface{})["stringValue"].(string)
		if !ok {
			fmt.Println("Invalid skin_pattern value format from Firestore")
			continue
		}

		// Add the skin_pattern value to the map if it doesn't already exist
		if _, exists := skinpattern[skinPatternValue]; !exists {
			skinpattern[skinPatternValue] = true
		}
	}

	// Convert the map keys to a slice
	var result []string
	for headShape := range skinpattern {
		result = append(result, headShape)
	}

	return result
}
