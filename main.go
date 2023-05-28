package main

import (
	"fmt"
	"log"
	"os"
	"pwdmgr/dbsecurity"
	"pwdmgr/login"
	"pwdmgr/new"
)

func main() {
	var username, mpwd string
	if len(os.Args) == 1 {
		fmt.Println("Enter your username:")
		var name string
		fmt.Scan(&name)
		fmt.Println("Create your master password:")
		var pwd string
		fmt.Scan(&pwd)

		// creating a new account for the user.
		new.Create(name, pwd)

	} else if len(os.Args) == 3 {
		username = os.Args[1]
		mpwd = os.Args[2]
		// logging in.
		login.Login(username, mpwd)
	}
	fmt.Println("Enter the operation you want to perform:")
	fmt.Println("1. Add new password item.")
	fmt.Println("2. Retrieve a password item from the database.")
	fmt.Println("3. Delete a password item.")
	fmt.Println("4. Update the password of an existing password item.")
	fmt.Print("5. Exit.\n\n")
	var opt int
	fmt.Scan(&opt)
	if opt == 1 {
		// creating a password item.
		var username, pwd, site string
		fmt.Println("Enter the username:")
		fmt.Scan(&username)

		// giving the user a choice to get a randomly generated password.
		fmt.Println("Do you want a randomly generated password? Enter 1 is Yes and 0 if No.")
		var r int
		fmt.Scan(&r)
		if r == 1 {
			fmt.Println("What do you want the length of the password to be?")
			var l int
			fmt.Scan(&l)
			pwd = dbsecurity.RandomPasswordGenerator(l)
		} else {
			fmt.Println("Enter the password:")
			fmt.Scan(&pwd)
		}
		fmt.Println("Enter the name of the site:")
		fmt.Scan(&site)
		// adding the password item to the database.
		login.AddNewItem(site, username, pwd)

	} else if opt == 2 {
		// reading from the database.
		fmt.Println("Enter the site for which you want to retrieve the username and password:")
		var site string
		fmt.Scan(&site)
		// select operation on the table in the database.
		login.RetrieveItem(site)

	} else if opt == 3 {
		// deleting a row from the database.
		var username, site string
		fmt.Println("Enter the username of the password item that you want to delete:")
		fmt.Scan(&username)
		fmt.Println("Enter the name of the site:")
		fmt.Scan(&site)
		login.DeleteItem(site, username)

	} else if opt == 4 {
		// updating a password in the database.
		var username, site, newpwd string
		fmt.Println("Enter the site:")
		fmt.Scan(&site)
		fmt.Println("Enter the username for which you want to update the password:")
		fmt.Scan(&username)
		fmt.Println("Do you want a randomly generated password? Enter 1 is Yes and 0 if No.")
		var r int
		fmt.Scan(&r)
		if r == 1 {
			fmt.Println("What do you want the length of the new password to be?")
			var l int
			fmt.Scan(&l)
			newpwd = dbsecurity.RandomPasswordGenerator(l)
		} else {
			fmt.Println("Enter the password:")
			fmt.Scan(&newpwd)
		}
		login.UpdateItem(site, username, newpwd)

	} else if opt == 5 {
		os.Exit(0)
	} else {
		// if the option entered by the user is neither 1, 2, 3, 4 or 5.
		log.Println("Invalid Option")
		os.Exit(1)
	}
}
