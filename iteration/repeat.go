package iteration

import "fmt"

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
		//keluarkan value dlu
		fmt.Printf("i = %d, repeated: %q\n", i, repeated)
		//i++ run last skali
	}
	return repeated
}
