package models

// Location model is a struct that contains the ID, roomNumber, buildingName, street city, postalCode, state.
type Location struct {
	ID           int    `db:"id"`
	RoomNumber   string `db:"roomNumber"`
	BuildingName string `db:"buildingName"`
	Street       string `db:"street"`
	City         string `db:"city"`
	PostalCode   string `db:"postalCode"`
	State        string `db:"state"`
}
