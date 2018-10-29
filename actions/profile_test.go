package actions

func (as *ActionSuite) Test_ProfileHandler() {
	res := as.JSON("/profile").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "This is the signup page!")
}
