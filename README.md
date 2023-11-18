# Automated Programming
Several methods of automated code generation were explored with the goal of producing a file of Go code that could run
a linear regression test against the data sets in Anscombe's Quartet with the results written out into the sections
below. Ultimately from the experiments done here it seems that AI generated code can certainly assist programmers and
lead to improved productivity, but in it's current state it can not come close to replacing an experienced developer.
Consistently the AI tools would return code that could not be compiled at all, returned erroneous data, or would return
data formatted in ways that were difficult to read and to work with in other functions. In it's current state AI is not
able to replace programmers, but it can help someone with strong work ethic and some talent increase their productivity
signifantly.

The original file that was used for the assignment "Testing Go with Statistics" is included with the project with the
name: `original_stats_test.go`

### Cloning the Repo
Clone  with git
- HTTP: `git clone https://github.com/Meowcenary/automated_programming.git`
- SSH: `git clone git@github.com:Meowcenary/automated_programming.git`

### Metaprogramming
Go's generate tool was used to implement a simple metaprogramming example that can be viewed at:
`go_generate/stats_test_generator.go`. The package "Jennifer" was also explored, but given time constraints this was
left in an incomplete state with very limited functionality. In this particular instance Go's generate tool was much
easier to work with because it drew upon Go's standard library whereas Jennifer required learning a specific syntax that
I was less experienced in. There's potential with Jennifer, but the cost of adding an additonal dependency and learning
to properly use it outweighs it's usefulness in my opinion.
- [Go generate](docs/GOGENERATE.md)
- [Go Jennifer](docs/GOJENNIFER.md)

### AI Assisted Code Generation
The middleground between metaprogramming and fully AI generated code is AI assisted coding. This can be thought of as
pair programming, but with one member of the pair being an AI. The first AI assistant that was explored was Github's
"Copilot" which is one of the currenty most popular AI assistants. Copilot is not free and requires a subscription of
$10/month or $100/year for individuals and $19/month for every user for businesses. The other AI assistant selected was
"Blackbox" which is free and conveniently available as a plugin for Visual Studio Code. Initially "Blackbox" was used
for code generation, but after having some challenges a second attempt was made to just refactor the initial project.
- [Github Copilot](docs/COPILOT.md)
- [Blackbox](docs/BLACKBOX.md)
- [Blackbox Refactor](docs/BLACKBOXREFACTOR.md)

### AI Generated Code
ChatGPT was used to attempt to generate code that ran the same tests as required for the assignment "Testing Go for
Statistics". The chat logs are written out in a file linked to below. Within the chat logs the prompts are written out
with a leading "->" while the responses are written out below following "Answer:". Large blocks of irrelevant response
text are omitted to save space.
- [ChatGPT Log](docs/CHATGPT.md)

### References
- [Go jennifer package](https://github.com/dave/jennifer)
- [Go generate package docs](https://pkg.go.dev/cmd/go#hdr-Generate_Go_files_by_processing_source)
- [Go convert templates to strings](https://coderwall.com/p/ns60fq/simply-output-go-html-template-execution-to-strings)
- [Go blog on Go generate](https://go.dev/blog/generate)
- [Go generate proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/go-generate.md)
- [Generating code with go generate - Daniel Milde](https://www.youtube.com/watch?v=ClW_g1iDGi4)
    - [Code from video](https://github.com/dundee/gogenerate)
- [About Githug Copilot](https://docs.github.com/en/copilot/overview-of-github-copilot/about-github-copilot-individual)
- [Visual Studio Blackbox](https://marketplace.visualstudio.com/items?itemName=Blackboxapp.blackbox)
- [Blackbox website](https://www.useblackbox.io/)
- [32 ChatGPT Tips for Beginners in 2023! (Become a PRO!)](https://www.youtube.com/watch?v=dUjWMdR_Kw8)
