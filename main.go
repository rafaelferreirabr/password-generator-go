package main

import (
	"fmt"
	"flag"
)



func main () {

	flagLength := flag.Int("length", 20, "password length")
	FlagNumber := flag.Bool("n", false, "Number")
	flagUppercase := flag.Bool("u", false, "UpperCase")
	flagLowercase := flag.Bool("l", false, "LoweCase")
	flagSpecial := flag.Bool("s", false, "special characters")
	flag.Parse()

	fmt.Println("[SYSTEM]> Starting Password Generator ...")
	fmt.Println("length: ", *flagLength)
	fmt.Println("number: ", *FlagNumber)
	fmt.Println("Uppercase: ", *flagUppercase)
	fmt.Println("flagLowercase: ", *flagLowercase)
	fmt.Println("flagSpecial: ", *flagSpecial)
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