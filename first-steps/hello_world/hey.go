

package main

// list of packages that are used on this script
// 1. fmt - package that allows the usage of function like Println
import "fmt"

const eng_hello_prefix = "Hello, "
const pt_hello_prefix = "Olá, "
const fr_hello_prefix = "Bonjour, "

func hello_world(name string, idiom string) string {
	if name == "" {
		name = "world"
	} 

	return defineMessage(idiom) + name
}


func defineMessage(idiom string) (message string) {
	switch idiom {
	case "portuguese":
		message = pt_hello_prefix 

	case "french":
		message = fr_hello_prefix 		 

	case "english":
		message = eng_hello_prefix
	}

	return message
}


func main(){
	fmt.Println(hello_world("Nisa", "french"))
	fmt.Println(hello_world("Nisa", "english"))
	fmt.Println(hello_world("", "portuguese"))
}
