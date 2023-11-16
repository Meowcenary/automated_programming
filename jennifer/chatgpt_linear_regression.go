package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
)

func anscombeQuartet() ([][]float64, [][]float64) {
// Dataset I
	x1 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y1 := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	// Dataset II
	x2 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y2 := []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.10, 6.13, 3.10, 9.13, 7.26, 4.74}

	// Dataset III
	x3 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y3 := []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}

	// Dataset IV
	x4 := []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}
	y4 := []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.50, 5.56, 7.91, 6.89}

	return [][]float64{x1, x2, x3, x4}, [][]float64{y1, y2, y3, y4}
}

func linearRegression(coordinates []stats.Coordinate) (float64, float64, error) {
	// Extract x and y values from the coordinates
	x := make([]float64, len(coordinates))
	y := make([]float64, len(coordinates))
	for i, coord := range coordinates {
		x[i] = coord.X
		y[i] = coord.Y
	}

	// Perform linear regression
	slope, intercept, err := stats.LinearRegression(x, y)
	if err != nil {
		return 0, 0, err
	}
	return slope, intercept, nil
}

func main() {
	x, y := anscombeQuartet()

	for i := 0; i < len(x); i++ {
		// Combine x and y values to create stats.Coordinate objects
		coordinates := make([]stats.Coordinate, len(x[i]))
		for j := 0; j < len(x[i]); j++ {
			coordinates[j] = stats.Coordinate{X: x[i][j], Y: y[i][j]}
		}

		// Perform linear regression for each dataset
		slope, intercept, err := linearRegression(coordinates)
		if err != nil {
			fmt.Printf("Error performing linear regression for dataset %d: %v\n", i+1, err)
			continue
		}

		fmt.Printf("Dataset %d Linear Regression:\n", i+1)
		fmt.Printf("Slope: %f\n", slope)
		fmt.Printf("Intercept: %f\n", intercept)
		fmt.Println()
	}
}

