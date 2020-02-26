package lib

import (
	"math"
	"testing"
	"gonum.org/v1/gonum/mat"
)

func TestVandermonde(t *testing.T) {
	if _, err := Vandermonde([]float64{}, 0, 0); err == nil {
		t.Errorf("gave an empty matrix, got no error, but should have")
	}

	two := []float64{2}

	for axis := 0; axis <= 1; axis++ {
		one_half, _ := Vandermonde(two, -1, axis)
		one     , _ := Vandermonde(two,  0, axis)
		also_two, _ := Vandermonde(two,  1, axis)
		four    , _ := Vandermonde(two,  2, axis)

		if one_half.At(0, 0) != 0.5 { t.Errorf("vand([2], -1, 0) should be [1/2]") }
		if one     .At(0, 0) !=   1 { t.Errorf("vand([2],  0, 0) should be [  1]") }
		if also_two.At(0, 0) !=   2 { t.Errorf("vand([2],  1, 0) should be [  2]") }
		if four    .At(0, 0) !=   4 { t.Errorf("vand([2],  2, 0) should be [  4]") }
	}

	_123 := []float64{1,2,3}
	vandermonde_of_123 := mat.NewDense(3, 3, []float64{
		1, 1, 1,
		1, 2, 4,
		1, 3, 9,
	})

	v123, _ := Vandermonde(_123, 0, 0)
	matrix_equality("v123", t, vandermonde_of_123, v123)

	v123_t, _ := Vandermonde(_123, 0, 1)
	matrix_equality("v123_t", t, vandermonde_of_123.T(), v123_t)

	_537 := []float64{5,3,7}
	high_powers := mat.NewDense(3, 3, []float64{
		25, 125, 625,
		9, 27, 81,
		49, 343, 2401,
	})

	v537, _ := Vandermonde(_537, 2, 0)
	matrix_equality("v123", t, high_powers, v537)

	v537_t, _ := Vandermonde(_537, 2, 1)
	matrix_equality("v123_t", t, high_powers.T(), v537_t)

	floats := []float64{ 0.6, 3.14, 6.9 }
	random := mat.NewDense(3, 3, []float64{
		2.777778,   1.66667, 1,
		0.101424,  0.318471, 1,
		0.021004,  0.144928, 1,
	})

	v_rand, _ := Vandermonde(floats, -2, 0)
	matrix_equality("floats", t, random, v_rand)

	v_rand_t, _ := Vandermonde(floats, -2, 1)
	matrix_equality("floats_t", t, random.T(), v_rand_t)
}

func matrix_equality(msg string, t *testing.T, expected, actual mat.Matrix) {
	row_expected, col_expected := expected.Dims()
	row_actual, col_actual := actual.Dims()
	if row_actual != row_expected || col_expected != col_actual {
		t.Errorf("%v: dimension mismatch: expected (%v, %v), got (%v, %v)", msg, row_expected, col_expected, row_actual, col_actual)
		return
	}

	for i := 0; i < row_actual; i++ {
		for j := 0; j < col_actual; j++ {
			actual_val := actual.At(i, j)
			expected_val := expected.At(i, j)
			if math.Abs(expected_val - actual_val) > 0.00001 {
				t.Errorf("%v: val mismatch: expected %v, got %v", msg, expected_val, actual_val)
			}
		}
	}
}
