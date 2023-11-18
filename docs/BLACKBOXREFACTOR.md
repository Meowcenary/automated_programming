## Blackbox Refactor
[Back to README](../README.md)

After the initial attempt to generate code with Blackbox it became clear that Blackbox was not as strong at replying to
prompts as ChatGPT or Google Bard. One interesting advantage Blackbox had over these tools was a plugin to integrate
into Visual Studio Code. The original Go code was copied and loaded into VS Code and the Blackbox plugin was installed
alongside it. The plugin itself uses a similar prompt and reply model to other conversational AIs with the chat
integrated directly into the IDE.

Note: For some reason Blackbox was referring to a file I had never opened or edited which is why it keeps mentioning
Python
**-> Can you help me refactor the code in this file?**
Answer: Too many print statements can be consolidated using string concatenation or f-strings.
For example, in your create_character() function, you can consolidate all the print statements into one:

```python
def create_character():
    ...
    print("Welcome to the world of Legends, your chosen race is:", race,
          "and your chosen class is:", char_class,
          "Please note that all abilities, strengths, and weaknesses of each class are subject to change as the game evolves.")
```
Repeated print statements can be refactored into a separate function.
For example, you have several instances of print("What would you like to do?") in your start() function. You can create a separate function to handle this:

```python
def prompt_user():
    print("What would you like to do?")

def start():
    ...
    while True:
        prompt_user()
        choice = input().lower()
        ...
```
Too many global variables can be refactored into a dictionary or a class.
For example, instead of using several global variables like char_class, race, etc., you can use a dictionary:

```python
Copy code
character_data = {
    "char_class": "",
    "race": "",
    ...
}
```
And update the code to use this dictionary instead of global variables.

Functions can be broken down into smaller functions.
For example, in your start() function, you have a lot of code in a single function. You can break this down into smaller functions based on the specific task each section is performing.

Some conditions in your code can be simplified.
For example, in your create_character() function, you have the following code:

```python
if race == "human" or race == "elf" or race == "dwarf" or race == "orc":
    ...
else:
    print("That is not a valid race.")
```
You can simplify this condition using the in operator:

```python
if race in ["human", "elf", "dwarf", "orc"]
```

Here are some follow-up suggestion:
How can I improve the efficiency of this code?
How can I make this code more readable?
How can I make this code more maintainable?

Note: Tried to specify the file this time around in hopes that Blackbox would give a better response
**-> Can you help me refactor the code that is in the file: "original_stats_test.go"**
Too many print statements can be consolidated using string concatenation or f-strings.
For example, in your create_character() function, you can consolidate all the print statements into one:

python
Copy code
def create_character():
    ...
    print("Welcome to the world of Legends, your chosen race is:", race,
          "and your chosen class is:", char_class,
          "Please note that all abilities, strengths, and weaknesses of each class are subject to change as the game evolves.")
Repeated print statements can be refactored into a separate function.
For example, you have several instances of print("What would you like to do?") in your start() function. You can create a separate function to handle this:

python
Copy code
def prompt_user():
    print("What would you like to do?")

def start():
    ...
    while True:
        prompt_user()
        choice = input().lower()
        ...
Too many global variables can be refactored into a dictionary or a class.
For example, instead of using several global variables like char_class, race, etc., you can use a dictionary:

python
Copy code
character_data = {
    "char_class": "",
    "race": "",
    ...
}
And update the code to use this dictionary instead of global variables.

Functions can be broken down into smaller functions.
For example, in your start() function, you have a lot of code in a single function. You can break this down into smaller functions based on the specific task each section is performing.

Some conditions in your code can be simplified.
For example, in your create_character() function, you have the following code:

python
Copy code
if race == "human" or race == "elf" or race == "dwarf" or race == "orc":
    ...
else:
    print("That is not a valid race.")
You can simplify this condition using the in operator:

python
Copy code
if race in ["human", "elf", "dwarf", "orc"]



Can you help me refactor the code that is in the file: "original_stats_test.go"


Go ahead and remove the original "original_stats_test.go" file from your project.

Instead of having multiple tests for different inputs in the same test file, Go promotes a single test per file philosophy.

So, you need to create separate test files for each of your tests. In your case, you have 4 different tests.

