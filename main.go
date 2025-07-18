package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func IsUnique(s string) bool {
	cartina := make(map[rune]struct{})
	runification := []rune(s)

	for _, value := range runification {

		if unicode.IsSpace(value) {
			continue
		}

		if _, ok := cartina[value]; ok {
			return false
		}
		cartina[value] = struct{}{}
	}

	return true
}

func main() {

	fmt.Println("Введите строку")
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	line = strings.ToLower(line)

	result := IsUnique(line)

	fmt.Printf("Все символы уникальны?: %v\n", result)
}
