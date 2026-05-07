package main

func rotateTheBox(boxGrid [][]byte) [][]byte {
	m := len(boxGrid)
	n := len(boxGrid[0])
	// Step 1: Apply gravity in original orientation (stones fall to the right)
	for r := 0; r < m; r++ {
		writePos := n - 1
		for c := n - 1; c >= 0; c-- {
			if boxGrid[r][c] == '#' {
				// Move stone to writePos
				boxGrid[r][c] = '.'
				boxGrid[r][writePos] = '#'
				writePos--
			} else if boxGrid[r][c] == '*' {
				// Obstacle found, reset writePos to position before obstacle
				writePos = c - 1
			}
		}
	}
	// Step 2: Rotate 90° clockwise
	// Original (r, c) -> rotated (c, m-1-r)
	rotated := make([][]byte, n)
	for i := 0; i < n; i++ {
		rotated[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			rotated[i][j] = '.'
		}
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			rotated[c][m-1-r] = boxGrid[r][c]
		}
	}

	return rotated
}
