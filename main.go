package main

import "fmt"



func main () {
	fmt.Println("[SYSTEM]> Starting Password Generator ...")
	number := charsets{
		min: 48,
		max: 57,
	}
	uppercase := charsets{
		min: 65,
		max: 90,
	}
	lowercase := charsets{
		min: 97,
		max: 122,
	}
	special := charsets{
		min: 33,
		max: 46,
	}

	passwd := password{
		length: 30,
	}
	passwd.types = append(passwd.types, number)
	passwd.types = append(passwd.types, lowercase)
	passwd.types = append(passwd.types, uppercase)
	passwd.types = append(passwd.types, special)

	
	fmt.Println("[SYSTEM]> Generating password ...")
	output := passwd.generate()
	fmt.Println("[PASSWORD]> " + output)
	fmt.Println("[SYSTEM]> Closing system")

}