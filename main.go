package main

import (
	"fmt"
	"flag"
)



func main () {

	flagLength := flag.Int("length", 20, "password length")
	flagNumber := flag.Bool("n", false, "Number")
	flagUppercase := flag.Bool("u", false, "UpperCase")
	flagLowercase := flag.Bool("l", false, "LoweCase")
	flagSpecial := flag.Bool("s", false, "special characters")
	flag.Parse()



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
		length: *flagLength,
	}

	flagsMap := map[charsets]bool{
		number:*flagNumber,
		uppercase:*flagUppercase,
		lowercase:*flagLowercase,
		special:*flagSpecial,
	}
	fmt.Println("[system]> length: ", *flagLength)

	for key, value := range flagsMap {
		if value {
			passwd.types = append(passwd.types, key)
		}
	}
	
	fmt.Println("[SYSTEM]> Generating password ...")
	output := passwd.generate()
	fmt.Println("[PASSWORD]> " + output)
	fmt.Println("[SYSTEM]> Closing system")

}