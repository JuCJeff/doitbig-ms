package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/jucjeff/doitb1g_backend/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
