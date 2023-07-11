https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/10b92b5a-6687-474e-9df0-e215b7d0a46d/351c4674-1c31-4148-b98f-1179dbcaac81

PASSING VARIABLES BY VALUE
Variables in Go are passed by value (except for a few data types we haven't covered yet). "Pass by value" means that when a variable is passed into a function, that function receives a copy of the variable. The function is unable to mutate the caller's original data.

func main(){
    x := 5
    increment(x)

    fmt.Println(x)
    // still prints 5,
    // because the increment function
    // received a copy of x
}

func increment(x int){
    x++
}
ASSIGNMENT
It's critical in Textio that we keep track of how many SMS messages we send on behalf of our clients. Fix the bug to accurately track the number of SMS messages sent.

Alter the incrementSends function so that it returns the result after incrementing the sendsSoFar.
Alter main() to capture the return value from incrementSends() and overwrite the previous sendsSoFar value.