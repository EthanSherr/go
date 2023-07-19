[link](https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/9bfdb774-7cf6-4d6b-95b7-9ef498a340d8/129d9a6f-1b4b-4fbd-92c9-6bb70850e29d)

# FIRST CLASS AND HIGHER ORDER FUNCTIONS
A programming language is said to have "first-class functions" when functions in that language are treated like any other variable. For example, in such a language, a function can be passed as an argument to other functions, can be returned by another function and can be assigned as a value to a variable.

A function that returns a function or accepts a function as input is called a Higher-Order Function.

Go supports [first-class](https://developer.mozilla.org/en-US/docs/Glossary/First-class_Function) and higher-order functions. Another way to think of this is that a function is just another type -- just like `ints` and `strings` and `bools`.

For example, to accept a function as a parameter:

```go
func add(x, y int) int {
  return x + y
}

func mul(x, y int) int {
  return x * y
}

// aggregate applies the given math function to the first 3 inputs
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
  return arithmetic(arithmetic(a, b), c)
}

func main(){
  fmt.Println(aggregate(2,3,4, add))
  // prints 9
  fmt.Println(aggregate(2,3,4, mul))
  // prints 24
}
```

## ASSIGNMENT
Textio is launching a new email messaging product, "Mailio"!

Fix the compile-time bug in the `getFormattedMessages` function. The function body is correct, but the function signature is not.