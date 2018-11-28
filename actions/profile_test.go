package actions

import "gitlab.com/doitbig-ms/models"

// Tests valid URI
func (as *ActionSuite) Test_ProfileHandler() {
	res := as.JSON("/profile").Get()
	as.Equal(200, res.Code)
}

// Tests fetching valid data from DB
func (as *ActionSuite) Test_ProfileGetter() {
	// testUser := &models.User{ID: 12345, FirstName: "John", LastName: "Test", Email: "Test", Password: "Test"}
	// err := as.DB.Create(testUser)
	// as.DB.Update(testUser)
	// as.NoError(err)
	err := as.DB.RawQuery("insert into users (id, firstname, lastname, email, password, image) values (12345, \"John\", \"Appleseed\",\"japple@wisc.edu\", \"password\", \"asdf\")")
	as.T().Log("Error: ", err)
	user := models.User{}
	as.DB.Find(&user, 12345)
	as.T().Log("THIS IS IT: ", user.ID)
	res := as.JSON("/profile?uid=12345").Get()
	as.Equal(200, res.Code)
	// as.Contains(res.Body.String(), "John")
}
