### Cloning the Repo
Clone  with git
- HTTP: `git clone https://github.com/Meowcenary/automated_programming.git`
- SSH: `git clone git@github.com:Meowcenary/automated_programming.git`

### Metaprogramming
Go's generate tool was used to implement a simple metaprogramming example that can be viewed at:
`go_generate/stats_test_generator.go`. The package "Jennifer" was also explored, but given time constraints this was
left in an incomplete state.

- [Go generate](docs/GOGENERATE.md)
- [Go Jennifer](https://github.com/dave/jennifer)

### AI Assisted Code Generation
TODO: Pick a code assistance tool to use

### AI Generated Code
Both ChatGPT and Google Bard were used to attempt to generate code that ran the same tests as required for the
assignment "Testing Go for Statistics". The chat logs are linked below. Within the chat logs the prompts are written out
following "->" with the response below following "answer:". Large blocks of irrelevant response text are omitted to save
space.

- [ChatGPT Log](docs/CHATGPT.md)
- [Google Bard Log](#)

### References
- Go jennifer - https://github.com/dave/jennifer
- Go generate package docs - https://pkg.go.dev/cmd/go#hdr-Generate_Go_files_by_processing_source
- Go convert templates to strings - https://coderwall.com/p/ns60fq/simply-output-go-html-template-execution-to-strings
- Go blog on Go generate - https://go.dev/blog/generate
- Go generate proposal - https://go.googlesource.com/proposal/+/refs/heads/master/design/go-generate.md
- Generating code with go generate - Daniel Milde - https://www.youtube.com/watch?v=ClW_g1iDGi4
    - Code from video: https://github.com/dundee/gogenerate
- 32 ChatGPT Tips for Beginners in 2023! (Become a PRO!) - https://www.youtube.com/watch?v=dUjWMdR_Kw8
