package array

func Apply(op func(a, b float64) float64, a, b []float64) []float64 {
	assertSize(a, b)

	size := len(a)
	arr := make([]float64, size)
	for i := 0; i < size; i++ {
		arr[i] = op(a[i], b[i])
	}

	return arr
}

func Product(a, b []float64) []float64 {
	return Apply(func(a, b float64) float64 { return a * b }, a, b)
}

func Addition(a, b []float64) []float64 {
	return Apply(func(a, b float64) float64 { return a + b }, a, b)
}

func Substitute(a, b []float64) []float64 {
	return Apply(func(a, b float64) float64 { return a - b }, a, b)
}

func Divide(a, b []float64) []float64 {
	return Apply(func(a, b float64) float64 { return a / b }, a, b)
}

func Dot(a, b []float64) float64 {
	result := Product(a, b)

	scalar := 0.0
	for _, val := range result {
		scalar += val
	}

	return scalar
}

func MatrixVectorProduct(vec []float64, matrix [][]float64) []float64 {
	if len(vec) != len(matrix) {
		panic("size mismatch")
	}

	output := make([]float64, len(vec))

	matrixSize := len(matrix)

	for i := 0; i < matrixSize; i++ {
		output[i] = Dot(vec, matrix[i])
	}

	return output
}

func MatrixTranspose(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]float64, cols)

	for i := 0; i < cols; i++ {
		transposed[i] = make([]float64, rows)
		for k := 0; k < rows; k++ {
			transposed[i][k] = matrix[k][i]
		}
	}

	return transposed
}

func assertSize(a, b []float64) {
	if len(a) != len(b) {
		panic("size mismatch")
	}
}
