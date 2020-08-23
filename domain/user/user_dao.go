package user

import (
	"github.com/pgrau/bookstore-user-api/datasource"
	"github.com/pgrau/bookstore-user-api/lib/date"
	"github.com/pgrau/bookstore-user-api/lib/error"
	"github.com/pgrau/bookstore-user-api/lib/mysql"
)

const(
	queryInsert = "INSERT INTO user(first_name, last_name, email, created_at) VALUES (?, ?, ?, ?);"
	queryUpdate = "UPDATE user SET first_name = ?, last_name = ?, email = ? WHERE id = ?;"
	queryDelete = "DELETE FROM user WHERE id = ?;"
	queryGetById = "SELECT id, first_name, last_name, email, created_at FROM user WHERE id = ?;"
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
		return mysql.ParseError(err)
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
		return mysql.ParseError(err)
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return mysql.ParseError(err)
	}

	user.Id = userId

	return nil
}

func (user *User) Update() *error.RestErr {
	stmt, err := datasource.MysqlClient.Prepare(queryUpdate)
	if err != nil {
		return error.InternalServerError(err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *error.RestErr {
	stmt, err := datasource.MysqlClient.Prepare(queryDelete)
	if err != nil {
		return error.InternalServerError(err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return mysql.ParseError(err)
	}

	return nil
}
