package actions

import "github.com/gobuffalo/buffalo"
import "fmt"

// SignupHandler is a handler to serve up
// a signup page.
func SignupHandler(c buffalo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	//this line's used to check if we've received the fields in http request body json at backend
	fmt.Fprintf(c.Response(), "Successfully received frontend json in request body\n first name is %s\n last name is %s\n email is %s\n password is %s", u.FirstName, u.LastName, u.Email, u.Password)

	//then we can manipulate with u.FirstName u.LastName u.Email u.Password

	return c.Render(200, r.JSON(map[string]string{"message": "This is the signup page!"}))
}
