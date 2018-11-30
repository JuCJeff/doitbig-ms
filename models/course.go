package models

// Course model is a struct that contains the name, timeid, locationid.
type Course struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Location string `db:"location"`
	Time     string `db:"time"`
}
