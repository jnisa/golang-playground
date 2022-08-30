
package main

import "testing"

func Test_select_password_features(t *testing.T){
	t.Run("low complexity: just consider symbols", func(t *testing.T){
		result := select_password_features(false, false, true, false)
		expected := "0123456789"

		if result != expected {
			t.Errorf("TEST FAILED (select_password_features): the chars selected don't match!")
		}
	})
}