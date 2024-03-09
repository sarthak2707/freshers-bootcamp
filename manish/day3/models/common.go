package models

type InsertMarksRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Subject     string `json:"subject"`
	Marks       int    `json:"marks"`
	Address     string `json:"address"`
	DateOfBirth int    `json:"date_of_birth"`
}
