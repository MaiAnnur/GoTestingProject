package iteration

import "fmt"

func ExampleRepeat() {
	repeated := Repeat("b", 3)
	fmt.Println(repeated)
	// Output: bbb
}
