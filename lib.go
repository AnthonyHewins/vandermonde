package vandermonde

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

// Generate a standard Vandermonde matrix, returning a gonum Dense matrix.
// Go here for more info: https://proofwiki.org/wiki/Definition:Vandermonde_Matrix
func Vandermonde(x []float64, start_deg, axis int) (*mat.Dense, error) {
	n := len(x)
	if n == 0 { return nil, fmt.Errorf("x must have contents") }

	//buf := mat.NewDense(n, n, nil)
	X := mat.NewDense(n, n, nil)
	if axis != 1 {
		X.Apply(func(row, col int, v float64) (float64) {
			// Caller desires (assuming start_deg is 0):
			//
			// [ 1 x1 x1^2 x1^3 ... ]
			// [ 1 x2 x2^2 x2^3 ... ] which is x[row] ^ (column + start_deg)
			// [ 1 x3 x3^2 x3^3 ... ]
			return math.Pow(x[row], float64(col + start_deg))
		}, X)
	} else {
		X.Apply(func(row, col int, v float64) (float64) {
			// Caller desires (assuming start_deg is 0):
			//
			// [  1     1     1     1   ... ]
			// [ x1    x2    x3    x4   ... ] which is x[column] ^ (column + start_deg)
			// [ x1^2  x2^2  x3^2  x4^3 ... ]
			return math.Pow(x[col], float64(row + start_deg))
		}, X)
	}

	return X, nil
}
