package login

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"pwdmgr/database"
	"pwdmgr/dbsecurity"
	"pwdmgr/error"
	"pwdmgr/hash"
)

var (
	Dbname string
	Key    string
)

func Login(name string, pwd string) {
	// fetching the existing database names.
	var names []string
	f, err := os.Open("databases.txt")
	error.CheckError(err)
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		names = append(names, sc.Text())
	}

	// the name of the database file is hash of the username concatenated with the hash of the master password.
	h := hash.Hash(name) + hash.Hash(pwd)
	Dbname = h
	Key = hash.GetMD5Hash(pwd)
	var flag bool
	var i string
	for _, i = range names {
		if i == h {
			flag = true
			break
		} else {
			flag = false
		}
	}

	// if the user has entered either the username or the password wrong.
	if !flag {
		log.Println("The entered username and password don't exist. Please enter the correct username and password.")
		os.Exit(1)
	}
	fmt.Print("\n\nLogged in.\n\n")
}

func AddNewItem(site string, username string, pwd string) {
	db, err := sql.Open("sqlite3", Dbname+".db")
	error.CheckError(err)
	site = dbsecurity.Encrypt(site, Key)
	username = dbsecurity.Encrypt(username, Key)
	pwd = dbsecurity.Encrypt(pwd, Key)
	// creating a record in the database that stores the encrypted form of the site, username and password.
	database.InsertItem(db, site, username, pwd)
}

func RetrieveItem(site string) {
	db, err := sql.Open("sqlite3", Dbname+".db")
	error.CheckError(err)
	// this map stores key value pairs where the key is the username and the value is the password.
	mp := database.GetItem(db, dbsecurity.Encrypt(site, Key))

	// looping over the map and displaying the usernames and their corresponding passwords for the requested site
	for k, v := range mp {
		fmt.Println("Username:", dbsecurity.Decrypt(k, Key), "\tPassword:", dbsecurity.Decrypt(v, Key))
	}
}

func DeleteItem(site string, username string) {
	db, err := sql.Open("sqlite3", Dbname+".db")
	error.CheckError(err)
	site = dbsecurity.Encrypt(site, Key)
	username = dbsecurity.Encrypt(username, Key)
	database.DeleteItem(db, site, username)
}

func UpdateItem(site string, username string, newpwd string) {
	db, err := sql.Open("sqlite3", Dbname+".db")
	error.CheckError(err)
	site = dbsecurity.Encrypt(site, Key)
	username = dbsecurity.Encrypt(username, Key)
	newpwd = dbsecurity.Encrypt(newpwd, Key)
	database.UpdateItem(db, site, username, newpwd)
}
