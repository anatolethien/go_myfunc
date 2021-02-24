package myfunc

import (
	"errors"
	"os"
)

func PrintRune(r rune) error {

	// Verifies if the rune is an ASCII character
	if r < 0 || r > 127 {
		return errors.New("invalid rune")
	}

	// Creating the byte containing the character
	var p = make([]byte, 1)
	p[0] = byte(r)

	// Writing the character
	var _, err = os.Stdout.Write(p)

	return err
}

func PrintString(s string) error {

	var err error

	for _, v := range s {
		err = PrintRune(v)
	}

	return err
}

func Printf()  {
	
}
