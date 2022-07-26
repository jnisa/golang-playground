
package arraySum

import "testing"

func Test_arrayIntSum(t *testing.T){
	t.Run("First Test: low complexity test", func(t *testing.T){
		result := arrayIntSum([]int{2, 4, 6})
		expected := 12

		if result != expected{
			t.Errorf("TEST FAILED - Result: %d | Expected: %d", result, expected)
		}
	})

	t.Run("Second Test: empty array provided", func(t *testing.T){
		result := arrayIntSum([]int{})
		expected := 0

		if result != expected {
			t.Errorf("TEST FAILED - Result: %d | Expected: %d", result, expected)
		}
	})
}

func Test_arrayStrSum(t *testing.T){
	t.Run("First Test: low complexity test", func(t *testing.T){
		result := arrayStrSum([]string{"abc", "def", "ghi"})
		expected := "abcdefghi"

		if result != expected {
			t.Errorf("TEST FAILED - Result: %s | Expected: %s", result, expected)
		}
	})

	t.Run("Second Test: empty array provided", func(t *testing.T){
		result := arrayStrSum([]string {})
		expected := ""

		if result != expected{
			t.Errorf("TEST FAILED - Result: %s | Expected: %s", result, expected)
		}
	})
}
