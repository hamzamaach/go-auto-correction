package reloaded

func Capitalize(s string) string {
	result := ""
	makeNextUpper := true
	for _, char := range s {
		if (char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') {
			if makeNextUpper {
				if char >= 'a' && char <= 'z' {
					result += string(char - ('a' - 'A'))
				} else {
					result += string(char)
				}
				makeNextUpper = false
			} else {
				if char >= 'A' && char <= 'Z' {
					result += string(char + 32)
				} else {
					result += string(char)
				}
			}
		} else {
			result += string(char)
			makeNextUpper = true
		}
	}
	return result
}
