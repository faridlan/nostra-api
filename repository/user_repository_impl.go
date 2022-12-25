package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/faridlan/nostra-api/helper"
	"github.com/faridlan/nostra-api/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return UserRepositoryImpl{}
}

func (repository UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into user(username,email,password) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)

	return user
}

func (repository UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "update user set username = ?, email = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
	SQL := "delete from user where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (repository UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userid int) (domain.User, error) {
	SQL := "select id,username,email from user where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userid)
	helper.PanicIfError(err)

	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		helper.PanicIfError(err)

		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {

	// txs, err := repository.DB.Begin()
	// if err != nil {
	// 	panic(err)
	// }
	SQL := "select id,username,email from user"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	users := []domain.User{}
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		helper.PanicIfError(err)

		users = append(users, user)

	}

	return users
}
