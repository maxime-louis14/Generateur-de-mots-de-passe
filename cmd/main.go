package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func generatePassword(length int) string {
	lowercaseChars := []rune("abcdefghijklmnopqrstuvwxyz")
	uppercaseChars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digits := []rune("0123456789")
	specialChars := []rune("~`!@#$%^&*()_-+={[}]|:;»‘<,>.?/")

	allChars := append(lowercaseChars, uppercaseChars...)
	allChars = append(allChars, digits...)
	allChars = append(allChars, specialChars...)

	if length < 4 || length > 40 {
		fmt.Println("La longueur du mot de passe doit être entre 4 et 40 caractères.")
		return ""
	}

	password := make([]rune, length)
	for i := range password {
		randomIndex := randInt(len(allChars))
		password[i] = allChars[randomIndex]
	}
	return string(password)
}

func randInt(max int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(n.Int64())
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrez la longueur du mot de passe : ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	length, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("La longueur doit être un nombre entier.")
		return
	}

	password := generatePassword(length)
	if password != "" {
		fmt.Printf("Mot de passe généré (%d caractères) : %s\n", length, password)
	}
}
