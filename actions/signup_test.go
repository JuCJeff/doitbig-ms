package actions

func (as *ActionSuite) Test_SignupHandler() {
	res := as.HTML("/signup").Get()
	as.Equal(200, res.Code)
}
