package user

import (
	"fmt"
	"github.com/pgrau/bookstore-user-api/datasource"
	"github.com/pgrau/bookstore-user-api/lib/error"
	"github.com/pgrau/bookstore-user-api/lib/logger"
	"github.com/pgrau/bookstore-user-api/lib/mysql"
)

const(
	queryInsert = "INSERT INTO user(first_name, last_name, email, created_at, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	queryUpdate = "UPDATE user SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryDelete = "DELETE FROM user WHERE id = ?;"
	queryGetById = "SELECT id, first_name, last_name, email, created_at, status FROM user WHERE id = ?;"
	queryFindByStatus = "SELECT id, first_name, last_name, email, created_at, status FROM user WHERE status = ?;"
)

var(
	userDB = make(map[int64]*User)
)

func (user *User) Get() *error.RestErr {
	stmt, err := datasource.MysqlClient.Prepare(queryGetById)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)

		return error.InternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Status); err != nil {
		return mysql.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *error.RestErr) {
	stmt, err := datasource.MysqlClient.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find by status user statement", err)

		return nil, error.InternalServerError("database error")
	}

	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to get rows of find by status user", err)

		return nil, error.InternalServerError("database error")
	}

	defer rows.Close();

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Status); err != nil {
			logger.Error("error when trying to to fetch an user of find by status", err)

			return nil, mysql.ParseError(err)
		}

		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, error.NotFound(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}

func (user *User) Save() *error.RestErr {
	stmt, err := datasource.MysqlClient.Prepare(queryInsert)
	if err != nil {
		logger.Error("error when trying to save an user", err)

		return error.InternalServerError("Database error")
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt, user.Status, user.Password)
	if err != nil {
		logger.Error("error when trying to save an user", err)

		return mysql.ParseError(err)
	}

	userId, err := result.LastInsertId()
	if err != nil {
		logger.Error("error when trying to save an user", err)

		return mysql.ParseError(err)
	}

	user.Id = userId

	return nil
}

func (user *User) Update() *error.RestErr {
	stmt, err := datasource.MysqlClient.Prepare(queryUpdate)
	if err != nil {
		logger.Error("error when trying to update an user", err)

		return error.InternalServerError(err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to update an user", err)

		return mysql.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *error.RestErr {
	stmt, err := datasource.MysqlClient.Prepare(queryDelete)
	if err != nil {
		logger.Error("error when trying to delete an user", err)

		return error.InternalServerError(err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		logger.Error("error when trying to delete an user", err)

		return mysql.ParseError(err)
	}

	return nil
}
