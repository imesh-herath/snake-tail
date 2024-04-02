package entities

import "time"

type SnakeResponse struct {
	Name       string    `json:"name"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Fields     Body      `json:"fields"`
}

type SS struct{
	Doc []SnakeResponse `json:"documents"`

}

type Body struct {
	ImageUrl       string `json:"image_url"`
	Description    string `json:"description"`
	ScientificName string `json:"scientific_name"`
	SnakeName      string `json:"snake_name"`
	SkinColor      string `json:"skin_color"`
	SkinPattern    string `json:"skin_pattern"`
	HeadShape      string `json:"head_shape"`
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

type SnakeSpec struct {
	Fields struct {
		SkinColor   string `json:"skin_color"`
		SkinPattern string `json:"skin_pattern"`
		HeadShape   string `json:"head_shape"`
	} `json:"fields"`
}
