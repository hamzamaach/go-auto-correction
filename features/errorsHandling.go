package reloaded

import (
	"fmt"
)

func VerifyArgs(args []string) bool {
	err := false
	if len(args) == 0 {
		err = true
		fmt.Println("Error: The name of the source file and the output file are missing!")
	} else if len(args) == 1 {
		err = true
		fmt.Println("Error: The name of the output file is missing!")
	} else if len(args) > 2 {
		err = true
		fmt.Println("Error: Too many arguments!")
	}
	return err
}
