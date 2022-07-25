

package main

import "testing"

func Test_sum(t *testing.T) {
	t.Run("basic sum test", func(t *testing.T) {
		result := sum(2, 3)
		expected := 5

		if result != expected {
			t.Errorf("YOU'RE DUMB - Result: %d | Expected: %d", result, expected)
		}
	})
}