So, you need to create 4 separate test files. Name them according to the functionality they are testing. For example, you can name them as follows:

knight_stats_test.go
wizard_stats_test.go
rogue_stats_test.go
invalid_class_stats_test.go
In each of these test files, write a test function that calls your original stats function with the appropriate input. For example, in the "knight_stats_test.go" file, you can write a test function as follows:
go
Copy code
func TestKnightStats(t *testing.T) {
    stats := CalculateStats("Knight")
    assert.Equal(t, stats.Strength, 10)
    assert.Equal(t, stats.Dexterity, 5)
    assert.Equal(t, stats.Intelligence, 2)
}
Similarly, you can write test functions for each of the other 3 tests.

Don't forget to import the necessary packages at the beginning of each test file. For example:

go
Copy code
import (
    "testing"
    "github.com/stretchr/testify/assert"
)
Now, you can run these tests individually using the go test command followed by the file name. For example:
bash
Copy code
go test knight_stats_test.go
This will allow you to run the tests individually, which is useful when you want to run a single test during development or when running the tests in a continuous integration pipeline.

Here are some follow-up suggestion:
How to improve the efficiency of the code?
How to make the code more readable?
How to make the code more maintainable?

**-> Can you refactor these functions <list of Anscombe and AnscombeCoefficient functions from the original assignment**
Answer: 2
```go
func calculateRank(value float64, sortedSeries stats.Series) float64 {
    sorted := stats.Series{value, 0}
    sort.Float64s(sorted)
    for i, val := range sortedSeries {
        if val == value {
            return float64(i+1) / float64(len(sortedSeries))
        }
    }
    return 0
}

func anscombe(sortedSeries stats.Series) ([]stats.Coordinate, []stats.Coordinate) {
    var X []float64
    var Y []float64

    for _, coordinate := range sortedSeries {
        X = append(X, coordinate.X)
        Y = append(Y, coordinate.Y)
    }

    var sortedX stats.Series
    var sortedY stats.Series

    for _, val := range X {
        sortedX = append(sortedX, stats.Coordinate{X: val, Y: 0})
    }

    for _, val := range Y {
        sortedY = append(sortedY, stats.Coordinate{X: val, Y: 0})
    }

    sort.Sort(sortedX)
    sort.Sort(sortedY)

    var rankedX []float64
    var rankedY []float64

    for _, coordinate := range sortedSeries {
        rankedX = append(rankedX, calculateRank(coordinate.X, sortedX))
        rankedY = append(rankedY, calculateRank(coordinate.Y, sortedY))
    }

    var Xstar []float64
    var Ystar []float64

    meanRankedX := stats.Mean(rankedX)
    meanRankedY := stats.Mean(rankedY)

    for _, coordinate := range sortedSeries {
        Xstar = append(Xstar, 2*(coordinate.X-meanRankedX))
        Ystar = append(Ystar, 2*(coordinate.Y-meanRankedY))
    }

    XstarMean := stats.Mean(Xstar)
    YstarMean := stats.Mean(
```

**-> Is there a way I should format my prompts?**

Answer:
1. Include all relevant details: To avoid any confusion, include all necessary details in your prompt.
2. Start with a clear introduction: Begin your prompt with a clear and concise introduction. This will make it easier
for the person you are addressing to understand your point.
3.  Be concise: While it is important to provide all the necessary details, try to be concise in your explanations. This
will ensure that your prompt remains easily digestible.
4.  Maintain a respectful tone: Avoid using abrupt or demanding language. Instead, use a respectful tone to ensure that
the person you are addressing feels comfortable engaging with your prompt.
5. Provide examples: If it is necessary to provide examples to clarify your point, ensure that these examples are
relevant and easily understood.
6. Avoid unnecessary formatting: While it is acceptable to use formatting tools like bold or italics to emphasize
specific points, avoid overusing these formatting tools as it may come across as excessively emphasizing the point.
7.  Keep your tone conversational: Aim to maintain a conversational tone throughout your prompt.
This will make it feel more approachable and engaging for the person you are addressing.

Remember, the key to successful communication is clear and concise explanation of your point. By following these
guidelines, you can ensure that your prompts are well-formatted and easy to understand.

---
I decided to stop here in the interest of time. I was expecting Blackbox to more actively review the code and it doesn't
seem like it is able to do so.
