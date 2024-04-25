package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Add a space after each punctuation mark (.,!?:;)
func addSpaceAfterPunctuation(str string) string {
	var modifiedContent strings.Builder
	length := len(str)

	for i := 0; i < length; i++ {
		modifiedContent.WriteByte(str[i])

		if str[i] == '.' || str[i] == ',' || str[i] == '!' || str[i] == '?' || str[i] == ':' || str[i] == ';' {
			// Check if the next character is not a space and exists
			if i < length-1 && str[i+1] != ' ' {
				modifiedContent.WriteByte(' ')
			}
		}
	}
	return modifiedContent.String()
}

// Converts the input string into a slice of strings (words)
func stringToSlice(str string) []string {
	// Unification of form of instructions
	var modifiedContent string
	modifiedContent = strings.ReplaceAll(str, "\n", "{-n-}")
	modifiedContent = strings.ReplaceAll(modifiedContent, ")", ") ")
	modifiedContent = strings.ReplaceAll(modifiedContent, "(up)", "(up,1)")
	modifiedContent = strings.ReplaceAll(modifiedContent, "(low)", "(low,1)")
	modifiedContent = strings.ReplaceAll(modifiedContent, "(cap)", "(cap,1)")
	sliceContent := strings.Fields(modifiedContent)

	// remove the space in <<(low, 3)>> to became <<(low,3)>>
	for i := 0; i < len(sliceContent); i++ {
		if (sliceContent[i][0] >= '0' && sliceContent[i][0] <= '9') && sliceContent[i][len(sliceContent[i])-1] == ')' {
			sliceContent[i-1] += sliceContent[i]
			sliceContent = append(sliceContent[:i], sliceContent[i+1:]...)
		}
	}

	return sliceContent
}

func extractInsideParentheses(input string) string {
	inParentheses := false
	currentContent := ""

	for _, char := range input {
		if char == '(' {
			inParentheses = true
			continue
		}
		if char == ')' {
			if inParentheses {
				break
			}
		}
		if inParentheses {
			currentContent += string(char)
		}
	}
	return currentContent
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
	// capitalize first character and concatenates with the rest of the string
	s = toLowerCase(s)
	return strings.ToUpper(string(s[0])) + s[1:]
}

/*
v := "aAeE"
strings.Countains(v, string(str[0]))
*/

// Corrects the use of indefinite articles ("a" and "A") to "an" or "An" when followed by a vowel sound
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

// Processes certain actions specified within parentheses in the input slice (such as converting to hex, bin, up, low, or cap)
func ProcessContentActions(sliceContent []string) []string {
	for i := 0; i < len(sliceContent); i++ {
		if strings.Contains(sliceContent[i], "(") &&
			strings.Contains(sliceContent[i], ")") &&
			(strings.Contains(sliceContent[i], "hex") ||
				strings.Contains(sliceContent[i], "bin") ||
				strings.Contains(sliceContent[i], "up") ||
				strings.Contains(sliceContent[i], "low") ||
				strings.Contains(sliceContent[i], "bin") ||
				strings.Contains(sliceContent[i], "cap")) {
			if strings.Contains(sliceContent[i], ",") {

				// remove Parentheses
				// withoutParentheses := sliceContent[i][1 : len(sliceContent[i])-1]
				withoutParentheses := extractInsideParentheses(sliceContent[i])
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
				// remove Parentheses
				// action := sliceContent[i][1 : len(sliceContent[i])-1]
				action := extractInsideParentheses(sliceContent[i])
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

// remove spaces from all the elems of a slice of strings given
func removeSpaces(sliceString []string) []string {
	var updatedStrings []string
	// Remove all spaces from the string
	for _, str := range sliceString {
		updatedStr := strings.ReplaceAll(str, " ", "")
		updatedStrings = append(updatedStrings, updatedStr)
	}

	return updatedStrings
}

// Adjusts white spaces and quote marks in the given slice of strings
func adjustWhitespaceAndQuotes(sliceString []string) []string {
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

// Adds spaces after punctuation marks (.,!?:;) and handles repeated punctuation correctly
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

			// Initialize a variable to handle repeated punctuation
			start := i + 1
			for start < len(str) && str[start] == str[i] {
				result += string(str[start])
				start++
			}
			// Update the index to skip past repeated punctuation
			i = start

			// Add a space after the punctuation
			if i < len(str)-1 && str[i] != ' ' {
				if str[i+1] != '.' &&
					str[i+1] != ',' &&
					str[i+1] != '!' &&
					str[i+1] != '?' &&
					str[i+1] != ':' &&
					str[i+1] != ';' {
					result += " "
				}
			}
		} else {
			result += string(str[i])
			i++
		}
	}

	return result
}

// Saves the modified string content to a specified file
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
		updatedString := adjustWhitespaceAndQuotes(sliceContent)
		for _, v := range updatedString {
			result += v
		}
		result = addSpacesAfterSymbols(result)
		result = strings.ReplaceAll(result, "{-n-}", "\n")
		SaveFile("result.txt", result)
	}
}
