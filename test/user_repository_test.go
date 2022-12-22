package test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/faridlan/nostra-api/helper"
	"github.com/faridlan/nostra-api/model/domain"
	"github.com/faridlan/nostra-api/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	Tx             *sql.Tx
	Mock           sqlmock.Sqlmock
	UserRepository repository.UserRepository
}

func (suite *Suite) SetupSuite() {
	var db *sql.DB
	var err error

	db, suite.Mock, err = sqlmock.New()

	suite.Tx, err = db.Begin()
	defer helper.CommitOrRollbak(suite.Tx)

	require.NoError(suite.T(), err)

	suite.UserRepository = repository.NewUserRepository()
}

func (Suite *Suite) TestUserRepo(T *testing.T) {

	user := domain.User{
		Id:       1,
		Username: "Udin",
		Email:    "udin@mail,com",
		Password: "secret",
	}

	Suite.Mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT id, username, email, password FROM user`,
	)).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"username",
		"email",
		"password",
	}).AddRow(user.Id, user.Username, user.Email, user.Password))

	result := Suite.UserRepository.FindAll(context.Background(), Suite.Tx)

	user2 := domain.User{
		Id:       999,
		Username: "Udin",
		Email:    "udin@mail,com",
		Password: "secret",
	}

	require.Equal(T, user2, result)

}

type User struct {
	Name string
}

var userRepo = repository.NewUserRepository()

func Test(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx, err := db.Begin()
	defer helper.CommitOrRollbak(tx)

	var (
		id       = 1
		username = "udin"
		email    = "udin@mail.com"
	)

	mock.ExpectQuery("select id,username,email from user").WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"username",
		"email",
	}).AddRow(id, username, email))

	result := userRepo.FindAll(context.Background(), tx)

	user := []domain.User{
		{Id: 1,
			Username: "Udin",
			Email:    "udin@mail,com",
		}}

	assert.Equal(t, user, result)
}
