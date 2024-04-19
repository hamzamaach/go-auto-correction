package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Convert a string to uppercase
func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

// Convert a string to lowercase
func toLowerCase(s string) string {
	return strings.ToLower(s)
}

// Capitalize a string
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		content, _ := ioutil.ReadFile(args[0])

		// Unification of form of instructions
		var modifiedContent string
		modifiedContent = strings.ReplaceAll(string(content), "(up)", "(up,1)")
		modifiedContent = strings.ReplaceAll(string(modifiedContent), "(low)", "(low,1)")
		modifiedContent = strings.ReplaceAll(string(modifiedContent), "(cap)", "(cap,1)")
		sliceContent := strings.Split(modifiedContent, " ")

		// remove the space in <<(low, 3)>> to became <<(low,3)>>
		for i := 0; i < len(sliceContent); i++ {
			if (sliceContent[i][0] >= '0' && sliceContent[i][0] <= '9') && sliceContent[i][len(sliceContent[i])-1] == ')' {
				sliceContent[i-1] += sliceContent[i]
				sliceContent = append(sliceContent[:i], sliceContent[i+1:]...)
			}
		}

		for i := 0; i < len(sliceContent); i++ {
			if strings.Contains(sliceContent[i], "(") && strings.Contains(sliceContent[i], ")") {
				if strings.Contains(sliceContent[i], ",") {
					withoutParentheses := sliceContent[i][1 : len(sliceContent[i])-1]
					options := strings.Split(withoutParentheses, ",")
					if len(options) == 2 {
						action := options[0]
						howMany, err := strconv.Atoi(options[1])
						if err == nil {
							for j := i - 1; j >= i-howMany; j-- {
								if j >= 0 {
									switch action {
									case "up":
										sliceContent[j] = toUpperCase(sliceContent[j])
									case "low":
										sliceContent[j] = toLowerCase(sliceContent[j])
									case "cap":
										sliceContent[j] = capitalize(sliceContent[j])
									}
								}
							}
						}
						sliceContent = append(sliceContent[:i], sliceContent[i+1:]...)
						i--
					}
				} else {
					action := sliceContent[i][1 : len(sliceContent[i])-1]
					switch action {
					case "hex":
						// Convert hex string to int64
						hexValue, err := strconv.ParseInt(sliceContent[i-1], 16, 64)
						if err == nil {
							// Convert int64 to string and update the slice
							sliceContent[i-1] = strconv.FormatInt(hexValue, 10)
						}
					case "bin":
						// Convert binary string to int64
						binValue, err := strconv.ParseInt(sliceContent[i-1], 2, 64)
						if err == nil {
							// Convert int64 to string and update the slice
							sliceContent[i-1] = strconv.FormatInt(binValue, 10)
						}
					}
					sliceContent = append(sliceContent[:i], sliceContent[i+1:]...)
					i--
				}
			}
		}

		for _, v := range sliceContent {
			fmt.Printf("%s ", v)
		}
		fmt.Printf("\n")
	}
}
