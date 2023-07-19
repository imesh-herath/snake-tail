package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Snake struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Color          string             `bson:"color"`
	Description    string             `bson:"description"`
	FirstAid       string             `bson:"firstAid"`
	HeadShape      string             `bson:"headShape"`
	Image          string             `bson:"image"`
	Name           string             `bson:"name"`
	OtherName      string             `bson:"otherName"`
	Pattern        string             `bson:"pattern"`
	ScientificName string             `bson:"scientificName"`
	VenomLevel     string             `bson:"venomLevel"`
}