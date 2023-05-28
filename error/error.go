package error

import "log"

// this is just a function that is used for checking errors and handling them if they are not nil.
func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}
