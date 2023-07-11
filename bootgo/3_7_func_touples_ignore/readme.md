https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/10b92b5a-6687-474e-9df0-e215b7d0a46d/185e65bf-8d3a-4419-abd0-258a457f0b88

IGNORING RETURN VALUES
A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore: _

For example:

func getPoint() (x int, y int) {
  return 3, 4
}

// ignore y value
x, _ := getPoint()
Even though getPoint() returns two values, we can capture the first one and ignore the second.

WHY WOULD YOU IGNORE A RETURN VALUE?
There could be many reasons. For example, maybe a function called getCircle returns the center point and the radius, but you really only need the radius for your calculation. In that case, you would ignore the center point variable.

This is crucial to understand because the Go compiler will throw an error if you have unused variable declarations in your code, so you need to ignore anything you don't intend to use.

ASSIGNMENT
In this code snippet, we only need our customer's first name. Ignore the last name so that the code compiles.