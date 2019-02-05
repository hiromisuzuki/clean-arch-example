package model

import "time"

//Timestamps ISO 8601 format Timestamp https://ja.wikipedia.org/wiki/ISO_8601
type Timestamps struct {
	CreatedAt *time.Time `json:"created_at" validate:"date-time" example:"2017-07-21T17:32:28Z"`
	UpdatedAt *time.Time `json:"updated_at" validate:"date-time" example:"2017-07-21T17:32:28Z"`
}
