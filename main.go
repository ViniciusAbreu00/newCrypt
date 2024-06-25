package main

import (
	"bufio"
	"fmt"
	"newCrypt/encryption"
	"os"
	"strings"

	"github.com/sqweek/dialog"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Would you like to encrypt or decrypt a file? (e/d): ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	fmt.Println("Enter the password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	filepath, err := dialog.File().Filter("Text files", "txt").Load()
	if err != nil {
		fmt.Println("Error loading file", err)
		return
	}

	switch choice {
	case "e":
		err := encryption.EncryptFile(filepath, password)
		if err != nil {
			fmt.Println("Error encryptin file", err)
			return
		}
	case "d":
		err := encryption.DecryptFile(filepath, password)
		if err != nil {
			fmt.Println("Error decrypting file", err)
			return
		}
	default:
		fmt.Println("Unknown choice")
	}

}
