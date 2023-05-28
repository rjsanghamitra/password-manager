package database

import (
	"database/sql"
	"os"
	"pwdmgr/error"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable(db sql.DB, name string) {
	com := "CREATE TABLE IF NOT EXISTS pwditems (\"site\" TEXT, \"username\" TEXT, \"password\" TEXT);"
	stmt, err := db.Prepare(com)
	error.CheckError(err)
	stmt.Exec()
	stmt.Close()
}

func CreateDb(name string) *sql.DB {
	_, err := os.Create(name + ".db")
	error.CheckError(err)
	db, err := sql.Open("sqlite3", name+".db")
	error.CheckError(err)
	return db
}

func InsertItem(db *sql.DB, site string, username string, pwd string) {
	com := "INSERT INTO pwditems (site, username, password) VALUES(?, ?, ?)"
	stmt, err := db.Prepare(com)
	error.CheckError(err)
	stmt.Exec(site, username, pwd)
	defer stmt.Close()
}

func DeleteItem(db *sql.DB, site string, username string) {
	com := "DELETE FROM pwditems WHERE site = ? AND username = ?"
	stmt, err := db.Prepare(com)
	error.CheckError(err)
	stmt.Exec(site, username)
	defer stmt.Close()
}

func GetItem(db *sql.DB, site string) map[string]string {
	com := "SELECT * FROM pwditems WHERE site = ?"
	stmt, err := db.Prepare(com)
	error.CheckError(err)
	rows, err := stmt.Query(site)
	error.CheckError(err)
	defer rows.Close()
	// storing the username and password as key-value pairs in a map
	mp := make(map[string]string, 1000)
	for rows.Next() {
		var temp, uname, pwd string
		err := rows.Scan(&temp, &uname, &pwd)
		error.CheckError(err)
		mp[uname] = pwd
	}
	return mp
}

func UpdateItem(db *sql.DB, site string, username string, newpwd string) {
	com := "UPDATE pwditems SET password = ? WHERE site = ? AND username = ?"
	stmt, err := db.Prepare(com)
	error.CheckError(err)
	stmt.Exec(newpwd, site, username)
}
