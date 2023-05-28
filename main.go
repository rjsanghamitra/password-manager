package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pwdmgr/dbsecurity"
	"pwdmgr/error"
	"pwdmgr/login"
	"pwdmgr/new"
)

func main() {
	fmt.Println("Hello!\nAre you a new user? Enter y for Yes and n for No.")
	reader := bufio.NewReader(os.Stdin)
	opt, _, err := reader.ReadRune()
	error.CheckError(err)

	if opt == 'y' {
		fmt.Println("Enter your username:")
		var name string
		fmt.Scan(&name)
		fmt.Println("Create your master password:")
		var pwd string
		fmt.Scan(&pwd)

		// creating a new account for the user.
		new.Create(name, pwd)
	} else {
		// if the user already has an account.
		fmt.Println("Enter your username:")
		var name string
		fmt.Scan(&name)
		fmt.Println("Enter your master password:")
		var pwd string
		fmt.Scan(&pwd)
		// logging in.
		login.Login(name, pwd)

		fmt.Println("Enter the operation you want to perform:")
		fmt.Println("1. Add new password item.")
		fmt.Println("2. Retrieve a password item from the database.")
		fmt.Println("3. Delete a password item.")
		fmt.Print("4. Update the password of an existing password item.\n\n")
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
				pwd = dbsecurity.RandomPasswordGenerator()
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
				pwd = dbsecurity.RandomPasswordGenerator()
			} else {
				fmt.Println("Enter the password:")
				fmt.Scan(&pwd)
			}
			login.UpdateItem(site, username, newpwd)

		} else {
			// if the option entered by the user is neither 1, 2, 3, or 4.
			log.Println("Invalid Option")
			os.Exit(1)
		}
	}
}
