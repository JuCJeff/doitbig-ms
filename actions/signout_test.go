package actions

func (as *ActionSuite) Test_SignoutHandler() {
	res := as.HTML("/signout").Get()
	as.Equal(200, res.Code)
}
