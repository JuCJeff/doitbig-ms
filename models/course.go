package models

// Course model is a struct that contains the name, time_ID, Location_ID.
type Course struct {
	ID          int    `db:"id"`
	name        string `db:"name"`
	time_ID     string `db:"time_ID"`
	location_ID string `db:"location_ID"`
}
