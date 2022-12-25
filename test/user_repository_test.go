package test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/faridlan/nostra-api/helper"
	"github.com/faridlan/nostra-api/model/domain"
	"github.com/faridlan/nostra-api/repository"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Id       int
	Username string
	Email    string
	Password string
}

func getTx(db *sql.DB) *sql.Tx {

	tx, err := db.Begin()
	helper.PanicIfError(err)

	return tx
}

func sqlMock() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		return nil, nil, err
	} else {
		return db, mock, nil
	}
}

var userRepo = repository.NewUserRepository()

func TestRepoFindAllSuccess(t *testing.T) {
	db, mock, err := sqlMock()
	helper.PanicIfError(err)

	defer db.Close()

	actualUser1 := User{
		Id:       1,
		Username: "Udin",
		Email:    "udin@mail.com",
	}

	actualUser2 := User{
		Id:       2,
		Username: "john",
		Email:    "john@mail.com",
	}

	mock.ExpectBegin()
	tx := getTx(db)
	mock.ExpectQuery("select id,username,email from user").WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"username",
		"email",
	}).AddRow(actualUser1.Id, actualUser1.Username, actualUser1.Email).AddRow(actualUser2.Id, actualUser2.Username, actualUser2.Email),
	)
	mock.ExpectCommit()

	result := userRepo.FindAll(context.Background(), tx)

	user := []domain.User{
		{
			Id:       1,
			Username: "Udin",
			Email:    "udin@mail.com",
		},
		{
			Id:       2,
			Username: "john",
			Email:    "john@mail.com",
		},
	}

	assert.Equal(t, user, result)
}

func TestRepoFindByIdSuccess(t *testing.T) {
	db, mock, err := sqlMock()
	helper.PanicIfError(err)

	defer db.Close()

	user := User{
		Id:       1,
		Username: "Udin",
		Email:    "udin@mail.com",
	}

	mock.ExpectBegin()
	tx := getTx(db)
	mock.ExpectQuery("select id,username,email from user where id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"usernmae",
		"email",
	}).AddRow(user.Id, user.Username, user.Email))
	mock.ExpectCommit()

	rows, err := userRepo.FindById(context.Background(), tx, 1)
	userResponse := domain.User{
		Id:       1,
		Username: "Udin",
		Email:    "udin@mail.com",
	}
	helper.PanicIfError(err)

	assert.Nil(t, err)
	assert.Equal(t, rows, userResponse)
}

func TestRepoFindByIdFailed(t *testing.T) {
	db, mock, err := sqlMock()
	helper.PanicIfError(err)

	defer db.Close()

	mock.ExpectBegin()
	tx := getTx(db)
	mock.ExpectQuery("select id,username,email from user where id = ?").WithArgs(2).WillReturnError(fmt.Errorf("user not found"))
	mock.ExpectRollback()

	_, err = userRepo.FindById(context.Background(), tx, 2)

	// userResponse := domain.User{}

	assert.NotNil(t, err)

}

func TestRepoSaveSuccess(t *testing.T) {

	db, mock, err := sqlMock()
	helper.PanicIfError(err)

	defer db.Close()

	user := User{
		Id:       1,
		Username: "Udin",
		Email:    "udin@mail.com",
		Password: "secret",
	}

	mock.ExpectBegin()
	tx := getTx(db)
	mock.ExpectExec("insert into user(username,email,password) values (?,?,?)").WithArgs(user.Username, user.Email, user.Password).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	userReq := domain.User{
		Id:       1,
		Username: "Udin",
		Email:    "udin@mail.com",
		Password: "secret",
	}

	userResponse := userRepo.Save(context.Background(), tx, userReq)

	assert.Equal(t, userResponse, userReq)
}

func TestRepoUpdateUser(t *testing.T) {
	db, mock, err := sqlMock()
	helper.PanicIfError(err)

	defer db.Close()

	user := User{
		Id:       1,
		Username: "Udin",
		Email:    "udin@mail.com",
	}

	mock.ExpectBegin()
	tx := getTx(db)
	mock.ExpectExec("update user set username = ?, email = ? where id = ?").WithArgs(user.Username, user.Email, user.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	userReq := domain.User{
		Id:       1,
		Username: "Udin",
		Email:    "udin@mail.com",
	}

	userResponse := userRepo.Update(context.Background(), tx, userReq)

	assert.Equal(t, userResponse, userReq)
}

func TestRepoDeleteUser(t *testing.T) {
	db, mock, err := sqlMock()
	helper.PanicIfError(err)

	defer db.Close()

	user := domain.User{
		Id:       1,
		Username: "Udin",
		Email:    "udin@mail.com",
	}

	mock.ExpectBegin()
	tx := getTx(db)
	mock.ExpectExec("delete from user where id = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	userRepo.Delete(context.Background(), tx, user)

}
