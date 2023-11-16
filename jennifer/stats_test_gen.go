package main

import (
	"fmt"
	"os"

	"github.com/dave/jennifer/jen"
)

func generateAnscombe1Function() *jen.Statement {
	return jen.Func().
		Id("Anscombe1").
		Params().
		Index().Qual("github.com/montanaflynn/stats", "Coordinate").
		Block(
			jen.Return(
				jen.Index().Qual("github.com/montanaflynn/stats", "Coordinate").Values(
					jen.Lit(10), jen.Lit(8.04),
				),
				jen.Index().Qual("github.com/montanaflynn/stats", "Coordinate").Values(
					jen.Lit(8), jen.Lit(6.95),
				),
				jen.Index().Qual("github.com/montanaflynn/stats", "Coordinate").Values(
					jen.Lit(13), jen.Lit(7.58),
				),
				jen.Index().Qual("github.com/montanaflynn/stats", "Coordinate").Values(
					jen.Lit(9), jen.Lit(8.81),
				),
				jen.Index().Qual("github.com/montanaflynn/stats", "Coordinate").Values(
					jen.Lit(11), jen.Lit(8.33),
				),
				jen.Index().Qual("github.com/montanaflynn/stats", "Coordinate").Values(
					jen.Lit(14), jen.Lit(9.96),
				),
				jen.Index().Qual("github.com/montanaflynn/stats", "Coordinate").Values(
					jen.Lit(6), jen.Lit(7.24),
				),
				jen.Index().Qual("github.com/montanaflynn/stats", "Coordinate").Values(
					jen.Lit(4), jen.Lit(4.26),
				),
				jen.Index().Qual("github.com/montanaflynn/stats", "Coordinate").Values(
					jen.Lit(12), jen.Lit(10.84),
				),
				jen.Index().Qual("github.com/montanaflynn/stats", "Coordinate").Values(
					jen.Lit(7), jen.Lit(4.82),
				),
				jen.Index().Qual("github.com/montanaflynn/stats", "Coordinate").Values(
					jen.Lit(5), jen.Lit(5.68),
				),
			),
		)
}

func main() {
	f := jen.NewFile("main")
	f.Add(generateAnscombe1Function())

	err := f.Render(os.Stdout)
	if err != nil {
		fmt.Println("Error rendering code:", err)
		return
	}
}
