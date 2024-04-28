package reloaded

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func convertToDecimal(input string, base int) (string, error) {
	value, err := strconv.ParseInt(input, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(value, 10), nil
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
	s = toLowerCase(s)
	// capitalize first character and concatenates with the rest of the string
	return strings.ToUpper(string(s[0])) + s[1:]
}

func ExtractInsideParentheses(input string) string {
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

func HandleIndefiniteArticles(str string) string {
	lines := strings.Split(str, "\n")
	vowels := "aeiouhAEIOUH"
	modifiedLines := []string{}

	for _, line := range lines {
		words := strings.Fields(line)
		for i := 0; i < len(words); i++ {
			if (words[i] == "a" || words[i] == "A") && i+1 < len(words) {
				if words[i] == "a" {
					if strings.Contains(vowels, string(words[i+1][0])) {
						words[i] = "an"
					}
				} else if words[i] == "A" {
					if strings.Contains(vowels, string(words[i+1][0])) {
						words[i] = "An"
					}
				}
			}
		}
		modifiedLines = append(modifiedLines, strings.Join(words, " "))
	}
	result := strings.Join(modifiedLines, "\n")
	return result
}

func FormatText(str string) string {
	// Unification of form of instructions
	str = strings.ReplaceAll(str, "(up)", "(up,1)")
	str = strings.ReplaceAll(str, "(low)", "(low,1)")
	str = strings.ReplaceAll(str, "(cap)", "(cap,1)")
	// add whitespace before and after each action
	str = strings.ReplaceAll(str, ")", ") ")
	str = strings.ReplaceAll(str, "(", " (")
	lines := strings.Split(str, "\n")
	modifiedLines := []string{}
	for _, line := range lines {
		words := strings.Fields(line)
		// remove the space in <<(low, 3)>> to became <<(low,3)>>
		for i := 0; i < len(words); i++ {
			if (words[i][0] >= '0' && words[i][0] <= '9') && words[i][len(words[i])-1] == ')' {
				words[i-1] += words[i]
				words = append(words[:i], words[i+1:]...)
			}
		}
		modifiedLines = append(modifiedLines, strings.Join(words, " "))
	}
	result := strings.Join(modifiedLines, "\n")
	return result
}

// Processes actions (converting to hex, bin, up, low, or cap)
func ProcessContentActions(str string) string {
	lines := strings.Split(str, "\n")
	modifiedLines := []string{}
	for _, line := range lines {
		words := strings.Fields(line)
		i := 0
		for i < len(words) {
			word := words[i]
			// Check if word contains an action
			if strings.Contains(word, "(") && strings.Contains(word, ")") && (strings.Contains(word, "hex") ||
				strings.Contains(word, "bin") ||
				strings.Contains(word, "up") ||
				strings.Contains(word, "low") ||
				strings.Contains(word, "cap")) {
				// If word contains a "," and is one of the actions "bin" or "hex" do nothing
				if strings.Contains(word, ",") && (strings.Contains(word, "bin") || strings.Contains(word, "hex")) {
					i++
					continue
				}
				if strings.Contains(word, ",") {
					// Handle case with multiple actions
					withoutParentheses := ExtractInsideParentheses(word)
					options := strings.Split(withoutParentheses, ",")
					if len(options) != 2 {
						i++
						continue
					}
					action := options[0]
					howMany, err := strconv.Atoi(options[1])
					if err != nil {
						fmt.Println("Error:", err)
						i++
						continue
					}
					// Apply the action to the specified number of words
					for j := i - 1; j >= i-howMany && j >= 0; j-- {
						if !isOnlycharacters(words[j]) {
							howMany++
							continue
						}
						if j >= 0 {
							switch action {
							case "up":
								words[j] = toUpperCase(words[j])
							case "low":
								words[j] = toLowerCase(words[j])
							case "cap":
								words[j] = capitalize(words[j])
							}
						}
					}
					words = append(words[:i], words[i+1:]...)
					i--
				} else {
					action := ExtractInsideParentheses(word)
					var base int
					switch action {
					case "hex":
						base = 16
					case "bin":
						base = 2
					default:
						i++
						continue
					}
					if i-1 >= 0 {
						decimalStr, err := convertToDecimal(words[i-1], base)
						if err == nil {
							words[i-1] = decimalStr
						}
					}
					words = append(words[:i], words[i+1:]...)
					i--
				}
			}
			i++
		}
		modifiedLines = append(modifiedLines, strings.Join(words, " "))
	}
	return strings.Join(modifiedLines, "\n")
}

// Adds spaces after punctuation marks (.,!?:;) and handles repeated punctuation correctly
func AdjustWhitespacesBeforeSymbols(str string) string {
	lines := strings.Split(str, "\n")
	modifiedLines := []string{}

	for _, line := range lines {
		words := strings.Fields(line)

		for i := 1; i < len(words); i++ {
			if len(words[i]) > 0 && strings.Contains(".,!?;:", string(words[i][0])) {
				words[i-1] += words[i]
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}

		modifiedLines = append(modifiedLines, strings.Join(words, " "))
	}

	result := strings.Join(modifiedLines, "\n")
	return result
}

func AdjustWhitespacesAfterSymbols(str string) string {
	var modifiedContent strings.Builder
	length := len(str)

	for i := 0; i < length; i++ {
		modifiedContent.WriteByte(str[i])

		if str[i] == '.' ||
			str[i] == ',' ||
			str[i] == '!' ||
			str[i] == '?' ||
			str[i] == ':' ||
			str[i] == ';' {
			// Check if the next character is not a space and exists
			if i < length-1 && str[i+1] != ' ' {
				modifiedContent.WriteByte(' ')
			}
		}
		// add space after single quote if there's a space before it

		if str[i] == '\'' || str[i] == '"' {
			if i > 0 && str[i-1] == ' ' {
				if i < length-1 && str[i+1] != ' ' {
					modifiedContent.WriteByte(' ')
				}
			}
		}
	}
	return modifiedContent.String()
}

func AdjustQuotes(str string) string {
	lines := strings.Split(str, "\n")
	modifiedLines := []string{}

	for _, line := range lines {
		words := strings.Fields(line)
		openSingleQuote := false
		openDoubleQuote := false

		for i := 0; i < len(words); i++ {
			if words[i] == "'" {
				if openSingleQuote {
					if i > 0 {
						words[i-1] += words[i]
						words = append(words[:i], words[i+1:]...)
						i--
					}
					openSingleQuote = false
				} else {
					if i < len(words)-1 {
						words[i] = words[i] + words[i+1]
						words = append(words[:i+1], words[i+2:]...)
					}
					openSingleQuote = true
				}
			} else if words[i] == "\"" {
				if openDoubleQuote {
					if i > 0 {
						words[i-1] += words[i]
						words = append(words[:i], words[i+1:]...)
						i--
					}
					openDoubleQuote = false
				} else {
					if i < len(words)-1 {
						words[i] = words[i] + words[i+1]
						words = append(words[:i+1], words[i+2:]...)
					}
					openDoubleQuote = true
				}
			}
		}
		modifiedLines = append(modifiedLines, strings.Join(words, " "))
	}

	result := strings.Join(modifiedLines, "\n")
	return result
}

func AddSpacesAroundSymbols(input string) string {
	var result strings.Builder
	symbols := ".,!?;:"

	// Iterate through each character in the input string
	for i, char := range input {
		if strings.ContainsRune(symbols, char) {
			// Add space before the symbol if necessary
			if i > 0 && !unicode.IsSpace(rune(input[i-1])) {
				result.WriteString(" ")
			}
			// Add the symbol
			result.WriteRune(char)
			// Add space after the symbol if necessary
			if i < len(input)-1 && !unicode.IsSpace(rune(input[i+1])) {
				result.WriteString(" ")
			}
		} else {
			// Add the character as is
			result.WriteRune(char)
		}
	}

	return result.String()
}

func isOnlycharacters(word string) bool {
	for _, char := range word {
		if strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", char) {
			return true
		}
	}
	return false
}

func AddSpaceAfterSingleQuote(str string) string {
	var modifiedContent strings.Builder
	length := len(str)

	i := 0
	for i < length {
		currentChar := str[i]
		modifiedContent.WriteByte(currentChar)

		if currentChar == '\'' || currentChar == '"' {
			if (i == 0 || str[i-1] == ' ') && i < length-1 {
				nextChar := str[i+1]
				if nextChar != ' ' {
					// Add a space after the single quote
					modifiedContent.WriteByte(' ')
				}
			}
		}
		i++
	}

	return modifiedContent.String()
}
