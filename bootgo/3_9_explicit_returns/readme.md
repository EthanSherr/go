https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/10b92b5a-6687-474e-9df0-e215b7d0a46d/067df3cb-a240-4f10-8159-b04a737e5002

EXPLICIT RETURNS
Even though a function has named return values, we can still explicitly return values if we want to.
```go
func getCoords() (x, y int){
  return x, y // this is explicit
}
```

Using this explicit pattern we can even overwrite the return values:
```go
func getCoords() (x, y int){
  return 5, 6 // this is explicit, x and y are NOT returned
}
```

Otherwise, if we want to return the values defined in the function signature we can just use a naked return (blank return):
```go
func getCoords() (x, y int){
  return // implicitly returns x and y
}
```

ASSIGNMENT
Fix the function to return the named values implicitly.