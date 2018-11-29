package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	value, err := c.Cookies().Get("user_id")
	if err != nil {
		//TODO: render login page if cookie not found
		c.Redirect(307, "http://localhost:3000/signin")
	}
	//if cookie found, then further proceed with the user_id
	//this line if temporarily for Postman test with cookie return results
	//TODO: extract required info from DB and render homepage
	return c.Render(200, r.String(value))
	//return c.Redirect(307, "http://localhost:3000")
}
