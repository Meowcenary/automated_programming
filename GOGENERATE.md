## Go Generate
[Back to README](README.md)

Generate is a tool built into Go that scans files for directives in the form of comments that give instructions for what
generate should do. The instruction set itself is fairly limited and a common approach to using generate is to write
standard Go code that uses string formatting or the template library to generate a different Go file and then mark the
file as a generate command using: `the command`

Generate is particularly useful for writing boilerplate code, but can also be used to generate code based on
configurations or other data. In this instance, recreating the tests for the Anscombe Quartet, it did not prove very
helpful. The functionality required was relatively simple and as a result there was not much benefit from
templatization. More generally generate could be useful as a means assist in writing large blocks of templatized code,
things like constructors or writing many version of a function for different types as an alternative to generics, but it
is challenging enough to work with that it would take an experienced developer to build anything with it and any
resulting tool would be unlikely to truly replace a developer in terms of code output and value added to a project.
