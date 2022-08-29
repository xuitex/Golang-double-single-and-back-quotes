package main

import "fmt

func main() {
	// double vs back quotes
	doubleQuotes := "Hello\n Go!"
	backQuotes := `Hello\n Go!`

	fmt.Printf("doubleQuotes: %s, type: %T\n", doubleQuotes, doubleQuotes)
	fmt.Printf("backQuotes: %s, type: %T\n", backQuotes, backQuotes)

	/* output
	doubleQuotes: Hello
	Go!, type: string
	backQuotes: Hello\n Go!, type: string
	*/

	// error: more than one character in rune literal
	// singleQuotes := 'Hello\n Go!'
	
	// double vs back quotes
	escapeCharactersIndoubleQuotes := "\n"
	escapeCharactersInbackQuotes := `\n`

	fmt.Printf("escape characters in double quotes: %s, length: %d\n", escapeCharactersIndoubleQuotes, len(escapeCharactersIndoubleQuotes))
	fmt.Printf("escape characters in back quotes: %s, length: %d\n", escapeCharactersInbackQuotes, len(escapeCharactersInbackQuotes))
}
