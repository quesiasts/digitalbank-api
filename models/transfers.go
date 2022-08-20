package models

import "time"

type Transfers struct {
	Id                     string    `json:"id"`
	Account_origin_id      int       `json:"account_origin_id,omitempty"`
	Account_destination_id int       `json:"account_destinations_id,omitempty"`
	Ammount                float64   `json:"ammount,omitempty"`
	Created_at             time.Time `json:"created_at,omitempty"`
}
