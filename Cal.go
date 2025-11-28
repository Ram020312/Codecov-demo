package cal
import (
	"errors"
)
func add(a, b int) (int, error) {
	return a + b, nil
}

func sub(a, b int) (int, error) {
	return a - b, nil 
}

func mul(a, b int) (int, error){
	return a * b, nil 
}

func Div(a, b int) (int, error) {
	if (b == 0) {
		return -1, errors.New("error: denominator can't be zero")
	}
	return a / b, nil
}