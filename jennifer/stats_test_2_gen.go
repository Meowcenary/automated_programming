package main

import (
	"fmt"
	"os"

	"github.com/dave/jennifer/jen"
	"github.com/montanaflynn/stats"
)

func generateCoordinateArrayReturnFunction(arrayParam []stats.Coordinate) *jen.Statement {
	return jen.Func().
		Id("returnCoordinateArray").
		Params().
		Index().Qual("github.com/montanaflynn/stats", "Coordinate").
		Block(
			jen.Return(
				jen.Id("arrayParam"),
			),
		)
}

func main() {
	f := jen.NewFile("main")

	// Sample array of Coordinate passed as an argument
	coordinateParam := []stats.Coordinate{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	f.Add(generateCoordinateArrayReturnFunction(coordinateParam))

	err := f.Render(os.Stdout)
	if err != nil {
		fmt.Println("Error rendering code:", err)
		return
	}
}
