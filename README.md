# Vandermonde

### What's a Vandermonde matrix?

A [Vandermonde matrix](https://proofwiki.org/wiki/Definition:Vandermonde_Matrix) is a square matrix that satisfies

```
a_ij = x_i ^ (j − 1)
a_ij = x_j ^ i
a_ij = x_i ^ (n − j)
```

Usually what you want is the first form, i.e.

```
[1 x x^2 x^3 ... x ^ (n - 1)]
[1 y y^2 y^3 ... y ^ (n - 1)]
...
[1 s s^2 s^3 ... s ^ (n - 1)]
```

### Importing
```
import "github.com/AnthonyHewins/vandermonde" 
```

### Usage

Generally, what you want is this:

Standard Vandermonde:
```go
my_data := []float64{1,2,3}
matrix, err := vandermonde.Vandermonde(my_data, 0, 0)

// matrix = [1 1 1]
//          [1 2 4]
//          [1 3 9]
```

Transposed version:
```go
my_data := []float64{1,2,3}
matrix, err := vandermonde.Vandermonde(my_data, 0, 1) // axis=1, similar to pandas and other data science langs

// matrix = [1 1 1]
//          [1 2 3]
//          [1 4 9]

// Note: the returned data is a gonum *mat.Dense, so you can also
// do matrix.T() if you need both the matrix and its transpose.
```

Starting with a different exponent:
```go
my_data := []float64{1,2,3}
matrix, err := vandermonde.Vandermonde(my_data, 2, 0)
// matrix = [1  1  1]
//          [4  8 16]
//          [9 27 81]

my_data := []float64{1,2,3}
matrix, err := vandermonde.Vandermonde(my_data, -1, 0)
// matrix = [1   1 1]
//          [1/2 1 2]
//          [1/3 1 3]
```

For more advanced usage when you only a slice of the Vandermonde matrix, which
can be larger or smaller:

``` go
my_data := []float64{1,2,3}

// start with exponent 2, generate 4 terms (instead of 3), default axis
matrix, err := vandermonde.VandermondeWindow(my_data, 2, 4, 0)
// matrix = [1  1  1   1]
//          [4  8 16  32]
//          [9 27 81 243]


// start with exponent -1, generate only 1 row, transposed axis
matrix, err := vandermonde.VandermondeWindow(my_data, -1, 1, 0)
// matrix = [1 1/2 1/3]
```
