
package main

const NUMBERS = "0123456789"
const SYMBOLS = "~`!@#$%^&*()-_+={}[]|/:;'<>,.?"
const LOWERCASE_LETTERS = "abcdefghijklmnopqrstuvwxyz"
const UPPERCASE_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

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
	select_password_features(false, false, false, false)
}