package actions

func (as *ActionSuite) Test_LoginHandler() {
	res := as.HTML("/login").Get()
	as.Equal(200, res.Code)
}

