package errx

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		ErrorX(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		ErrorX(errorCommit)
	}
}
