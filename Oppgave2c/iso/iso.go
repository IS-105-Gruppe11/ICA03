package iso

import "fmt"


//Oppgave 2B
func GreetingExtendedASCII() string{
	greetings := []byte{83, 97, 108, 117, 116, 44, 32, 231, 97, 32, 118, 97, 32, 176, 45, 41, 32, 128, 53, 48}
	for i := 0; i < len(greetings); i++ {
		fmt.Printf("%c", greetings[i])
	}
	retur := string(greetings)
	return retur
}
