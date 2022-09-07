# Golang Restful API using GORM ORM (MySQL)

## Getting Started

### Folder Structure
```
.
|-- .env
|-- main.go
|-- README.md
|   +-- app
|   |   +-- controllers
|   |   |   |-- usersController.go
|   |   |   +-- auth
|   |   |   |   |-- authController.go
|   |   +-- models
|   |   |   |-- base.go
|   |   |   |-- user.go
|   +-- routes
|   |  |-- api.go
|   +-- utils
|   |   |-- utils.go
```
Ensure you create directory in your directory.

`git clone https://github.com/aminshokripwa/Golang-Restful-API-using-GORM.git`

## Download the packages used to create this rest API
Run the following Golang commands to install all the necessary packages. These packages will help you set up a web server, ORM for interacting with your db, mysql driver for db connection, load your environment variables from the .env file and generate JWT tokens.

```
go get -u github.com/gorilla/mux
go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql
go get -u github.com/joho/godotenv
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/gorilla/handlers

go get -u go get -u github.com/gorilla/handlers
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/gorilla/context
```

## Setting configuration file
Edit .env file in the root of the project and set the parameters for connect to database

## Running the project

`go run main.go`

## Database Table Creation Statement
Use the following DDL (Data Definition Language) to create the users table.

``` SQL
CREATE TABLE `users` (
  `id` int(11) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `tokens` varchar(200) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

## API Endpoints & Usage

To be able to login, you need to use the create new user endpoint to set up a user account.

* POST    `api/v1/login` login with username and password in your db to get token back

```
{
	"Username": "aminpwa@fake-domain.com",
	"Password": "secret"
}
```

*** Output ***

Note that the current implementation still returns the encrypted password, this needs to be removed from the response.

```
{
    "data": {
        "access-token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjkwMDAwMDAwMDAwMCwidXNlcm5hbWUiOiJhbWluLnB3YTFAZmFrZS1kb21haW4uY29tIn0.DIP2RmjA-ixG52ae1txevJIWZHVKvS8XbFuQxKPktkc"
    },
    "message": "success",
    "status": true
}
```

* POST    `api/v1/register` create a new account with a username, password and name

```
{
	"Name": "Amin_pwa",
	"Username": "aminpwa@fake-domain.com",
    "Password": "secret"
}
```

*** Output ***

Note that the current implementation still returns the encrypted password, this needs to be removed from the response.

```
{
    "message": "success",
    "status": true,
    "user": {
        "ID": 1,
        "CreatedAt": "2022-09-06T00:54:22.09382+01:00",
        "UpdatedAt": "2019-09-06T00:54:22.09382+01:00",
        "DeletedAt": null,
        "Name": "Amin_pwa",
        "Username": "Aminpwa@fake-domain.com"
    }
}

* GET     `api/v1/users?page=1&limit=4` Retrieving the paged list of users
* GET     `api/v1/users/1` Retrieve user with id = 1
* PUT     `api/v1/users/1` Update the record with id = 1
* DELETE  `api/v1/users/1` Delete the user with id = 1

Now make all calls pass the token in the header as a ***Bearer Token***.
