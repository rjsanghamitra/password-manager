package new

import (
	"fmt"
	"os"
	"pwdmgr/database"
	"pwdmgr/error"
	"pwdmgr/hash"
)

// this is a function to create a new account for a user.
func Create(name string, pwd string) {
	pwdhash := hash.Hash(pwd)
	namehash := hash.Hash(name)
	dbname := namehash + pwdhash

	db := database.CreateDb(dbname)
	database.CreateTable(*db, "pwditems")
	defer db.Close()

	file, err := os.OpenFile("databases.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	error.CheckError(err)
	defer file.Close()
	file.WriteString(dbname + "\n")

	fmt.Println("Account created!")
}
