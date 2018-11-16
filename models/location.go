package models

// Course model is a struct that contains the name, time_ID, Location_ID.
type Course struct {
	ID        		int    `db:"id"`
	roomNumber 	  	string `db:"roomNumber"`
	buildingName   	string `db:"buildingName"`
	street     		string `db:"street"`
	city			string `db:"city"`
	postalCode		string `db:"postalCode"`
	state			string `db:"state"`
}