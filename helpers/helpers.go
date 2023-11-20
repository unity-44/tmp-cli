package helpers

import (
	"log"
	"os"
	"os/exec"
)

// Creates a file with name from flag
func CreateFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

// Open file in VS code
func OpenFileInCode(filePath string) {
	cmd := exec.Command("code", filePath)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
