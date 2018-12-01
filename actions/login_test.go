package actions

func (as *ActionSuite) Test_LoginHandler() {
	res := as.HTML("/login").Get()
	as.Equal(307, res.Code)
}
