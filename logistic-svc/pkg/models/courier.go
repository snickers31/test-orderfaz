package models

import "gorm.io/gorm"

type Courier struct {
	gorm.Model
	LogisticName    string `json:"logistic_name"`
	OriginName      string `json:"origin_name"`
	DestinationName string `json:"destination_name"`
	Amount          int64  `json:"amount"`
	Duration        string `json:"duration"`
}
