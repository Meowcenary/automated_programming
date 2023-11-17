//go:generate go run stats_test_generator.go

package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/montanaflynn/stats"
)

// Coordinate represents a pair of x and y values
type Coordinate struct {
	X float64
	Y float64
}

func executeTemplate(tmpl *template.Template, data interface{}) (string, error) {
	var result bytes.Buffer
	err := tmpl.Execute(&result, data)
	if err != nil {
		return "", err
	}
	return result.String(), nil
}

func createTemplate(templateName, templateStr string) (*template.Template, error) {
	template, err := template.New(templateName).Parse(templateStr)
	if err != nil {
		return nil, err
	}
	return template, nil
}

func writeToFile(fileName, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func headTemplateStr() string {
	return `package stats_test

import (
	"fmt"
	"math"
	"slices"
	"testing"

	"github.com/montanaflynn/stats"
)
`
}

func anscombeFunctionsTemplateStr() string {
return `{{range $functionName, $dataset := .}}
func {{ $functionName }}() []stats.Coordinate {
	return []stats.Coordinate{
		{{range $index, $element := $dataset}}{ {{printf "%.16f" $element.X}}, {{printf "%.16f" $element.Y}}},{{end}}
	}
}
{{end}}
`
}

// There isn't anything to templatize at this point, but this could be refactored if there is a need to generate
// Similar functions to this
func findGradientAndInterceptTemplate() string {
	return `func FindGradientAndIntercept(s []stats.Coordinate) (gradientIntercept []float64, err error) {
	if len(s) == 0 {
		// Cannot return nil for type float64
		return []float64{0.0, 0.0}, stats.EmptyInputErr
	}

	// Placeholder for the math to be done
	var sum [5]float64

	// Loop over data keeping index in place
	i := 0
	for ; i < len(s); i++ {
		sum[0] += s[i].X
		sum[1] += s[i].Y
		sum[2] += s[i].X * s[i].X
		sum[3] += s[i].X * s[i].Y
		sum[4] += s[i].Y * s[i].Y
	}

	// Find gradient and intercept
	f := float64(i)
	gradient := (f*sum[3] - sum[0]*sum[1]) / (f*sum[2] - sum[0]*sum[0])
	intercept := (sum[1] / f) - (gradient * sum[0] / f)

	return []float64{gradient, intercept}, nil
}

`
}

func roundFloatTemplate() string {
	return `func roundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}

`
}

func testAnscombeLinearRegressionGradientAndInterceptsTemplate() string {
	return `func TestAnscombeLinearRegressionGradientsAndIntercepts(t *testing.T) {
    type test struct {
        input []stats.Coordinate
        want  []float64
    }

    /*
    input is one of the Anscombe data fixtures, want is a slice of floats where the first value is the gradient and the
    second value is the intercept
    */
    tests := []test{
        {input: Anscombe1(), want: []float64{0.500091, 3.000091}},
        {input: Anscombe2(), want: []float64{0.500000, 3.000909}},
        {input: Anscombe3(), want: []float64{0.499727, 3.002455}},
        {input: Anscombe4(), want: []float64{0.499909, 3.001727}},
    }

    for i, test := range tests {
        got, _ := FindGradientAndIntercept(test.input)

        if (roundFloat(got[0], 6) != roundFloat(test.want[0], 6)) || (roundFloat(got[1], 6) != roundFloat(test.want[1], 6)) {
            t.Errorf("Linear regression gradient and intercept for Anscombe %d do not match expectations\nWant:\nGradient: %f\nIntercept: %f\nGot:\nGradient: %f\nIntercept: %f",
                     i+1,
                     roundFloat(test.want[0], 6),
                     roundFloat(test.want[1], 6),
                     roundFloat(got[0], 6),
                     roundFloat(got[1], 6))
        }
    }
}

`
}

func testAnscombeLinearRegressionCoefficientsTemplate() string {
	return `func TestAnscombeLinearRegressionCoefficients(t *testing.T) {
    type test struct {
        input []stats.Coordinate
        want  []stats.Coordinate
    }

    tests := []test{
        {input: Anscombe1(), want: Anscombe1Coefficients()},
        {input: Anscombe2(), want: Anscombe2Coefficients()},
        {input: Anscombe3(), want: Anscombe3Coefficients()},
        {input: Anscombe4(), want: Anscombe4Coefficients()},
    }

    for i, test := range tests {
        got, _ := stats.LinearRegression(test.input)

        if !slices.Equal(got, test.want) {
            t.Errorf("Linear regression coefficients for Anscombe %d do not match expectations\n%s\n%s",
                     i+1,
                     fmt.Sprint(got),
                     fmt.Sprint(test.want))
        }
    }
}

`
}

func benchmarkTemplate() string {
	return `var blackhole []stats.Coordinate
func BenchmarkAnscombeLinearRegression(b *testing.B) {
		result, _ := stats.LinearRegression(Anscombe1())
    result, _ = stats.LinearRegression(Anscombe2())
    result, _ = stats.LinearRegression(Anscombe3())
    result, _ = stats.LinearRegression(Anscombe4())
    blackhole = result
}

`
}

func main() {
	// Define your datasets and coefficients
	datasets := map[string][]Coordinate{
		"Anscombe1": {
			{10, 8.04},
			{8, 6.95},
			{13, 7.58},
			{9, 8.81},
			{11, 8.33},
			{14, 9.96},
			{6, 7.24},
			{4, 4.26},
			{12, 10.84},
			{7, 4.82},
			{5, 5.68},
		},
		"Anscombe2": {
			{10, 9.14},
			{8, 8.14},
			{13, 8.74},
			{9, 8.77},
			{11, 9.26},
			{14, 8.1},
			{6, 6.13},
			{4, 3.1},
			{12, 9.13},
			{7, 7.26},
			{5, 4.74},
		},
		"Anscombe3": {
			{10, 7.46},
			{8, 6.77},
			{13, 12.74},
			{9, 7.11},
			{11, 7.81},
			{14, 8.84},
			{6, 6.08},
			{4, 5.39},
			{12, 8.15},
			{7, 6.42},
			{5, 5.73},
		},
		"Anscombe4": {
			{8, 6.58},
			{8, 5.76},
			{8, 7.71},
			{8, 8.84},
			{8, 8.47},
			{8, 7.04},
			{8, 5.25},
			{19, 12.5},
			{8, 5.56},
			{8, 7.91},
			{8, 6.89},
		},
	}

	// the string argument to template.new is the name of the generated template
	// https://pkg.go.dev/text/template
	anscombefunctionstemplate, err := template.New("datasets").Parse(anscombeFunctionsTemplateStr())

	if err != nil {
		panic(err)
	}

	// execute the anscombefunctionstemplate into a string
	anscombeFunctionsResult, err := executeTemplate(anscombefunctionstemplate, datasets)

	coefficients := map[string]stats.Series{
		"Anscombe1Coefficients": {
			{10, 8.001000000000001},
			{8, 7.000818181818185},
			{13, 9.501272727272724},
			{9, 7.500909090909093},
			{11, 8.501090909090909},
			{14, 10.001363636363633},
			{6, 6.000636363636369},
			{4, 5.000454545454553},
			{12, 9.001181818181816},
			{7, 6.500727272727277},
			{5, 5.5005454545454615},
		},
		"Anscombe2Coefficients": {
			{10, 8.00090909090909},
			{8, 7.000909090909092},
			{13, 9.500909090909088},
			{9, 7.500909090909091},
			{11, 8.50090909090909},
			{14, 10.000909090909087},
			{6, 6.000909090909094},
			{4, 5.000909090909095},
			{12, 9.00090909090909},
			{7, 6.500909090909093},
			{5, 5.500909090909094},
    },
		"Anscombe3Coefficients": {
			{10, 7.999727272727271},
			{8, 7.000272727272732},
			{13, 9.498909090909081},
			{9, 7.500000000000002},
			{11, 8.499454545454540},
			{14, 9.998636363636351},
			{6, 6.000818181818192},
			{4, 5.001363636363653},
			{12, 8.999181818181810},
			{7, 6.5005454545454615},
			{5, 5.501090909090922},
		},
		"Anscombe4Coefficients": {
			{8, 7.000999999999998},
			{8, 7.000999999999998},
			{8, 7.000999999999998},
			{8, 7.000999999999998},
			{8, 7.000999999999998},
			{8, 7.000999999999998},
			{8, 7.000999999999998},
			{19, 12.500000000000018},
			{8, 7.000999999999998},
			{8, 7.000999999999998},
			{8, 7.000999999999998},
		},
	}

	// the string argument to template.new is the name of the generated template
	// https://pkg.go.dev/text/template
	coefficientsTemplate, err := template.New("coefficients").Parse(anscombeFunctionsTemplateStr())
	if err != nil {
		panic(err)
	}
	// execute the  into a string
	coefficientsResult, err := executeTemplate(coefficientsTemplate, coefficients)
	if err != nil {
		panic(err)
	}

	// Combine the strings to form the completed Go file
	combinedResult := headTemplateStr() + anscombeFunctionsResult + coefficientsResult + findGradientAndInterceptTemplate() +
		roundFloatTemplate() + testAnscombeLinearRegressionGradientAndInterceptsTemplate() + testAnscombeLinearRegressionCoefficientsTemplate() +
		benchmarkTemplate()

	// Write the combined result to a file
	fileName := "stats_test.go"
	err = writeToFile(fileName, combinedResult)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
	fmt.Println("Combined template result written to:", fileName)
}
