package controllers


import (
	"encoding/json"
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
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Color           string             `json:"color,omitempty" bson:"color,omitempty"`
	Description     string             `json:"description,omitempty" bson:"description,omitempty"`
	FirstAid        string             `json:"firstAid,omitempty" bson:"firstAid,omitempty"`
	HeadShape       string             `json:"headShape,omitempty" bson:"headShape,omitempty"`
	Image           string             `json:"image,omitempty" bson:"image,omitempty"`
	Name            string             `json:"name,omitempty" bson:"name,omitempty"`
	OtherName       string             `json:"otherName,omitempty" bson:"otherName,omitempty"`
	Pattern         string             `json:"pattern,omitempty" bson:"pattern,omitempty"`
	ScientificName  string             `json:"scientificName,omitempty" bson:"scientificName,omitempty"`
	VenomLevel      string             `json:"venomLevel,omitempty" bson:"venomLevel,omitempty"`
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
	snakes := []Snake{}
	cur, err := collection.Find(nil, bson.D{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to fetch data from database"))
		return
	}
	defer cur.Close(nil)

	for cur.Next(nil) {
		var snake Snake
		err := cur.Decode(&snake)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unable to fetch data from database"))
			return
		}
		snakes = append(snakes, snake)
	}

	json.NewEncoder(w).Encode(snakes)
}

func GetSnakeByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid snake ID"))
		return
	}

	var snake Snake
	err = collection.FindOne(nil, bson.M{"_id": id}).Decode(&snake)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to fetch data from database"))
		return
	}

	json.NewEncoder(w).Encode(snake)
}

func CreateSnake(w http.ResponseWriter, r *http.Request) {
	var snake Snake
	err := json.NewDecoder(r.Body).Decode(&snake)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	snake.ID = primitive.NewObjectID()

	_, err = collection.InsertOne(nil, snake)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to create snake"))
		return
	}

	response := struct {
		Message string             `json:"message"`
		Data    primitive.ObjectID `json:"data"`
	}{
		Message: "New Snake Created!",
		Data:    snake.ID,
	}

	json.NewEncoder(w).Encode(response)
}

func GetSnakeFromSpec(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Pattern    string `json:"pattern"`
		HeadShape  string `json:"headShape"`
		Color      string `json:"color"`
	}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil || params.Pattern == "" || params.HeadShape == "" || params.Color == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Params cannot be null"))
		return
	}

	filter := bson.M{
		"pattern":    params.Pattern,
		"headShape":  params.HeadShape,
		"color":      bson.M{"$regex": params.Color},
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
