package models

// Course model is a struct that contains the name, timeid, locationid.
type Course struct {
<<<<<<< HEAD
	ID         int    `db:"id"`
	Name       string `db:"name"`
	TimeID     string `db:"time_ID"`
	LocationID string `db:"location_ID"`
=======
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Location string `db:"location"`
	Time     string `db:"time"`
>>>>>>> 09def4daa89fbc989d71def47ff43a6dd7a8c32d
}
