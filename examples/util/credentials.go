package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func Credentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}
	username = strings.TrimRight(username, "\r\n")
	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	}
	fmt.Print("\n")
	password := string(bytePassword)
	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}

func Totp() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter TOTP: ")
	totp, err := reader.ReadString('\n')
	if err != nil {
		return totp, err
	}
	totp = strings.TrimRight(totp, "\r\n")
	return totp, nil
}
