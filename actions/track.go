package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"gitlab.com/doitbig-ms/models"
)

// TrackHandler is a handler to serve up
// a tracks page.
func TrackHandler(c buffalo.Context) error {
	track := models.Track{}
	db := c.Value("tx").(*pop.Connection)
	db.Find(&track, c.Param("tid"))
	return c.Render(200, r.JSON(track))
}
