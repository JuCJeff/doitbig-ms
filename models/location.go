package models

// Location model is a struct that contains the ID, roomNumber, buildingName, street city, postalCode, state.
type Location struct {
	ID        		int    `db:"id"`
	roomNumber 	  	string `db:"roomNumber"`
	buildingName   	string `db:"buildingName"`
	street     		string `db:"street"`
	city			string `db:"city"`
	postalCode		string `db:"postalCode"`
	state			string `db:"state"`
}