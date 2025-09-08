package main

func createMatrix(rows, cols int) [][]int {
	mat := make([][]int, rows)
	for i := range rows {
		mat[i] = make([]int, cols)
		for j := range cols {
			mat[i][j] = i * j
		}
	}
	return mat
}
