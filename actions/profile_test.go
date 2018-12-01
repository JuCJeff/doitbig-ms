package actions

// Tests valid URI
func (as *ActionSuite) Test_ProfileHandler() {
	res := as.JSON("/profile").Get()
	as.Equal(200, res.Code)
}

// Tests fetching valid data from DB
func (as *ActionSuite) Test_ProfileGetter() {
	res := as.JSON("/profile?uid=12345").Get()
	as.Equal(200, res.Code)
}
