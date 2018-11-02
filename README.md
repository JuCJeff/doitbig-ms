# Welcome to doitb1g-ms!

This comprises the backend of the DoITB1G STS Project. The backend is responsible for making queries to the database and interfacing this data with the sister project doitb1g-ui. 

## Requirements

- Golang (v1.11+)
- SQLite 3
- Buffalo: https://gobuffalo.io/en/docs/installation

## Starting the Application

Our project is based on an underlying Go Web boilerplate that facilitates speed of development called Buffalo. Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" message.

**Congratulations!** You now have our application up and running.

## Supported Routes

Your can post GET requests to the following routes with the following syntax:

- http://localhost:3000/login
- http://localhost:3000/signup
- http://localhost:3000/signout
- http://localhost:3000/
- http://localhost:3000/profile?uid=xxxxx with `xxxxx` being a valid User ID. Omitting the query string will return blank data.

## Testing

To run tests, Buffalo features a test runner that runs the underlying Go tests of the project. 

	$ buffalo test

