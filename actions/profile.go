package actions

import (
	"github.com/gobuffalo/buffalo"
)

// ProfileHandler is a handler to serve up
// a signup page.
func ProfileHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "This is the profile page! uid = " + c.Param("uid")}))
}
