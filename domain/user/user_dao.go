package user

import (
	"fmt"
	"github.com/pgrau/bookstore-user-api/datasource"
	"github.com/pgrau/bookstore-user-api/lib/date"
	"github.com/pgrau/bookstore-user-api/lib/error"
	"strings"
)

const(
	indexUniqueEmail = "email_UNIQUE"
	queryInsert = "INSERT INTO user(first_name, last_name, email, created_at) VALUES (?, ?, ?, ?);"
	queryGetById = "SELECT id, first_name, last_name, email, created_at FROM user WHERE id = ?;"
	errorNoRows = "no rows in result set"
)

var(
	userDB = make(map[int64]*User)
)

func (user *User) Get() *error.RestErr {
	stmt, err := datasource.MysqlClient.Prepare(queryGetById)
	if err != nil {
		return error.InternalServerError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return error.NotFound(fmt.Sprintf("user %d not found", user.Id));
		}

		return error.InternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}

	return nil
}

func (user *User) Save() *error.RestErr {
	stmt, err := datasource.MysqlClient.Prepare(queryInsert)
	if err != nil {
		return error.InternalServerError(err.Error())
	}

	defer stmt.Close()

	user.CreatedAt = date.GetNowString();

	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return error.Conflict(fmt.Sprintf("email %s already exists", user.Email))
		}
		return error.InternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return error.InternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	user.Id = userId

	return nil
}
