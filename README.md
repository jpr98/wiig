# wiig
*Implementation of the Monkey programming language.*

This repo follows the book [Writing an Interpreter in Go](https://interpreterbook.com) written by Thorsten Ball.
The end goal is to create an interpreted programming language called Monkey. 
It is implemented in Go, with no third-party libraries.


### Parser
Monkey has a recursive descent parser. Specifically a Pratt parser, which has an operator precedence approach.
If you want to know more about this parsing algorithm [this](https://dev.to/jrop/pratt-parsing) is a great article explaining it.
