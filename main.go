package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		content, _ := ioutil.ReadFile(args[0])
		var modifiedContent string
		modifiedContent = strings.ReplaceAll(string(content), "(up)", "(up,1)")
		modifiedContent = strings.ReplaceAll(string(modifiedContent), "(low)", "(low,1)")
		modifiedContent = strings.ReplaceAll(string(modifiedContent), "(cap)", "(cap,1)")
		sliceContent := strings.Split(modifiedContent, " ")
		for i := 0; i < len(sliceContent); i++ {
			if (sliceContent[i][0] >= '0' && sliceContent[i][0] <= '9') && sliceContent[i][len(sliceContent[i])-1] == ')' {
				sliceContent[i-1]+= sliceContent[i]
				sliceContent = append(sliceContent[:i], sliceContent[i+1:]...)
			}
		}


		for _, v := range sliceContent {
			fmt.Printf("%s|",v)
		}
		fmt.Printf("\n")
	}
}
