package main

import (
	"log"
	"ricer/internal/filesys"
)

func main() {
	themes := filesys.GetThemes()
	if len(themes) == 0 {
		log.Panic("No themes")
	}
	filesys.SubmitTheme(themes[1])
}
