package cal

import ( 
	"testing"
	"github.com/stretchr/testify/assert"
	)

func TestAdd(t *testing.T) {
    got, _ := add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add() = %d; want %d", got, want)
    }
}

func TestSub(t *testing.T) {
    got, _ := sub(5, 2)
    want := 3
    if got != want {
        t.Errorf("Sub() = %d; want %d", got, want)
    }
}

func TestMul(t *testing.T) {
    got, _ := mul(4, 3)
    want := 12
    if got != want {
        t.Errorf("Mul() = %d; want %d", got, want)
    }
}

func TestDiv(t *testing.T) {
    // normal case
    got, _ := Div(10, 2)
    want := 5
    if got != want {
        t.Errorf("Div() = %d; want %d", got, want)
    }

    // division by zero
    _, err := Div(10, 0)
    if err == nil {
        t.Error("Div() expected error for denominator 0, got nil")
    }
}
