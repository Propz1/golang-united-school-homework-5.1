package square

import (
	"testing"
)

var (
	p Point = Point{x: 10, y: 7}

	input = Square{
		start: p,
		a:     5,
	}
)

func TestSquare(t *testing.T) {

	t.Run("End", func(t *testing.T) {

		var result = Point{x: 5, y: 12}

		realResult := input.End()

		if realResult != result {
			t.Errorf("Expected result %v != %v real result", result, realResult)
		}

	})

	t.Run("Area", func(t *testing.T) {

		var result uint = 25

		realResult := input.Area()

		if realResult != result {
			t.Errorf("Expected result %v != %v real result", result, realResult)
		}

	})

	t.Run("Perimeter", func(t *testing.T) {

		var result uint = 20

		realResult := input.Perimeter()

		if realResult != result {
			t.Errorf("Expected result %v != %v real result", result, realResult)
		}

	})
}
