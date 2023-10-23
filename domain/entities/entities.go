package entities

import "time"

type SnakeResponse struct {
	Name   string `json:"name"`
	Fields struct {
		PatientName struct {
			StringValue string `json:"stringValue"`
		} `json:"patient_name"`
		BittenTime struct {
			TimestampValue time.Time `json:"timestampValue"`
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
