package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func addSpaceAfterPunctuation(str string) string {
	var modifiedContent strings.Builder
	length := len(str)

	for i := 0; i < length; i++ {
		char := str[i]
		modifiedContent.WriteByte(char)

		// Check if the character is one of the specified punctuation marks
		if char == '.' || char == ',' || char == '!' || char == '?' || char == ':' || char == ';' {
			// Check if the next character is not a space and exists
			if i < length-1 && str[i+1] != ' ' {
				modifiedContent.WriteByte(' ')
			}
		}
	}

	return modifiedContent.String()
}

func stringToSlice(str string) []string {
	// Unification of form of instructions
	var modifiedContent string
	modifiedContent = strings.ReplaceAll(str, "(up)", "(up,1)")
	modifiedContent = strings.ReplaceAll(modifiedContent, "(low)", "(low,1)")
	modifiedContent = strings.ReplaceAll(modifiedContent, "(cap)", "(cap,1)")
	sliceContent := strings.Split(modifiedContent, " ")

	// remove the space in <<(low, 3)>> to became <<(low,3)>>
	for i := 0; i < len(sliceContent); i++ {
		if (sliceContent[i][0] >= '0' && sliceContent[i][0] <= '9') && sliceContent[i][len(sliceContent[i])-1] == ')' {
			sliceContent[i-1] += sliceContent[i]
			sliceContent = append(sliceContent[:i], sliceContent[i+1:]...)
		}
	}

	return sliceContent
}

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

func FixIndefiniteArticles(str []string) []string {
	updatedString := []string{}
	for i, v := range str {
		if v == "a" || v == "A" {
			if str[i+1][0] == 'a' ||
				str[i+1][0] == 'e' ||
				str[i+1][0] == 'o' ||
				str[i+1][0] == 'i' ||
				str[i+1][0] == 'h' ||
				str[i+1][0] == 'u' {
				if v == "a" {
					v = "an"
				}
				if v == "A" {
					v = "An"
				}
			}
		}
		updatedString = append(updatedString, v)
	}
	return updatedString
}

func ProcessContentActions(sliceContent []string) []string {
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
	return sliceContent
}

func removeSpaces(sliceString []string) []string {
	var updatedStrings []string
	// Remove all spaces from the string
	for _, str := range sliceString {
		updatedStr := strings.ReplaceAll(str, " ", "")
		updatedStrings = append(updatedStrings, updatedStr)
	}

	return updatedStrings
}

func handleSpaces(sliceString []string) []string {
	openQuote := false
	for i := 0; i < len(sliceString)-1; i++ {
		if sliceString[i+1][0] == '.' ||
			sliceString[i+1][0] == ',' ||
			sliceString[i+1][0] == '!' ||
			sliceString[i+1][0] == '?' ||
			sliceString[i+1][0] == ':' ||
			sliceString[i+1][0] == ';' {
			continue
		} else if sliceString[i+1][0] == '\'' {
			if !openQuote {
				continue
			} else {
				sliceString[i] += " "
			}
		} else {
			sliceString[i] += " "
		}
		if sliceString[i][0] == '\'' && !openQuote {
			sliceString[i] = strings.Trim(sliceString[i], " ")
		}
	}
	return sliceString
}
func addSpacesAfterSymbols(str string) string {
	result := ""
	i := 0

	for i < len(str) {
		if str[i] == '.' ||
			str[i] == ',' ||
			str[i] == '!' ||
			str[i] == '?' ||
			str[i] == ':' ||
			str[i] == ';' {
			result += string(str[i])

			start := i + 1
			for start < len(str) && str[start] == str[i] {
				result += string(str[start])
				start++
			}
			i = start

			// if i < len(str) && str[i] != ' ' {
			// 	result += " "
			// }
		} else {
			result += string(str[i])
			i++
		}
	}

	return result
}

// func addSpacesAfeterSymbols(str string) string {
// 	result := ""
// 	for i, char := range str {
// 		if (char == '.' ||
// 			char == ',' ||
// 			char == '!' ||
// 			char == '?' ||
// 			char == ':' ||
// 			char == ';') && i != len(str)-1 {
// 			if str[i+1] != ' ' {
// 				result += string(char)
// 			}
// 		} else {
// 			result += string(char)
// 		}
// 	}
// 	return result
// }

func SaveFile(fileName string, str string) {
	file, _ := os.Create(fileName)
	defer file.Close()
	data := []byte(str)
	file.Write(data)
}

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		result := ""
		content, _ := ioutil.ReadFile(args[0])
		stringContent := addSpaceAfterPunctuation(string(content))
		sliceContent := stringToSlice(stringContent)
		sliceContent = ProcessContentActions(sliceContent)
		sliceContent = FixIndefiniteArticles(sliceContent)
		sliceContent = removeSpaces(sliceContent)
		updatedString := handleSpaces(sliceContent)
		for _, v := range updatedString {
			result += v
		}
		result = addSpacesAfterSymbols(result)
		// fmt.Println(result)
		SaveFile("result.txt", result)

	}
}
