// comment
package main  

import "fmt"

func main() {
	noDeclareType := 'c'
	fmt.Printf("%T\n", noDeclareType)

	var declareTypeByte byte = 'c'
	fmt.Printf("%T\n", declareTypeByte)

	var declareTypeRune rune = 'c'
	fmt.Printf("%T\n", declareTypeRune)
}
