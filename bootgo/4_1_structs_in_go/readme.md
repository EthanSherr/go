https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/59d90390-f479-4472-bb16-9af599285229/81362a78-09b3-459e-bd16-c0312d3ef2d2

STRUCTS IN GO
We use structs in Go to represent structured data. It's often convenient to group different types of variables together. For example, if we want to represent a car we could do the following:

```go
type car struct {
  Make string
  Model string
  Height int
  Width int
}
```
This creates a new struct type called car. All cars have a Make, Model, Height and Width.

In Go, you will often use a struct to represent information that you would have used a dictionary for in Python, or an object literal for in JavaScript.

ASSIGNMENT
Complete the messageToSend struct definition. It needs two fields:

phoneNumber - an integer
message - a string.