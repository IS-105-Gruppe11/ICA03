package iso

func IterateOverASCIIStringLiteral(ascii []byte) {
	for i := 0; i < len(ascii); i++ {
		fmt.Printf("%X %q %b \n", ascii[i], ascii[i], ascii[i])

}

// Kode for Oppgave 2b
func GreetingExtendedASCII() {}
