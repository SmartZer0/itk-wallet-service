package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

func main() {

	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	printType("numDecimal", numDecimal)
	printType("numOctal", numOctal)
	printType("numHexadecimal", numHexadecimal)
	printType("pi", pi)
	printType("name", name)
	printType("isActive", isActive)
	printType("complexNum", complexNum)

	joined := joinToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("Объединённая строка:", joined)

	runes := []rune(joined)
	fmt.Println("Количество рун:", utf8.RuneCountInString(joined))

	hashed := hashWithSalt(runes, "go-2024")
	fmt.Println("SHA256 хеш:", hashed)
}

func printType(name string, val interface{}) {
	fmt.Printf("Тип %s: %s\n", name, reflect.TypeOf(val).String())
}

func joinToString(values ...interface{}) string {
	var sb strings.Builder
	for _, v := range values {
		sb.WriteString(fmt.Sprintf("%v", v))
	}
	return sb.String()
}

func hashWithSalt(runes []rune, salt string) string {
	mid := len(runes) / 2
	runesWithSalt := append(runes[:mid], append([]rune(salt), runes[mid:]...)...)
	hashed := sha256.Sum256([]byte(string(runesWithSalt)))
	return hex.EncodeToString(hashed[:])
}
