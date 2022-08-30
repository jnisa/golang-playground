
package main

import "math/rand"

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

func main() {
	get_fixed_length_password(12, "abcdefg1234")
}