

package arraySum

func arrayIntSum(input []int) int {
	var sum_value int
	for _, entry := range(input) {
		sum_value += entry	
	}
	return sum_value
}

func arrayStrSum(input []string) string{
	var sum_strings string
	for _,entry := range(input){
		sum_strings += entry
	}
	return sum_strings
}
