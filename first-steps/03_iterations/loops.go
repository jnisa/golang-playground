
package iterations


func repeatedString(input string, reps_number int) string {
	var repeated_string string
	for i := 0; i < reps_number; i++ {
		repeated_string+=input
	}
	return repeated_string
}
