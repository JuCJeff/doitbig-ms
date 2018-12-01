package models

// User model is a struct that contains the id, firstName, lastNAme, email, password, and image of student struct.
type User struct {
	ID        int    `db:"id"`
	FirstName string `db:"firstname"`
	LastName  string `db:"lastname"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Image     string `db:"image"`
}
