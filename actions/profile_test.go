package actions

import "github.com/gobuffalo/pop"

func (as *ActionSuite) Test_ProfileHandler() {
	res := as.JSON("/profile").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_ProfileGetter() {
	db := c.Value("tx").(*pop.Connection)
	res := as.JSON("/profile?uid=12345").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "John")
}
