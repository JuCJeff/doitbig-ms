package actions

import "github.com/gobuffalo/buffalo"

//import "fmt"
import "time"

// LoginHandler is a handler to serve up
// a signup page.
func LoginHandler(c buffalo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	//this line's used to check if we've received the fields in http request body json at backend
	//fmt.Fprintf(c.Response(), "login route successfully received frontend json in request body\n email is %s\n password is %s", u.Email, u.Password)

	//TODO: should also check cookies but I'll leave that for later

	//TODO: then we can manipulate with u.Email u.Password to find the user_id

	//this id is just for test, we'll use real userID later
	userID := "testCookieFromLogin"
	exp := time.Now().Add(15 * 24 * time.Hour) // expire in 15 days
	c.Cookies().SetWithExpirationTime("user_id", userID, exp)
	//successfully set cookies, now redirect to homepage
	return c.Redirect(307, "rootPath()")
}
