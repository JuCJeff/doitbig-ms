package actions

import (
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"gitlab.com/doitbig-ms/models"
)

// ProfileHandler is a handler to serve up
// user profile data from the database
func ProfileHandler(c buffalo.Context) error {
	fmt.Println(c.Value("tx"))
	user := models.User{}
	db := c.Value("tx").(*pop.Connection)
	db.Find(&user, c.Param("uid"))
	return c.Render(200, r.JSON(user))
}
