package helper

import "database/sql"

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {

		errorRollBack := tx.Rollback()
		PanicIfErr(errorRollBack)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfErr(errorCommit)
	}
}

func Offset(query string, tx *sql.Tx) *sql.Stmt {
	stmt, err := tx.Prepare(query)
	PanicIfErr(err)
	return stmt
}
