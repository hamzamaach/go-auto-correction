package reloaded

import (
	"os"
)

func SaveFile(fileName string, str string) {
	file, _ := os.Create(fileName)
	defer file.Close()
	data := []byte(str+"\n")
	file.Write(data)
}
