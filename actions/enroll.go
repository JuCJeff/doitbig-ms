package actions

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/buffalo"
)

// EnrollGetHandler is a handler to manage course enrollment
func EnrollGetHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "This is the enroll page!" +
		"- you typed " + c.Param("uid") + " for your uid"}))
}

// EnrollPostHandler is a handler to manage course enrollment
func EnrollPostHandler(c buffalo.Context) error {

	fmt.Println("POST: Enroll Handler")

	var uid = c.Param("uid")
	var cid = c.Param("cid")
	var nowTime = time.Now().Format("2006-01-02 15:04:05")

	db, _ := sql.Open("mysql", "sqladmin:asdfghjkl;'@tcp(doitbig.c3lglnwntifb.us-east-1.rds.amazonaws.com:3306)/doitb1gdb")

	insert, err := db.Query("INSERT INTO courseEnroll VALUES ('" + uid + "', '" + cid +
		"', '" + nowTime + "')")
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	return c.Render(200, r.JSON(map[string]string{"message": "UserID " + uid +
		" has successfully enrolled into classID " + cid + " at time " + nowTime}))
}
