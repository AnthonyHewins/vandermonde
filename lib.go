package vandermonde

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

// Generate a standard Vandermonde matrix, returning a gonum Dense matrix.
// Go here for more info: https://proofwiki.org/wiki/Definition:Vandermonde_Matrix
//
// This is a special case of VandermondeWindow.
func Vandermonde(x []float64, start_deg, axis int) (*mat.Dense, error) {
	return vandermonde_core(x, start_deg, len(x), axis)
}

// Generate a Vandermonde matrix, stopping at a certain polynomial degree. This allows it to be rectangular instead
// of square. You can essentially capture any window of the Vandermonde matrix for any vector x.
//
// The last argument, axis, transposes the result if axis == 1 (the transpose imposes no runtime cost)
func VandermondeWindow(x []float64, start_deg, cutoff, axis int) (*mat.Dense, error) {
	if cutoff < 1 { return nil, fmt.Errorf("cutoff must be > 1; empty matrices not allowed") }
	return vandermonde_core(x, start_deg, cutoff, axis)
}

func vandermonde_core(x []float64, start_deg, cutoff, axis int) (*mat.Dense, error) {
	n := len(x)
	if n == 0 { return nil, fmt.Errorf("x must have contents")}

	if axis != 1 {
		X := mat.NewDense(n, cutoff, nil)

		X.Apply(func(row, col int, v float64) (float64) {
			// Caller desires (assuming start_deg is 0):
			//
			// [ 1 x1 x1^2 x1^3 ... ]
			// [ 1 x2 x2^2 x2^3 ... ] which is x[row] ^ (column + start_deg)
			// [ 1 x3 x3^2 x3^3 ... ]
			return math.Pow(x[row], float64(col + start_deg))
		}, X)

		return X, nil
	} else {
		X := mat.NewDense(cutoff, n, nil)

		X.Apply(func(row, col int, v float64) (float64) {
			// Caller desires (assuming start_deg is 0):
			//
			// [  1     1     1     1   ... ]
			// [ x1    x2    x3    x4   ... ] which is x[column] ^ (column + start_deg)
			// [ x1^2  x2^2  x3^2  x4^3 ... ]
			return math.Pow(x[col], float64(row + start_deg))
		}, X)

		return X, nil
	}
}
