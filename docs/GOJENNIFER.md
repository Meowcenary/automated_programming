## Go Jennifer
[Back to README](../README.md)

Jennifer is a third party package that defines a domain specific language (DSL) for metaprogramming in Go. Personally I
struggled to produce anything particularly noteworthy using it, but I included the code examples undre the `jennifer/`
for the sake of comparison to the specific problem of generating Anscombe's Quartet and running validation tests against
it. The example code: `stats_test_gen.go` will output a function that returns the first Anscombe data set, but that was
all that could be completed with the time allotted.

### Dependencies
To run the generator you will need to install the package jennifer: `github.com/dave/jennifer` with this command:
`go get -u github.com/dave/jennifer/jen`

### Running the Code
Change directories to `jennifer/` and then run `go run stats_test_gen.go` which will output the generated Go code. From
a unix shell This can be chained with the `>` operator to redirect the output to a file E.g:
`go run stats_test_gen.go > stats_test.go`
