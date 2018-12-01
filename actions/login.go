package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"gitlab.com/doitbig-ms/models"
)

// LoginHandler is a handler to serve up
// a signup page.
func LoginHandler(c buffalo.Context) error {
	var email = c.Param("email")
	var password = c.Param("password")
	user := models.User{}
	db := c.Value("tx").(*pop.Connection)
	query := db.Where(`email = ?`, email)
	err := query.First(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.Password)
	if user.Password == password {
		return c.Redirect(307, "http://localhost:3000")
	}
	return c.Redirect(307, "http://localhost:3000/signin")
}
