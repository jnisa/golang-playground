
package main

import "testing"

const AVAILABLE_CHARS = NUMBERS + SYMBOLS + LOWERCASE_LETTERS + UPPERCASE_LETTERS

func Test_select_password_features(t *testing.T){
	t.Run("low complexity: just consider symbols", func(t *testing.T){
		result := select_password_features(false, false, true, false)
		expected := "0123456789"

		if result != expected {
			t.Errorf("TEST FAILED (select_password_features): the chars selected don't match!")
		}
	})
}

func Test_get_fixed_length_password(t *testing.T){
	t.Run("low complexity: ask for a 12 characters password", func(t *testing.T) {
		result := get_fixed_length_password(12, AVAILABLE_CHARS)
		expected := 12

		if len(result) != expected {
			t.Errorf("TEST FAILED (get_fixed_lenght_password): the sizes don't match!")
		}
	})

	t.Run("low complexity: ask for 25 character password", func(t *testing.T) {
		result := get_fixed_length_password(25, AVAILABLE_CHARS)
		expected := 25

		if len(result) != expected {
			t.Errorf("TEST FAILED (get_fixed_length_password): the sizes don't match!")
		}
	})
}