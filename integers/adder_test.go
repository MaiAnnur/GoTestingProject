package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestSubtract(t *testing.T) {
	difference := Subtract(5, 3)
	expected := 2

	if difference != expected {
		t.Errorf("expected '%d' but got '%d'", expected, difference)
	}
}

func TestMultiply(t *testing.T) {
	product := Multiply(2, 3)
	expected := 6

	if product != expected {
		t.Errorf("expected '%d' but got '%d'", expected, product)
	}
}

func TestDivide(t *testing.T) {
	quotient := Divide(6, 3)
	expected := 2

	if quotient != expected {
		t.Errorf("expected '%d' but got '%d'", expected, quotient)
	}
}

func TestDivideByZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic but did not get one")
		}
	}()
	Divide(6, 0)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)

}
