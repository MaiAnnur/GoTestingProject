package iteration

//const repeatCount = 5

func Repeat(character string, count int) string {
	//var repeated string
	repeated := ""
	for i := 0; i < count; i++ {
		repeated += character
		//keluarkan value dlu
		//fmt.Printf("i = %d, repeated: %q\n", i, repeated)
		//i++ run last skali
	}
	return repeated
}
