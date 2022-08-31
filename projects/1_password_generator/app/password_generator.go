
package main

import "fmt"
import "math/rand"

const NUMBERS = "0123456789"
const SYMBOLS = "~`!@#$%^&*()-_+={}[]|/:;'<>,.?"
const LOWERCASE_LETTERS = "abcdefghijklmnopqrstuvwxyz"
const UPPERCASE_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func get_fixed_length_password(length int, chars string) string {
    // Determines the length of the password that will be generated.

    // :param length: size that the password should have
    // :param chars: type of chars that can be used to create the password 
    // :return string: password string 

    password := make([]byte, length)
    for idx := range password {
        password[idx] = chars[rand.Intn(len(chars))]
    }
    return string(password)
}

func select_password_features(
	has_uppercase_letters bool,
	has_lowercase_letters bool,
	has_numbers bool,
	has_symbols bool) string {
	// Defines the features that the password must have.

	// :param has_uppercase_letters: enable the inclusion of uppercase letters
	// :param has_lowercase_letters: enable the inclusion of lowercase letters
	// :param has_numbers: enable the inclusion of numbers
	// :param has_symbols: enable the inclusion of symbols
	// :return string: string with all the chars enabled  

	allowed_characters := ""

	if has_uppercase_letters {
		allowed_characters += UPPERCASE_LETTERS
	}

	if has_lowercase_letters {
		allowed_characters += LOWERCASE_LETTERS
	}

	if has_numbers {
		allowed_characters += NUMBERS
	}

	if has_symbols {
		allowed_characters += SYMBOLS
	}
	return allowed_characters
}

func main() {
	// ask the user what features he needs on his password
	var password_length int
	var has_uppercase_letters bool
	var has_lowercase_letters bool
	var has_numbers bool
	var has_symbols bool

	fmt.Println("Password Length: ")
	fmt.Scanln(&password_length)
	fmt.Println("Uppercase Letters: ")
	fmt.Scanln(&has_uppercase_letters)
	fmt.Println("Lowercase Letters: ")
	fmt.Scanln(&has_lowercase_letters)
	fmt.Println("Numbers: ")
	fmt.Scanln(&has_numbers)
	fmt.Println("Symbols: ")
	fmt.Scanln(&has_symbols)

	available_chars := select_password_features(has_uppercase_letters, has_lowercase_letters, has_numbers, has_symbols)
	password_generated := get_fixed_length_password(password_length, available_chars)

	fmt.Printf("Password Generated: %s\n", password_generated)
}