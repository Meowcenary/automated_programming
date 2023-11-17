## Go Generate
[Back to README](../README.md)

Generate is a tool built into Go that scans files for directives in the form of comments that give instructions for what
generate should do. The instruction set itself is fairly limited and a common approach to using generate is to write
standard Go code that uses string formatting or the template library to generate a different Go file and then mark the
file as a generate command using something like: `//go:generate go run <generator file name>.go`

Generate is particularly useful for "metaprogramming", writing programs that can edit other programs and themself. A
common use case for metaprogramming is writing boilerplate code, but can also be used to generate code based on
configurations or other data. More generally generate could be useful as a means assist in writing large blocks of
templatized code, things like constructors or writing many version of a function for different types as an alternative
to generics. In this instance, recreating the tests for the Anscombe Quartet, it did not prove very helpful, but it did
elucidate the potential generate has.

From this brief example it seems that the major trade off of using generate is that adds significant complexity in an
effort to reduce verbosity. While this example was very simple, it still raised some frustrating bugs that were the
result of templates behaving in unexpected ways with particular problems keeping track of floating point precision.

In summary, Go's generate command has the potential to increase a programmer's productivity, but require's an
experienced user to know where it will be effective.

### Dependencies
To run the generator you will need to install this stats package: `github.com/montanaflynn/stats` with this command:
`go get github.com/montanaflynn/stats`

### Running the Generator
Change directories to `go_generate/` and then run `go generate stats_test_generator`. If successful a message will print
with the file that the generated code was written to `stats_test.go`. The generated code can then be run with
`go test stats_test.go`.

### Creating Go Generator Code With AI

**WIP - might be better to leave this out and focus on other aspects of this project**

After working through a few tutorial examples and reading up on Go's generate tool I was still having some challenges
wrapping my head around how to properly use it. To help understand how to use generate I decided to ask Google's Bard AI
for some help. The below log marks the prompt/question with a leading "->" and the answer with "Answer: "

**-> Can you help me write a go template string that accepts a struct as an argument?**

Google Bard was helpful in getting started, but I was struggling to combine the text templates together in a way that
made sense to me.
