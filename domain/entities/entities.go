package entities

import "time"

type SnakeResponse struct {
	Name   string `json:"name"`
	Fields struct {
		PatientName struct {
			StringValue string `json:"stringValue"`
		} `json:"patient_name"`
		BittenTime struct {
			TimestampValue string `json:"timestampValue"`
		} `json:"bitten_time"`
		SnakeImageUrl struct {
			StringValue string `json:"stringValue"`
		} `json:"snake_image_url"`
		Id struct {
			IntegerValue string `json:"integerValue"`
		} `json:"id"`
		PhoneNumber struct {
			IntegerValue string `json:"integerValue"`
		} `json:"phone_number"`
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
