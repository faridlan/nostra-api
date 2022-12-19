package helper

import "database/sql"

func CommitOrRollbak(tx *sql.Tx) {

	err := recover()
	if err != nil {
		err := tx.Rollback()
		PanicIfError(err)

		panic(err)
	} else {
		err := tx.Commit()
		PanicIfError(err)
	}
}
