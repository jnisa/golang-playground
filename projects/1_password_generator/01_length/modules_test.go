
package main

import "fmt"
import "testing"


func Test_get_fixed_length_password(t *testing.T){
	t.Run("low complexity: compare lengths", func(t *testing.T) {
		result := get_fixed_length_password(12)
		expected := 12

		fmt.Printf("The obtained password: %#v\n", result)

		if len(result) != expected {
			t.Errorf("TEST FAILED (test_get_fixed_lenght_password): the sizes don't match!")
		}
	})
}