package models

// Student model is a struct that contains the id, coursetaken, upcomingcourse, and tracknum of user struct.
type Student struct {
	ID             int    `db:"id"`
	CourseTaken    string `db:"coursetaken"`
	UpcomingCourse string `db:"upcomingcourse"`
	TrackNum       int    `db:"tracknumber"`
}
