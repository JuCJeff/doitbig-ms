package grifts

import (
	"github.com/gobuffalo/buffalo"
	"gitlab.com/doitbig-ms/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
