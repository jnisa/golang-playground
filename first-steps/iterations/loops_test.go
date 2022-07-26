

package iterations

import "testing"

func Test_repeatedString(t *testing.T) {
	t.Run("First Test: low complexity test", func(t *testing.T){
		result := repeatedString("Nisa", 3)
		expected := "NisaNisaNisa"	
		
		if result != expected {
			t.Errorf("TEST FAILED - Result: %s | Expected: %s", result, expected)
		}
	})

	t.Run("Second Test: when no repetitions number is provided", func(t *testing.T){
		result := repeatedString("abc", 0)
		expected := ""

		if result != expected {
			t.Errorf("TEST FAILED - Result: %s | Expected: %s", result, expected)
		}
	})
}


// perform some benchmarks
func Benchmark_repeatedString(b *testing.B) {
	for i := 0; i < b.N; i++{
		repeatedString("a", 10000)
	} 
}
