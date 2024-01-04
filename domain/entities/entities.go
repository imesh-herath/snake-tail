package entities

import "time"

type SnakeResponse struct {
	Name   string `json:"name"`
	Fields struct {
		ImageUrl struct {
			StringValue string `json:"stringValue"`
		} `json:"image_url"`
		Description struct {
			StringValue string `json:"stringValue"`
		} `json:"description"`
		ScientificName struct {
			StringValue string `json:"StringValue"`
		} `json:"scientific_name"`
		SnakeName struct {
			StringValue string `json:"StringValue"`
		} `json:"snake_name"`
	} `json:"fields"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
type PatientRequest struct {
	Fields struct {
		PatientName struct {
			StringValue string `json:"stringValue"`
		} `json:"patient_name"`

		PhoneNumber struct {
			IntegerValue string `json:"integerValue"`
		} `json:"phone_number"`

		BittenTime struct {
			TimestampValue time.Time `json:"timestampValue"`
		} `json:"bitten_time"`

		SnakeImage struct {
			SnakeImageUrl string `json:"stringValue"`
		} `json:"snake_image_url"`
	} `json:"fields"`
}
