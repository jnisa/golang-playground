
package main

import "testing"

func Test_hello_world_1(t *testing.T){
	t.Run("test when a name is provided", func(t *testing.T) {
		result := hello_world("Nisa", "english")
		expected := "Hello, Nisa"

		if result != expected {
			t.Errorf("TEST FAILED - Result: '%s' | Expected: '%s'", result, expected)
		}
	})

	t.Run("test when a name is not provided", func(t *testing.T) {
		result := hello_world("", "english")
		expected := "Hello, world"

		if result != expected {
			t.Errorf("TEST FAILED - Result: %s | Expected: %s", result, expected)
		}
	})

	t.Run("when portuguese message required", func(t *testing.T) {
		result := hello_world("Nisa", "portuguese")
		expected := "Olá, Nisa"

		if result != expected {
			t.Errorf("TEST FAILED - Result %s | Expected: %s", result, expected)
		}
	})

	t.Run("when french message required", func(t *testing.T) {
		result := hello_world("", "french")
		expected := "Bonjour, world"

		if result != expected {
			t.Errorf("TEST FAILED - Result %s | Expected: %s", result, expected)
		}
	})
}
