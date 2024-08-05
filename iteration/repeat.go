package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
		//keluarkan value dlu
		//fmt.Printf("i = %d, repeated: %q\n", i, repeated)
		//i++ run last skali
	}
	return repeated
}
