package actions

import "github.com/gobuffalo/buffalo"
import "fmt"

// LoginHandler is a handler to serve up
// a signup page.
func LoginHandler(c buffalo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	//this line's used to check if we've received the fields in http request body json at backend
	fmt.Fprintf(c.Response(), "Successfully received frontend json in request body\n email is %s\n password is %s", u.Email, u.Password)

	//then we can manipulate with u.Email u.Password

	return c.Render(200, r.JSON(map[string]string{"message": "This is the Login page!"}))
}

