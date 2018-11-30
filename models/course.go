package models

// Course model is a struct that contains the name, time_ID, Location_ID.
type Course struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	TimeID     string `db:"time_ID"`
	LocationID string `db:"location_ID"`
}
