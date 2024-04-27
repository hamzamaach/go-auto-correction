package main

import (
	"fmt"
	"os"

	ft "reloaded/features"
)

func main() {
	args := os.Args[1:]
	argsErr := ft.VerifyArgs(args)
	if !argsErr {
		content, fileErr := os.ReadFile(args[0])
		if fileErr == nil {
			StringContent := string(content)
			StringContent = ft.HandleIndefiniteArticles(StringContent)
			StringContent = ft.FormatText(StringContent)
			StringContent = ft.ProcessContentActions(StringContent)
			StringContent = ft.AdjustWhitespacesAfterSymbols(StringContent)
			StringContent = ft.AddSpacesAroundSymbols(StringContent)
			StringContent = ft.AdjustSingleQuotes(StringContent)
			StringContent = ft.AdjustWhitespacesBeforeSymbols(StringContent)
			ft.SaveFile(args[1], StringContent)
		} else {
			fmt.Print("Error: ", fileErr)
			fmt.Print("\n")
		}
	}
}
