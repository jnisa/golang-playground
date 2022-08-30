
package main

import "math/rand"

const NUMBERS = "0123456789"
const SPECIAL_CHARS = "~`!@#$%^&*()-_+={}[]|/:;'<>,.?"
const LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const AVAILABLE_CHARS = NUMBERS + SPECIAL_CHARS + LETTERS

func get_fixed_length_password(length int) string {
    password := make([]byte, length)
    for idx := range password {
        password[idx] = AVAILABLE_CHARS[rand.Intn(len(AVAILABLE_CHARS))]
    }
    return string(password)
}

func main() {
	get_fixed_length_password(12)
}