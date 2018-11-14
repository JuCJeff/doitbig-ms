package actions

import "github.com/gobuffalo/buffalo"

//import "fmt"
import "time"

// SignupHandler is a handler to serve up
// a signup page.
func SignupHandler(c buffalo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	//this line's used to check if we've received the fields in http request body json at backend
	//fmt.Fprintf(c.Response(), "Successfully received frontend json in request body\n first name is %s\n last name is %s\n email is %s\n password is %s", u.FirstName, u.LastName, u.Email, u.Password)

	//TODO: then we can manipulate with u.FirstName u.LastName u.Email u.Password
	//TODO: and give user a new user_id

	//this id is just for test, we'll use real userID later
	userID := "testCookieFromSignup"
	exp := time.Now().Add(15 * 24 * time.Hour) // expire in 15 days
	c.Cookies().SetWithExpirationTime("user_id", userID, exp)
	//successfully set cookies, now redirect to homepage
	return c.Redirect(307, "rootPath()")
	//return c.Render(200, r.JSON(map[string]string{"message": "This is the signup page!"}))
}
